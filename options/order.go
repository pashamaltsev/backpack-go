package options

import "github.com/UnipayFI/backpack-go/utils"

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

type Side string

const (
	Bid  Side = "Bid"
	Ask  Side = "Ask"
	Buy  Side = "Bid"
	Sell Side = "Ask"
)

type TimeInForce string

const (
	TimeInForceGTC TimeInForce = "GTC"
	TimeInForceIOC TimeInForce = "IOC"
	TimeInForceFOK TimeInForce = "FOK"
)

type OrderOptions struct {
	AutoLend               *bool                    `json:"autoLend,omitempty"`
	AutoLendRedeem         *bool                    `json:"autoLendRedeem,omitempty"`
	AutoBorrow             *bool                    `json:"autoBorrow,omitempty"`
	AutoBorrowRepay        *bool                    `json:"autoBorrowRepay,omitempty"`
	ClientID               *uint32                  `json:"clientId,omitempty"`
	PostOnly               *bool                    `json:"postOnly,omitempty"`
	ReduceOnly             *bool                    `json:"reduceOnly,omitempty"`
	SelfTradePrevention    *SelfTradePreventionType `json:"selfTradePrevention,omitempty"`
	StopLossLimitPrice     *float64                 `json:"stopLossLimitPrice,string,omitempty"`
	StopLossTriggerPrice   *float64                 `json:"stopLossTriggerPrice,string,omitempty"`
	TakeProfitLimitPrice   *float64                 `json:"takeProfitLimitPrice,string,omitempty"`
	TakeProfitTriggerPrice *float64                 `json:"takeProfitTriggerPrice,string,omitempty"`
	TimeInForce            *TimeInForce             `json:"timeInForce,omitempty"`
}

func (o *OrderOptions) ToParams() map[string]any {
	return utils.StructToMap[map[string]any](o)
}

type OrderOptionFn func(*OrderOptions)

func WithStopLoss(triggerPrice, limitPrice float64) OrderOptionFn {
	return func(o *OrderOptions) {
		o.StopLossLimitPrice = &limitPrice
		o.StopLossTriggerPrice = &triggerPrice
	}
}

func WithTakeProfit(triggerPrice, limitPrice float64) OrderOptionFn {
	return func(o *OrderOptions) {
		o.TakeProfitTriggerPrice = &triggerPrice
		o.TakeProfitLimitPrice = &limitPrice
	}
}

func WithAutoLend(autoLend bool) OrderOptionFn {
	return func(o *OrderOptions) {
		o.AutoLend = &autoLend
	}
}

func WithAutoLendRedeem(autoLendRedeem bool) OrderOptionFn {
	return func(o *OrderOptions) {
		o.AutoLendRedeem = &autoLendRedeem
	}
}

func WithAutoBorrow(autoBorrow bool) OrderOptionFn {
	return func(o *OrderOptions) {
		o.AutoBorrow = &autoBorrow
	}
}

func WithAutoBorrowRepay(autoBorrowRepay bool) OrderOptionFn {
	return func(o *OrderOptions) {
		o.AutoBorrowRepay = &autoBorrowRepay
	}
}

func WithClientID(clientID uint32) OrderOptionFn {
	return func(o *OrderOptions) {
		o.ClientID = &clientID
	}
}

func WithPostOnly(postOnly bool) OrderOptionFn {
	return func(o *OrderOptions) {
		o.PostOnly = &postOnly
	}
}

func WithReduceOnly(reduceOnly bool) OrderOptionFn {
	return func(o *OrderOptions) {
		o.ReduceOnly = &reduceOnly
	}
}

func WithSelfTradePrevention(selfTradePrevention SelfTradePreventionType) OrderOptionFn {
	return func(o *OrderOptions) {
		o.SelfTradePrevention = &selfTradePrevention
	}
}

func WithTimeInForce(timeInForce TimeInForce) OrderOptionFn {
	return func(o *OrderOptions) {
		o.TimeInForce = &timeInForce
	}
}
