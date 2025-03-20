package models

import "time"

type LimitOffset struct {
	Limit  *int `json:"limit,string"`
	Offset *int `json:"offset,string"`
}

type Trade struct {
	ID            int64     `json:"id"`
	Price         float64   `json:"price,string"`
	Quantity      float64   `json:"quantity,string"`
	QuoteQuantity float64   `json:"quoteQuantity,string"`
	Timestamp     time.Time `json:"timestamp" time_format:"unixmilli"`
	IsBuyerMaker  bool      `json:"isBuyerMaker"`
}
