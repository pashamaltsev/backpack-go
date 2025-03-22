package options

import "time"

type LimitOffset struct {
	Limit  *int `json:"limit,string"`
	Offset *int `json:"offset,string"`
}

type DateFilter struct {
	From *time.Time `json:"from" time_format:"unixmilli"`
	To   *time.Time `json:"to" time_format:"unixmilli"`
	LimitOffset
}
