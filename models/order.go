package models

import "time"

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
	SideBid Side = "Bid"
	SideAsk Side = "Ask"
)

type OrderType string

const (
	OrderTypeMarket     OrderType = "Market"
	OrderTypeLimit      OrderType = "Limit"
	OrderTypeStopMarket OrderType = "StopMarket"
	OrderTypeStopLimit  OrderType = "StopLimit"
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
	AutoLend               *bool                    `json:"autoLend"`
	AutoLendRedeem         *bool                    `json:"autoLendRedeem"`
	AutoBorrow             *bool                    `json:"autoBorrow"`
	AutoBorrowRepay        *bool                    `json:"autoBorrowRepay"`
	ClientID               *int32                   `json:"clientId"`
	PostOnly               *bool                    `json:"postOnly"`
	ReduceOnly             *bool                    `json:"reduceOnly"`
	SelfTradePrevention    *SelfTradePreventionType `json:"selfTradePrevention"`
	StopLossLimitPrice     *float64                 `json:"stopLossLimitPrice,string"`
	StopLossTriggerPrice   *float64                 `json:"stopLossTriggerPrice,string"`
	TakeProfitLimitPrice   *float64                 `json:"takeProfitLimitPrice,string"`
	TakeProfitTriggerPrice *float64                 `json:"takeProfitTriggerPrice,string"`
	TimeInForce            *TimeInForce             `json:"timeInForce"`
	TriggerPrice           *float64                 `json:"triggerPrice,string"`
	TriggerQuantity        *float64                 `json:"triggerQuantity,string"`
}
