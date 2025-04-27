package backpackgo

import (
	"fmt"
	"strings"
	"sync"
	"time"

	encodingjson "encoding/json"

	"github.com/UnipayFI/backpack-go/models"
	"github.com/UnipayFI/backpack-go/options"
	ops "github.com/UnipayFI/backpack-go/websocket"
	"github.com/go-json-experiment/json"
	"github.com/gorilla/websocket"
)

type BackpackWebsocket struct {
	BaseURL   string
	APIKey    string
	APISecret string
	Windows   time.Duration
	conn      *websocket.Conn
	handlers  map[string]func(encodingjson.RawMessage)
	pool      sync.Pool
}

type Message struct {
	Stream string                  `json:"stream"`
	Data   encodingjson.RawMessage `json:"data"`
}

func NewBackpackWebsocket(options ...ops.Options) (*BackpackWebsocket, error) {
	opts := ops.DefaultWebSocketOptions()
	dialer := websocket.DefaultDialer

	for _, option := range options {
		option(opts, dialer)
	}

	conn, _, err := dialer.Dial(opts.BaseURL, nil)
	if err != nil {
		return nil, err
	}

	websocket := &BackpackWebsocket{
		BaseURL:   opts.BaseURL,
		APIKey:    opts.APIKey,
		APISecret: opts.APISecret,
		Windows:   opts.Windows,
		conn:      conn,
		handlers:  make(map[string]func(encodingjson.RawMessage)),
		pool: sync.Pool{
			New: func() any {
				return &Message{}
			},
		},
	}
	go websocket.loop()
	return websocket, nil
}

func (client *BackpackWebsocket) SubscribeOrderUpdate(handler func(*models.OrderUpdate)) error {
	return SubscribePrivateStream(client, "account.orderUpdate", new(models.OrderUpdate), handler)
}

func (client *BackpackWebsocket) SubscribeOrderUpdateWithSymbol(symbol string, handler func(*models.OrderUpdate)) error {
	return SubscribePrivateStream(client, fmt.Sprintf("account.orderUpdate.%s", symbol), new(models.OrderUpdate), handler)
}

func (client *BackpackWebsocket) SubscribePositionUpdate(handler func(*models.PositionUpdate)) error {
	return SubscribePrivateStream(client, "account.positionUpdate", new(models.PositionUpdate), handler)
}

func (client *BackpackWebsocket) SubscribePositionUpdateWithSymbol(symbol string, handler func(*models.PositionUpdate)) error {
	return SubscribePrivateStream(client, fmt.Sprintf("account.positionUpdate.%s", symbol), new(models.PositionUpdate), handler)
}

func (client *BackpackWebsocket) SubscribeRFQUpdate(handler func(*models.RFQUpdate)) error {
	return SubscribePrivateStream(client, "account.rfqUpdate", new(models.RFQUpdate), handler)
}

func (client *BackpackWebsocket) SubscribeRFQUpdateWithSymbol(symbol string, handler func(*models.RFQUpdate)) error {
	return SubscribePrivateStream(client, fmt.Sprintf("account.rfqUpdate.%s", symbol), new(models.RFQUpdate), handler)
}

func (client *BackpackWebsocket) SubscribeBookTicker(symbol string, handler func(*models.BookTickerUpdate)) error {
	return SubscribePublicStream(client, fmt.Sprintf("bookTicker.%s", symbol), new(models.BookTickerUpdate), handler)
}

func (client *BackpackWebsocket) SubscribeDepth(symbol string, handler func(*models.DepthUpdate)) error {
	return SubscribePublicStream(client, fmt.Sprintf("depth.%s", symbol), new(models.DepthUpdate), handler)
}

func (client *BackpackWebsocket) Subscribe200msDepth(symbol string, handler func(*models.DepthUpdate)) error {
	return SubscribePublicStream(client, fmt.Sprintf("depth.200ms.%s", symbol), new(models.DepthUpdate), handler)
}

