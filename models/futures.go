package models

type Position struct {
	BreakEvenPrice           float64  `json:"breakEvenPrice,string"`
	EntryPrice               float64  `json:"entryPrice,string"`
	EstLiquidationPrice      float64  `json:"estLiquidationPrice,string"`
	Imf                      float64  `json:"imf,string"`
	ImfFunction              Function `json:"imfFunction"`
	MarkPrice                float64  `json:"markPrice,string"`
	Mmf                      float64  `json:"mmf,string"`
	MmfFunction              Function `json:"mmfFunction"`
	NetCost                  float64  `json:"netCost,string"`
	NetQuantity              float64  `json:"netQuantity,string"`
	NetExposureQuantity      float64  `json:"netExposureQuantity,string"`
	NetExposureNotional      float64  `json:"netExposureNotional,string"`
	PnlRealized              float64  `json:"pnlRealized,string"`
	PnlUnrealized            float64  `json:"pnlUnrealized,string"`
	CumulativeFundingPayment float64  `json:"cumulativeFundingPayment,string"`
	SubAccountID             int      `json:"subaccountId"`
	Symbol                   string   `json:"symbol"`
	UserID                   int      `json:"userId"`
	PositionID               string   `json:"positionId"`
	CumulativeInterest       float64  `json:"cumulativeInterest,string"`
}
