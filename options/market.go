package options

type TickerInterval string

const (
	OneDayTicker  TickerInterval = "1d"
	OneWeekTicker TickerInterval = "1w"
)

type MarketType string

const (
	SPOT       MarketType = "SPOT"
	PERP       MarketType = "PERP"
	IPERP      MarketType = "IPERP"
	DATED      MarketType = "DATED"
	PREDICTION MarketType = "PREDICTION"
	RFQ        MarketType = "RFQ"
)
