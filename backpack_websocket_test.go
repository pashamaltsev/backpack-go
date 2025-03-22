package backpackgo

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/feeeei/backpack-go/models"
	"github.com/feeeei/backpack-go/options"
	"github.com/feeeei/backpack-go/websocket"
)

func TestBackpackWebsocket(t *testing.T) {
	client, err := NewBackpackWebsocket(websocket.WithAPIToken(os.Getenv("API_KEY"), os.Getenv("API_SECRET")))
	if err != nil {
		t.Fatal(err)
	}

	symbol := "BTC_USDC"

	// test SubscribeBookTicker
	t.Run("test SubscribeBookTicker", func(t *testing.T) {
		wait := make(chan struct{})
		err = client.SubscribeBookTicker(symbol, func(bookTicker *models.BookTickerUpdate) {
			wait <- struct{}{}
		})
		if err != nil {
			t.Errorf("SubscribeBookTicker failed: %v", err)
		} else {
			<-wait
			client.Unsubscribe(fmt.Sprintf("bookTicker.%s", symbol))
			time.Sleep(time.Second)
			fmt.Println("OK: SubscribeBookTicker")
		}
	})

	// test SubscribeDepth
	t.Run("test SubscribeDepth", func(t *testing.T) {
		wait := make(chan struct{})
		err = client.SubscribeDepth(symbol, func(depth *models.DepthUpdate) {
			wait <- struct{}{}
		})
		if err != nil {
			t.Errorf("SubscribeDepth failed: %v", err)
		} else {
			<-wait
			client.Unsubscribe(fmt.Sprintf("depth.%s", symbol))
			time.Sleep(time.Second)
			fmt.Println("OK: SubscribeDepth")
		}
	})

	// test SubscribeKLine
	t.Run("test SubscribeKLine", func(t *testing.T) {
		wait := make(chan struct{})
		err = client.SubscribeKLine(options.KLineInterval1m, symbol, func(kline *models.KLineUpdate) {
			wait <- struct{}{}
		})
		if err != nil {
			t.Errorf("SubscribeKLine failed: %v", err)
		} else {
			<-wait
			client.Unsubscribe(fmt.Sprintf("kline.%s.%s", options.KLineInterval1m, symbol))
			time.Sleep(time.Second)
			fmt.Println("OK: SubscribeKLine")
		}
	})

	// test SubscribeLiquidation
	t.Run("test SubscribeLiquidation", func(t *testing.T) {
		wait := make(chan struct{})
		err = client.SubscribeLiquidation(symbol, func(liquidation *models.LiquidationUpdate) {
			wait <- struct{}{}
		})
		if err != nil {
			t.Errorf("SubscribeLiquidation failed: %v", err)
		} else {
			<-wait
			client.Unsubscribe(fmt.Sprintf("liquidation.%s", symbol))
			time.Sleep(time.Second)
			fmt.Println("OK: SubscribeLiquidation")
		}
	})

	// test SubscribeMarkPrice
	t.Run("test SubscribeMarkPrice", func(t *testing.T) {
		wait := make(chan struct{})
		err = client.SubscribeMarkPrice(symbol+"_PERP", func(markPrice *models.MarkPriceUpdate) {
			wait <- struct{}{}
		})
		if err != nil {
			t.Errorf("SubscribeMarkPrice failed: %v", err)
		} else {
			<-wait
			client.Unsubscribe(fmt.Sprintf("markPrice.%s", symbol+"_PERP"))
			time.Sleep(time.Second)
			fmt.Println("OK: SubscribeMarkPrice")
		}
	})

	// test SubscribeTicker
	t.Run("test SubscribeTicker", func(t *testing.T) {
		wait := make(chan struct{})
		err = client.SubscribeTicker(symbol, func(ticker *models.TickerUpdate) {
			wait <- struct{}{}
		})
		if err != nil {
			t.Errorf("SubscribeTicker failed: %v", err)
		} else {
			<-wait
			client.Unsubscribe(fmt.Sprintf("ticker.%s", symbol))
			time.Sleep(time.Second)
			fmt.Println("OK: SubscribeTicker")
		}
	})

	// test SubscribeTrade
	t.Run("test SubscribeTrade", func(t *testing.T) {
		wait := make(chan struct{})
		err = client.SubscribeTrade(symbol, func(trade *models.TradeUpdate) {
			wait <- struct{}{}
		})
		if err != nil {
			t.Errorf("SubscribeTrade failed: %v", err)
		} else {
			<-wait
			client.Unsubscribe(fmt.Sprintf("trade.%s", symbol))
			time.Sleep(time.Second)
			fmt.Println("OK: SubscribeTrade")
		}
	})

	// test SubscribeOrderUpdate
	t.Run("test SubscribeOrderUpdate", func(t *testing.T) {
		wait := make(chan struct{})
		err = client.SubscribeOrderUpdate(func(order *models.OrderUpdate) {
			wait <- struct{}{}
		})
		if err != nil {
			t.Errorf("SubscribeOrderUpdate failed: %v", err)
		} else {
			<-wait
			fmt.Println("OK: SubscribeOrderUpdate")
		}
	})

	// test SubscribePositionUpdate
	t.Run("test SubscribePositionUpdate", func(t *testing.T) {
		wait := make(chan struct{})
		err = client.SubscribePositionUpdate(func(position *models.PositionUpdate) {
			wait <- struct{}{}
		})
		if err != nil {
			t.Errorf("SubscribePositionUpdate failed: %v", err)
		} else {
			<-wait
			fmt.Println("OK: SubscribePositionUpdate")
		}
	})

	// TODO: test SubscribeRFQUpdate
	// // test SubscribeRFQUpdate
	// t.Run("test SubscribeRFQUpdate", func(t *testing.T) {
	// 	wait := make(chan struct{})
	// 	err = client.SubscribeRFQUpdate(func(rfq *models.RFQUpdate) {
	// 		wait <- struct{}{}
	// 	})
	// 	if err != nil {
	// 		t.Errorf("SubscribeRFQUpdate failed: %v", err)
	// 	} else {
	// 		<-wait
	// 		fmt.Println("OK: SubscribeRFQUpdate")
	// 	}
	// })
}