func (client *BackpackWebsocket) SubscribeKLine(interval options.KLineInterval, symbol string, handler func(*models.KLineUpdate)) error {
	return SubscribePublicStream(client, fmt.Sprintf("kline.%s.%s", interval, symbol), new(models.KLineUpdate), handler)
}

func (client *BackpackWebsocket) SubscribeLiquidation(symbol string, handler func(*models.LiquidationUpdate)) error {
	return SubscribePublicStream(client, fmt.Sprintf("liquidation.%s", symbol), new(models.LiquidationUpdate), handler)
}

func (client *BackpackWebsocket) SubscribeMarkPrice(symbol string, handler func(*models.MarkPriceUpdate)) error {
	return SubscribePublicStream(client, fmt.Sprintf("markPrice.%s", symbol), new(models.MarkPriceUpdate), handler)
}

func (client *BackpackWebsocket) SubscribeTicker(symbol string, handler func(*models.TickerUpdate)) error {
	return SubscribePublicStream(client, fmt.Sprintf("ticker.%s", symbol), new(models.TickerUpdate), handler)
}

func (client *BackpackWebsocket) SubscribeOpenInterest(symbol string, handler func(*models.OpenInterestUpdate)) error {
	return SubscribePublicStream(client, fmt.Sprintf("openInterest.%s", symbol), new(models.OpenInterestUpdate), handler)
}

func (client *BackpackWebsocket) SubscribeTrade(symbol string, handler func(*models.TradeUpdate)) error {
	return SubscribePublicStream(client, fmt.Sprintf("trade.%s", symbol), new(models.TradeUpdate), handler)
}

func (client *BackpackWebsocket) Unsubscribe(stream string) error {
	request, err := client.sign(stream, "unsubscribe")
	if err != nil {
		return err
	}
	return client.conn.WriteJSON(request)
}

func SubscribePrivateStream[T any](client *BackpackWebsocket, stream string, newable models.Newable, handler func(T)) error {
	request, err := client.sign(stream, "subscribe")
	if err != nil {
		return err
	}
	client.handlers[stream] = func(payload encodingjson.RawMessage) {
		obj := newable.New()
		json.Unmarshal(payload, &obj)
		handler(obj.(T))
	}
	return client.conn.WriteJSON(request)
}

func SubscribePublicStream[T any](client *BackpackWebsocket, stream string, newable models.Newable, handler func(T)) error {
	request := client.payload(stream, "subscribe")
	client.handlers[stream] = func(payload encodingjson.RawMessage) {
		obj := newable.New()
		json.Unmarshal(payload, &obj)
		handler(obj.(T))
	}
	return client.conn.WriteJSON(request)
}

func (client *BackpackWebsocket) loop() {
	for {
		_, message, err := client.conn.ReadMessage()
		if err != nil {
			return
		}
		client.handleMessage(message)
	}
}

func (client *BackpackWebsocket) handleMessage(message []byte) {
	msg := client.pool.Get().(*Message)
	json.Unmarshal(message, &msg)
	if client.handlers[msg.Stream] != nil {
		go client.handlers[msg.Stream](msg.Data)
	}
	client.pool.Put(msg)
}

func (client *BackpackWebsocket) sign(stream, instruction string) (map[string]any, error) {
	payload := client.payload(stream, instruction)
	timestamp := time.Now().UnixNano() / int64(time.Millisecond)
	signature, err := sign(client.APISecret, nil, instruction, timestamp, client.Windows.Milliseconds())
	if err == nil {
		payload["signature"] = []string{client.APIKey, signature, fmt.Sprintf("%v", timestamp), fmt.Sprintf("%v", client.Windows.Milliseconds())}
		return payload, nil
	} else {
		return nil, err
	}
}

func (client *BackpackWebsocket) payload(stream, instruction string) map[string]any {
	return map[string]any{"method": strings.ToUpper(instruction), "params": []string{stream}}
}
