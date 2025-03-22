package models

import (
	"time"

	"github.com/feeeei/backpack-go/options"
)

type Newable interface {
	New() any
}

type OrderUpdate struct {
	EventType               string                          `json:"e"`
	EventTime               time.Time                       `json:"E" time_format:"unixmicro"`
	Symbol                  string                          `json:"s"`
	ClientOrderID           int64                           `json:"c"`
	Side                    options.Side                    `json:"S"`
	OrderType               OrderType                       `json:"o"`
	TimeInForce             options.TimeInForce             `json:"f"`
	Quantity                float64                         `json:"q,string"`
	QuantityInQuote         float64                         `json:"Q,string"`
	Price                   float64                         `json:"p,string"`
	TriggerPrice            float64                         `json:"P,string"`
	TakeProfitTriggerPrice  float64                         `json:"a,string"`
	StopLossTriggerPrice    float64                         `json:"b,string"`
	TriggerQuantity         float64                         `json:"Y,string"`
	OrderStatus             OrderStatus                     `json:"X"`
	OrderExpiryReason       OrderExpiryReason               `json:"R"`
	OrderID                 string                          `json:"i"`
	TradeID                 int64                           `json:"t"`
	FillQuantity            float64                         `json:"l,string"`
	ExecutedQuantity        float64                         `json:"z,string"`
	ExecutedQuantityInQuote float64                         `json:"Z,string"`
	FillPrice               float64                         `json:"L,string"`
	IsMaker                 bool                            `json:"m"`
	Fee                     float64                         `json:"n,string"`
	FeeSymbol               string                          `json:"N"`
	SelfTradePrevention     options.SelfTradePreventionType `json:"V"`
	EngineTimestamp         time.Time                       `json:"T" time_format:"unixmicro"`
	Origin                  string                          `json:"O"`
}

func (u *OrderUpdate) New() any {
	return &OrderUpdate{}
}

type PositionUpdate struct {
	EventType                 string    `json:"e"`
	EventTime                 time.Time `json:"E" time_format:"unixmicro"`
	Symbol                    string    `json:"s"`
	BreakEventPrice           float64   `json:"b,string"`
	EntryPrice                float64   `json:"B,string"`
	EstimatedLiquidationPrice float64   `json:"l,string"`
	InitialMarginFraction     float64   `json:"f,string"`
	MarkPrice                 float64   `json:"M,string"`
	MaintenanceMarginFraction float64   `json:"m,string"`
	NetQuantity               float64   `json:"q,string"`
	NetExposureQuantity       float64   `json:"Q,string"`
	NetExposureNotional       float64   `json:"n,string"`
	PositionID                string    `json:"i"`
	PnLRealized               float64   `json:"p,string"`
	PnLUnrealized             float64   `json:"P,string"`
	EngineTimestamp           time.Time `json:"T" time_format:"unixmicro"`
}

func (u *PositionUpdate) New() any {
	return &PositionUpdate{}
}

type RFQUpdate struct {
	EventType       string       `json:"e"`
	EventTime       time.Time    `json:"E" time_format:"unixmicro"`
	RFQID           string       `json:"R"`
	Symbol          string       `json:"s"`
	QuoteID         string       `json:"Q"`
	ClientQuoteID   string       `json:"C"`
	Side            options.Side `json:"S"`
	Price           float64      `json:"p,string"`
	Quantity        float64      `json:"q,string"`
	SubmissionTime  time.Time    `json:"w" time_format:"unixmilli"`
	ExpiryTime      time.Time    `json:"W" time_format:"unixmilli"`
	Status          string       `json:"X"`
	EngineTimestamp time.Time    `json:"T" time_format:"unixmicro"`
}

func (u *RFQUpdate) New() any {
	return &RFQUpdate{}
}

