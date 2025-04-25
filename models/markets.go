package models

import (
	"net/http"
	"strconv"

	"github.com/UnipayFI/backpack-go/options"
	json "github.com/json-iterator/go"

	"time"
)

type Market struct {
	Symbol            string             `json:"symbol"`
	BaseSymbol        string             `json:"baseSymbol"`
	QuoteSymbol       string             `json:"quoteSymbol"`
	MarketType        options.MarketType `json:"marketType"`
	Filters           OrderBookFilters   `json:"filters"`
	IMFFunction       Function           `json:"imfFunction"`
	MMFFunction       Function           `json:"mmfFunction"`
	FundingInterval   *int64             `json:"fundingInterval"`
	OpenInterestLimit *float64           `json:"openInterestLimit,string"`
	OrderBookState    OrderBookState     `json:"orderBookState"`
	CreatedAt         time.Time          `json:"createdAt" time_format:"2006-01-02T15:04:05.999999"`
}

type OrderBookFilters struct {
	Price    PriceFilter    `json:"price"`
	Quantity QuantityFilter `json:"quantity"`
}

type PriceFilter struct {
	MinPrice                    float64               `json:"minPrice,string"`
	MaxPrice                    *float64              `json:"maxPrice,string"`
	TickSize                    float64               `json:"tickSize,string"`
	MaxMultiplier               *float64              `json:"maxMultiplier,string"`
	MinMultiplier               *float64              `json:"minMultiplier,string"`
	MaxImpactMultiplier         *float64              `json:"maxImpactMultiplier,string"`
	MinImpactMultiplier         *float64              `json:"minImpactMultiplier,string"`
	MeanMarkPriceBand           *PriceBandMarkPrice   `json:"meanMarkPriceBand"`
	MeanPremiumBand             *PriceBandMeanPremium `json:"meanPremiumBand"`
	BorrowEntryFeeMaxMultiplier *float64              `json:"borrowEntryFeeMaxMultiplier,string"`
	BorrowEntryFeeMinMultiplier *float64              `json:"borrowEntryFeeMinMultiplier,string"`
}

type QuantityFilter struct {
	MinQuantity float64  `json:"minQuantity,string"`
	MaxQuantity *float64 `json:"maxQuantity,string"`
	StepSize    float64  `json:"stepSize,string"`
}

type PriceBandMarkPrice struct {
	MaxMultiplier float64 `json:"maxMultiplier,string"`
	MinMultiplier float64 `json:"minMultiplier,string"`
}

type PriceBandMeanPremium struct {
	TolerancePct float64 `json:"tolerancePct,string"`
}

type Function struct {
	Type   string  `json:"type"`
	Base   float64 `json:"base,string"`
	Factor float64 `json:"factor,string"`
}

type OrderBookState string

const (
	OrderBookStateOpen       OrderBookState = "Open"
	OrderBookStateClosed     OrderBookState = "Closed"
	OrderBookStateCancelOnly OrderBookState = "CancelOnly"
	OrderBookStateLimitOnly  OrderBookState = "LimitOnly"
	OrderBookStatePostOnly   OrderBookState = "PostOnly"
)

type Ticker struct {
	Symbol             string  `json:"symbol"`
	FirstPrice         float64 `json:"firstPrice,string"`
	LastPrice          float64 `json:"lastPrice,string"`
	PriceChange        float64 `json:"priceChange,string"`
	PriceChangePercent float64 `json:"priceChangePercent,string"`
	High               float64 `json:"high,string"`
	Low                float64 `json:"low,string"`
	Volume             float64 `json:"volume,string"`
	QuoteVolume        float64 `json:"quoteVolume,string"`
	Trades             int     `json:"trades,string"`
}

type Depth struct {
	Asks         []DepthItem `json:"asks"`
	Bids         []DepthItem `json:"bids"`
	LastUpdateID string      `json:"lastUpdateId"`
	Timestamp    time.Time   `json:"timestamp" time_format:"unixmicro"`
}

type DepthItem [2]float64

func (di *DepthItem) UnmarshalJSON(data []byte) error {
	var arr [2]string
	if err := json.Unmarshal(data, &arr); err != nil {
		return err
	}

	var err error
	(*di)[0], err = strconv.ParseFloat(arr[0], 64)
	if err != nil {
		return err
	}

	(*di)[1], err = strconv.ParseFloat(arr[1], 64)
	if err != nil {
		return err
	}
	return nil
}

type Kline struct {
	Start       time.Time `json:"start" time_format:"DateTime"`
	End         time.Time `json:"end" time_format:"DateTime"`
	Open        *float64  `json:"open,string"`
	High        *float64  `json:"high,string"`
	Low         *float64  `json:"low,string"`
	Close       *float64  `json:"close,string"`
	Volume      *float64  `json:"volume,string"`
	QuoteVolume *float64  `json:"quoteVolume,string"`
	Trades      *int      `json:"trades,string"`
}

type MarkPrice struct {
	FundingRate          float64   `json:"fundingRate,string"`
	IndexPrice           float64   `json:"indexPrice,string"`
	MarkPrice            float64   `json:"markPrice,string"`
	NextFundingTimestamp time.Time `json:"nextFundingTimestamp" time_format:"unixmilli"`
	Symbol               string    `json:"symbol"`
}

type OpenInterest struct {
	Symbol       string    `json:"symbol"`
	OpenInterest float64   `json:"openInterest,string"`
	Timestamp    time.Time `json:"timestamp" time_format:"unixmilli"`
}

type PageHeaders struct {
	AccessControlExposeHeaders string `header:"ACCESS-CONTROL-EXPOSE-HEADERS"`
	PageCount                  int64  `header:"X-PAGE-COUNT"`
	CurrentPage                int64  `header:"X-CURRENT-PAGE"`
	PageSize                   int64  `header:"X-PAGE-SIZE"`
	Total                      int64  `header:"X-TOTAL"`
}

func ParseFundingRateHeaders(headers http.Header) *PageHeaders {
	header := &PageHeaders{}

	if exposeHeaders := headers.Get("ACCESS-CONTROL-EXPOSE-HEADERS"); exposeHeaders != "" {
		header.AccessControlExposeHeaders = exposeHeaders
	}

	if pageCount := headers.Get("X-PAGE-COUNT"); pageCount != "" {
		if val, err := strconv.ParseInt(pageCount, 10, 64); err == nil {
			header.PageCount = val
		}
	}

	if currentPage := headers.Get("X-CURRENT-PAGE"); currentPage != "" {
		if val, err := strconv.ParseInt(currentPage, 10, 64); err == nil {
			header.CurrentPage = val
		}
	}

	if pageSize := headers.Get("X-PAGE-SIZE"); pageSize != "" {
		if val, err := strconv.ParseInt(pageSize, 10, 64); err == nil {
			header.PageSize = val
		}
	}

	if total := headers.Get("X-TOTAL"); total != "" {
		if val, err := strconv.ParseInt(total, 10, 64); err == nil {
			header.Total = val
		}
	}

	return header
}

type FundingRate struct {
	Symbol               string    `json:"symbol"`
	IntervalEndTimestamp time.Time `json:"intervalEndTimestamp" time_format:"2006-01-02T15:04:05"`
	FundingRate          float64   `json:"fundingRate,string"`
}
