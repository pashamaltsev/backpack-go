package models

import "time"

type Trade struct {
	ID            int64     `json:"id"`
	Price         float64   `json:"price,string"`
	Quantity      float64   `json:"quantity,string"`
	QuoteQuantity float64   `json:"quoteQuantity,string"`
	Timestamp     time.Time `json:"timestamp,format:unixmilli"`
	IsBuyerMaker  bool      `json:"isBuyerMaker"`
}