type BookTickerUpdate struct {
	EventType  string    `json:"e"`
	EventTime  time.Time `json:"E" time_format:"unixmicro"`
	Symbol     string    `json:"s"`
	AskPrice   float64   `json:"a,string"`
	AskSize    float64   `json:"A,string"`
	BidPrice   float64   `json:"b,string"`
	BidSize    float64   `json:"B,string"`
	UpdateID   string    `json:"u"`
	EngineTime time.Time `json:"T" time_format:"unixmicro"`
}

func (u *BookTickerUpdate) New() any {
	return &BookTickerUpdate{}
}

type DepthUpdate struct {
	EventType     string      `json:"e"`
	EventTime     time.Time   `json:"E" time_format:"unixmicro"`
	Symbol        string      `json:"s"`
	Asks          []DepthItem `json:"a"`
	Bids          []DepthItem `json:"b"`
	FirstUpdateID int64       `json:"U"`
	LastUpdateID  int64       `json:"u"`
	EngineTime    time.Time   `json:"T" time_format:"unixmicro"`
}

func (u *DepthUpdate) New() any {
	return &DepthUpdate{}
}

type KLineUpdate struct {
	EventType string    `json:"e"`
	EventTime time.Time `json:"E" time_format:"unixmicro"`
	StartTime time.Time `json:"t" time_format:"2006-01-02T15:04:05"`
	CloseTime time.Time `json:"T" time_format:"2006-01-02T15:04:05"`
	Open      float64   `json:"o,string"`
	Close     float64   `json:"c,string"`
	High      float64   `json:"h,string"`
	Low       float64   `json:"l,string"`
	Volume    float64   `json:"v,string"`
	Trades    int       `json:"n"`
	IsClosed  bool      `json:"X"`
}

func (u *KLineUpdate) New() any {
	return &KLineUpdate{}
}

type LiquidationUpdate struct {
	EventType  string       `json:"e"`
	EventTime  time.Time    `json:"E" time_format:"unixmicro"`
	Quantity   float64      `json:"q,string"`
	Price      float64      `json:"p,string"`
	Side       options.Side `json:"S"`
	Symbol     string       `json:"s"`
	EngineTime time.Time    `json:"T" time_format:"unixmicro"`
}

func (u *LiquidationUpdate) New() any {
	return &LiquidationUpdate{}
}

type MarkPriceUpdate struct {
	EventType            string    `json:"e"`
	EventTime            time.Time `json:"E" time_format:"unixmicro"`
	Symbol               string    `json:"s"`
	MarkPrice            float64   `json:"p,string"`
	EstimatedFundingRate float64   `json:"f,string"`
	IndexPrice           float64   `json:"i,string"`
	NextFundingTimestamp time.Time `json:"n" time_format:"unixmicro"`
}

func (u *MarkPriceUpdate) New() any {
	return &MarkPriceUpdate{}
}

type TickerUpdate struct {
	EventType   string    `json:"e"`
	EventTime   time.Time `json:"E" time_format:"unixmicro"`
	Symbol      string    `json:"s"`
	FirstPrice  float64   `json:"o,string"`
	LastPrice   float64   `json:"c,string"`
	HighPrice   float64   `json:"h,string"`
	LowPrice    float64   `json:"l,string"`
	BaseVolume  float64   `json:"v,string"`
	QuoteVolume float64   `json:"V,string"`
	Trades      int       `json:"n"`
}

func (u *TickerUpdate) New() any {
	return &TickerUpdate{}
}

type TradeUpdate struct {
	EventType     string    `json:"e"`
	EventTime     time.Time `json:"E" time_format:"unixmicro"`
	Symbol        string    `json:"s"`
	Price         float64   `json:"p,string"`
	Quantity      float64   `json:"q,string"`
	BuyerOrderID  string    `json:"b"`
	SellerOrderID string    `json:"a"`
	TradeID       int       `json:"t"`
	EngineTime    time.Time `json:"T" time_format:"unixmicro"`
	IsMaker       bool      `json:"m"`
}

func (u *TradeUpdate) New() any {
	return &TradeUpdate{}
}
