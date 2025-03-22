package options

type WithdrawalOptions struct {
	ClientID       *string `json:"clientId"`
	TwoFactorCode  *string `json:"twoFactorCode"`
	AutoBorrow     *bool   `json:"autoBorrow"`
	AutoLendRedeem *bool   `json:"autoLendRedeem"`
}
