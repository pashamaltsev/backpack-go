package options

type BorrowLendSide string

const (
	Borrow BorrowLendSide = "Borrow"
	Lend   BorrowLendSide = "Lend"
)

type BorrowLendType string

const (
	BorrowLendEventTypeBorrow      BorrowLendType = "Borrow"
	BorrowLendEventTypeBorrowRepay BorrowLendType = "BorrowRepay"
	BorrowLendEventTypeLend        BorrowLendType = "Lend"
	BorrowLendEventTypeLendRedeem  BorrowLendType = "LendRedeem"
)

type BorrowHistoryOptions struct {
	Type       *BorrowLendType `json:"type"`
	Sources    *string         `json:"sources"`
	PositionID *string         `json:"positionId"`
	Symbol     *int64          `json:"symbol"`
	LimitOffset
}

type BorrowPostionHistoryOptions struct {
	Symbol *string          `json:"symbol"`
	Side   *BorrowLendSide  `json:"side"`
	State  *BorrowLendState `json:"state"`
	LimitOffset
}

type BorrowLendState string

const (
	BorrowLendStateOpen   BorrowLendState = "Open"
	BorrowLendStateClosed BorrowLendState = "Closed"
)
