package options

type AccountOrderLimitOptions struct {
	Price           *float64 `json:"price,string,omitempty"`
	ReduceOnly      *bool    `json:"reduceOnly,omitempty"`
	AutoBorrow      *bool    `json:"autoBorrow,omitempty"`
	AutoBorrowRepay *bool    `json:"autoBorrowRepay,omitempty"`
	AutoLendRedeem  *bool    `json:"autoLendRedeem,omitempty"`
}

type AccountWithdrawalLimitOptions struct {
	AutoBorrow     *bool `json:"autoBorrow,string,omitempty"`
	AutoLendRedeem *bool `json:"autoLendRedeem,string,omitempty"`
}
