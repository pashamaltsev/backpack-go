package models

import (
	"time"

	"github.com/feeeei/backpack-go/utils"
)

type Order struct {
	OrderType              OrderType               `json:"orderType"`
	ID                     string                  `json:"id"`
	ClientID               *int                    `json:"clientId"`
	CreatedAt              time.Time               `json:"createdAt" time_format:"unix"`
	ExecutedQuantity       float64                 `json:"executedQuantity,string"`
	ExecutedQuoteQuantity  float64                 `json:"executedQuoteQuantity,string"`
	Quantity               *float64                `json:"quantity,string"`
	QuoteQuantity          *float64                `json:"quoteQuantity,string"`
	ReduceOnly             *bool                   `json:"reduceOnly"`
	TimeInForce            TimeInForce             `json:"timeInForce"`
	SelfTradePrevention    SelfTradePreventionType `json:"selfTradePrevention"`
	Side                   Side                    `json:"side"`
	Status                 OrderStatus             `json:"status"`
	StopLossTriggerPrice   *float64                `json:"stopLossTriggerPrice,string"`
	StopLossLimitPrice     *float64                `json:"stopLossLimitPrice,string"`
	StopLossTriggerBy      *string                 `json:"stopLossTriggerBy"`
	Symbol                 string                  `json:"symbol"`
	TakeProfitTriggerPrice *float64                `json:"takeProfitTriggerPrice,string"`
	TakeProfitLimitPrice   *float64                `json:"takeProfitLimitPrice,string"`
	TakeProfitTriggerBy    *string                 `json:"takeProfitTriggerBy"`
	TriggerBy              *float64                `json:"triggerBy,string"`
	TriggerPrice           *float64                `json:"triggerPrice,string"`
	TriggerQuantity        *float64                `json:"triggerQuantity,string"`
	TriggeredAt            *time.Time              `json:"triggeredAt" time_format:"unix"`
}

type Side string

const (
	Bid  Side = "Bid"
	Ask  Side = "Ask"
	Buy  Side = "Bid"
	Sell Side = "Ask"
)

type OrderType string

const (
	OrderTypeMarket OrderType = "Market"
	OrderTypeLimit  OrderType = "Limit"
)

type SelfTradePreventionType string

const (
	SelfTradePreventionTypeRejectTaker SelfTradePreventionType = "RejectTaker"
	SelfTradePreventionTypeRejectMaker SelfTradePreventionType = "RejectMaker"
	SelfTradePreventionTypeRejectBoth  SelfTradePreventionType = "RejectBoth"
)

type OrderStatus string

const (
	OrderStatusCancelled       OrderStatus = "Cancelled"
	OrderStatusExpired         OrderStatus = "Expired"
	OrderStatusFilled          OrderStatus = "Filled"
	OrderStatusNew             OrderStatus = "New"
	OrderStatusPartiallyFilled OrderStatus = "PartiallyFilled"
	OrderStatusTriggerPending  OrderStatus = "TriggerPending"
	OrderStatusTriggerFilled   OrderStatus = "TriggerFilled"
)

type TimeInForce string

const (
	TimeInForceGTC TimeInForce = "GTC"
	TimeInForceIOC TimeInForce = "IOC"
	TimeInForceFOK TimeInForce = "FOK"
)

type OrderOptions struct {
	Quantity               *float64                 `json:"quantity,string,omitempty"`
	AutoLend               *bool                    `json:"autoLend,omitempty"`
	AutoLendRedeem         *bool                    `json:"autoLendRedeem,omitempty"`
	AutoBorrow             *bool                    `json:"autoBorrow,omitempty"`
	AutoBorrowRepay        *bool                    `json:"autoBorrowRepay,omitempty"`
	ClientID               *int32                   `json:"clientId,omitempty"`
	PostOnly               *bool                    `json:"postOnly,omitempty"`
	ReduceOnly             *bool                    `json:"reduceOnly,omitempty"`
	SelfTradePrevention    *SelfTradePreventionType `json:"selfTradePrevention,omitempty"`
	StopLossLimitPrice     *float64                 `json:"stopLossLimitPrice,string,omitempty"`
	StopLossTriggerPrice   *float64                 `json:"stopLossTriggerPrice,string,omitempty"`
	TakeProfitLimitPrice   *float64                 `json:"takeProfitLimitPrice,string,omitempty"`
	TakeProfitTriggerPrice *float64                 `json:"takeProfitTriggerPrice,string,omitempty"`
	TimeInForce            *TimeInForce             `json:"timeInForce,omitempty"`
	TriggerPrice           *float64                 `json:"triggerPrice,string,omitempty"`
	TriggerQuantity        *float64                 `json:"triggerQuantity,string,omitempty"`
}

