package options

type FillHistoryOptions struct {
	OrderID    string     `json:"orderId,omitempty"`
	Symbol     string     `json:"symbol,omitempty"`
	FillType   FillType   `json:"fillType,omitempty"`
	MarketType MarketType `json:"marketType,omitempty"`
	DateFilter
}

type FillType string

const (
	FillTypeUser                                   FillType = "User"
	FillTypeBookLiquidation                        FillType = "BookLiquidation"
	FillTypeAdl                                    FillType = "Adl"
	FillTypeBackstop                               FillType = "Backstop"
	FillTypeLiquidation                            FillType = "Liquidation"
	FillTypeAllLiquidation                         FillType = "AllLiquidation"
	FillTypeCollateralConversion                   FillType = "CollateralConversion"
	FillTypeCollateralConversionAndSpotLiquidation FillType = "CollateralConversionAndSpotLiquidation"
)

type FundingHistoryOptions struct {
	SubAccountID *int    `json:"subaccountId"`
	Symbol       *string `json:"symbol"`
	LimitOffset
}

type OrderHistoryOptions struct {
	OrderID    *string     `json:"orderId"`
	Symbol     *string     `json:"symbol"`
	MarketType *MarketType `json:"marketType"`
	LimitOffset
}

type PnlHistoryOptions struct {
	SubAccountID *int    `json:"subaccountId"`
	Symbol       *string `json:"symbol"`
	LimitOffset
}

type SettlementHistoryOptions struct {
	Source *string `json:"source"`
	LimitOffset
}
