package models

type MarketAssets struct {
	Symbol string       `json:"symbol"`
	Tokens []AssetToken `json:"tokens"`
}

type AssetToken struct {
	Blockchain        string  `json:"blockchain"`
	ContractAddress   *string `json:"contractAddress"`
	DepositEnabled    bool    `json:"depositEnabled"`
	MaximumWithdrawal float64 `json:"maximumWithdrawal,string"`
	MinimumDeposit    float64 `json:"minimumDeposit,string"`
	MinimumWithdrawal float64 `json:"minimumWithdrawal,string"`
	WithdrawEnabled   bool    `json:"withdrawEnabled"`
	WithdrawalFee     float64 `json:"withdrawalFee,string"`
}

type MarketCollateral struct {
	Symbol          string             `json:"symbol"`
	IMFFunction     CollateralFunction `json:"imfFunction"`
	MMFFunction     CollateralFunction `json:"mmfFunction"`
	HaircutFunction HaircutFunction    `json:"haircutFunction"`
}

type CollateralFunction struct {
	Type   string  `json:"type"`
	Base   float64 `json:"base,string"`
	Factor float64 `json:"factor,string"`
}

type HaircutFunction struct {
	Weight float64 `json:"weight,string"`
	Kind   Kind    `json:"kind"`
}

type Kind struct {
	Type string `json:"type"`
}