func (o *OrderOptions) ToParams() map[string]any {
	return utils.StructToMap[map[string]any](o)
}

type OrderOption func(*OrderOptions)

func WithQuantity(quantity float64) OrderOption {
	return func(o *OrderOptions) {
		o.Quantity = &quantity
	}
}

func WithTrigger(price, quantity float64) OrderOption {
	return func(o *OrderOptions) {
		o.TriggerPrice = &price
		o.TriggerQuantity = &quantity
	}
}

func WithStopLoss(triggerPrice, limitPrice float64) OrderOption {
	return func(o *OrderOptions) {
		o.StopLossLimitPrice = &limitPrice
		o.StopLossTriggerPrice = &triggerPrice
	}
}

func WithTakeProfit(triggerPrice, limitPrice float64) OrderOption {
	return func(o *OrderOptions) {
		o.TakeProfitTriggerPrice = &triggerPrice
		o.TakeProfitLimitPrice = &limitPrice
	}
}

func WithAutoLend(autoLend bool) OrderOption {
	return func(o *OrderOptions) {
		o.AutoLend = &autoLend
	}
}

func WithAutoLendRedeem(autoLendRedeem bool) OrderOption {
	return func(o *OrderOptions) {
		o.AutoLendRedeem = &autoLendRedeem
	}
}

func WithAutoBorrow(autoBorrow bool) OrderOption {
	return func(o *OrderOptions) {
		o.AutoBorrow = &autoBorrow
	}
}

func WithAutoBorrowRepay(autoBorrowRepay bool) OrderOption {
	return func(o *OrderOptions) {
		o.AutoBorrowRepay = &autoBorrowRepay
	}
}

func WithClientID(clientID int32) OrderOption {
	return func(o *OrderOptions) {
		o.ClientID = &clientID
	}
}

func WithPostOnly(postOnly bool) OrderOption {
	return func(o *OrderOptions) {
		o.PostOnly = &postOnly
	}
}

func WithReduceOnly(reduceOnly bool) OrderOption {
	return func(o *OrderOptions) {
		o.ReduceOnly = &reduceOnly
	}
}

func WithSelfTradePrevention(selfTradePrevention SelfTradePreventionType) OrderOption {
	return func(o *OrderOptions) {
		o.SelfTradePrevention = &selfTradePrevention
	}
}

func WithStopLossTriggerPrice(stopLossTriggerPrice float64) OrderOption {
	return func(o *OrderOptions) {
		o.StopLossTriggerPrice = &stopLossTriggerPrice
	}
}

func WithTakeProfitLimitPrice(takeProfitLimitPrice float64) OrderOption {
	return func(o *OrderOptions) {
		o.TakeProfitLimitPrice = &takeProfitLimitPrice
	}
}

func WithTakeProfitTriggerPrice(takeProfitTriggerPrice float64) OrderOption {
	return func(o *OrderOptions) {
		o.TakeProfitTriggerPrice = &takeProfitTriggerPrice
	}
}

func WithTimeInForce(timeInForce TimeInForce) OrderOption {
	return func(o *OrderOptions) {
		o.TimeInForce = &timeInForce
	}
}

func WithTriggerPrice(triggerPrice float64) OrderOption {
	return func(o *OrderOptions) {
		o.TriggerPrice = &triggerPrice
	}
}

func WithTriggerQuantity(triggerQuantity float64) OrderOption {
	return func(o *OrderOptions) {
		o.TriggerQuantity = &triggerQuantity
	}
}
