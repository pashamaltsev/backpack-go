package models

import (
	"time"

	"github.com/feeeei/backpack-go/options"
)

type Order struct {
	OrderType              OrderType                       `json:"orderType"`
	ID                     string                          `json:"id"`
	ClientID               *int                            `json:"clientId"`
	CreatedAt              time.Time                       `json:"createdAt" time_format:"unix"`
	ExecutedQuantity       float64                         `json:"executedQuantity,string"`
	ExecutedQuoteQuantity  float64                         `json:"executedQuoteQuantity,string"`
	Quantity               *float64                        `json:"quantity,string"`
	QuoteQuantity          *float64                        `json:"quoteQuantity,string"`
	ReduceOnly             *bool                           `json:"reduceOnly"`
	TimeInForce            options.TimeInForce             `json:"timeInForce"`
	SelfTradePrevention    options.SelfTradePreventionType `json:"selfTradePrevention"`
	Side                   options.Side                    `json:"side"`
	Status                 OrderStatus                     `json:"status"`
	StopLossTriggerPrice   *float64                        `json:"stopLossTriggerPrice,string"`
	StopLossLimitPrice     *float64                        `json:"stopLossLimitPrice,string"`
	StopLossTriggerBy      *string                         `json:"stopLossTriggerBy"`
	Symbol                 string                          `json:"symbol"`
	TakeProfitTriggerPrice *float64                        `json:"takeProfitTriggerPrice,string"`
	TakeProfitLimitPrice   *float64                        `json:"takeProfitLimitPrice,string"`
	TakeProfitTriggerBy    *string                         `json:"takeProfitTriggerBy"`
	TriggerBy              *float64                        `json:"triggerBy,string"`
	TriggerPrice           *float64                        `json:"triggerPrice,string"`
	TriggerQuantity        *float64                        `json:"triggerQuantity,string"`
	TriggeredAt            *time.Time                      `json:"triggeredAt" time_format:"unix"`
}

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
