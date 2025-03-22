package options

type InterestHistoryOptions struct {
	Asset      *string                `json:"asset"`
	Symbol     *string                `json:"symbol"`
	PositionID *string                `json:"positionId"`
	Source     *InterestHistorySource `json:"source"`
	LimitOffset
}

type InterestHistorySource string

const (
	InterestHistorySourceUnrealizedPnl InterestHistorySource = "UnrealizedPnl"
	InterestHistorySourceBorrowLend    InterestHistorySource = "BorrowLend"
)
