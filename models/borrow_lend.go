package models

type BorrowLend struct {
	CumulativeInterest  float64  `json:"cumulativeInterest,string"`
	ID                  string   `json:"id"`
	Imf                 float64  `json:"imf,string"`
	ImfFunction         Function `json:"imfFunction"`
	NetQuantity         float64  `json:"netQuantity,string"`
	MarkPrice           float64  `json:"markPrice,string"`
	Mmf                 float64  `json:"mmf,string"`
	MmfFunction         Function `json:"mmfFunction"`
	NetExposureQuantity float64  `json:"netExposureQuantity,string"`
	NetExposureNotional float64  `json:"netExposureNotional,string"`
	Symbol              string   `json:"symbol"`
}

type BorrowLendSide string

const (
	BorrowLendSideBorrow BorrowLendSide = "Borrow"
	BorrowLendSideLend   BorrowLendSide = "Lend"
)
