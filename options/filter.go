package options

import "time"

type LimitOffset struct {
	Limit  int `json:"limit,omitzero,string"`
	Offset int `json:"offset,omitzero,string"`
}

type DateFilter struct {
	From time.Time `json:"from,omitzero,format:unixmilli" `
	To   time.Time `json:"to,omitzero,format:unixmilli"`
	LimitOffset
}
