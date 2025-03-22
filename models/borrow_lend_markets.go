package models

import "time"

type BorrowLendMarketState string

const (
	BorrowLendMarketStateOpen      BorrowLendMarketState = "Open"
	BorrowLendMarketStateClosed    BorrowLendMarketState = "Closed"
	BorrowLendMarketStateRepayOnly BorrowLendMarketState = "RepayOnly"
)

type BorrowLendMarket struct {
	State                        BorrowLendMarketState `json:"state"`
	AssetMarkPrice               float64               `json:"assetMarkPrice,string"`
	BorrowInterestRate           float64               `json:"borrowInterestRate,string"`
	BorrowedQuantity             float64               `json:"borrowedQuantity,string"`
	Fee                          float64               `json:"fee,string"`
	LendInterestRate             float64               `json:"lendInterestRate,string"`
	LentQuantity                 float64               `json:"lentQuantity,string"`
	MaxUtilization               float64               `json:"maxUtilization,string"`
	OpenBorrowLendLimit          float64               `json:"openBorrowLendLimit,string"`
	OptimalUtilization           float64               `json:"optimalUtilization,string"`
	Symbol                       string                `json:"symbol"`
	Timestamp                    time.Time             `json:"timestamp" time_format:"RFC3339Nano"`
	ThrottleUtilizationThreshold float64               `json:"throttleUtilizationThreshold,string"`
	ThrottleUtilizationBound     float64               `json:"throttleUtilizationBound,string"`
	ThrottleUpdateFraction       float64               `json:"throttleUpdateFraction,string"`
	Utilization                  float64               `json:"utilization,string"`
	StepSize                     float64               `json:"stepSize,string"`
}

type BorrowLendMarketHistory struct {
	BorrowInterestRate float64   `json:"borrowInterestRate,string"`
	BorrowedQuantity   float64   `json:"borrowedQuantity,string"`
	LendInterestRate   float64   `json:"lendInterestRate,string"`
	LentQuantity       float64   `json:"lentQuantity,string"`
	Timestamp          time.Time `json:"timestamp" time_format:"RFC3339Nano"`
	Utilization        float64   `json:"utilization,string"`
}
