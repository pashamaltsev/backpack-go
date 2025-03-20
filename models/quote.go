package models

type Quote struct {
	RfqID    string      `json:"rfqId"`
	QuoteID  string      `json:"quoteId"`
	ClientID *int        `json:"clientId"`
	Price    float64     `json:"price,string"`
	Status   QuoteStatus `json:"status"`
}

type QuoteStatus string

const (
	QuoteStatusCancelled       QuoteStatus = "Cancelled"
	QuoteStatusExpired         QuoteStatus = "Expired"
	QuoteStatusFilled          QuoteStatus = "Filled"
	QuoteStatusNew             QuoteStatus = "New"
	QuoteStatusPartiallyFilled QuoteStatus = "PartiallyFilled"
	QuoteStatusTriggerPending  QuoteStatus = "TriggerPending"
	QuoteStatusTriggerFailed   QuoteStatus = "TriggerFailed"
)
