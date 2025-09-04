package models

import "github.com/pashamaltsev/backpack-go/options"

type Account struct {
	AccountUpdateble
	AutoRealizePnl  bool    `json:"autoRealizePnl"`
	BorrowLimit     float64 `json:"borrowLimit,string"`
	FuturesMakerFee float64 `json:"futuresMakerFee,string"`
	FuturesTakerFee float64 `json:"futuresTakerFee,string"`
	LimitOrders     int     `json:"limitOrders"`
	Liquidating     bool    `json:"liquidating"`
	PositionLimit   int     `json:"positionLimit,string"`
	SpotMakerFee    float64 `json:"spotMakerFee,string"`
	SpotTakerFee    float64 `json:"spotTakerFee,string"`
	TriggerOrders   int     `json:"triggerOrders"`
}

type AccountUpdateble struct {
	AutoBorrowSettlements *bool    `json:"autoBorrowSettlements,omitempty"`
	AutoLend              *bool    `json:"autoLend,omitempty"`
	AutoRepayBorrows      *bool    `json:"autoRepayBorrows,omitempty"`
	LeverageLimit         *float64 `json:"leverageLimit,string,omitempty"`
}

type AccountBorrowLimit struct {
	Symbol            string  `json:"symbol"`
	MaxBorrowQuantity float64 `json:"maxBorrowQuantity,string"`
}

type AccountOrderLimit struct {
	Symbol           string       `json:"symbol"`
	Side             options.Side `json:"side"`
	MaxOrderQuantity float64      `json:"maxOrderQuantity,string"`
	options.AccountOrderLimitOptions
}

type AccountWithdrawalLimit struct {
	MaxWithdrawalQuantity string `json:"maxWithdrawalQuantity"`
	Symbol                string `json:"symbol"`
	AutoBorrow            *bool  `json:"autoBorrow,omitempty"`
	AutoLendRedeem        *bool  `json:"autoLendRedeem,omitempty"`
}
