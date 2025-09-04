package backpackgo

import (
	"fmt"
	"testing"
	"time"

	"github.com/pashamaltsev/backpack-go/options"
)

func TestBackpackPublicREST(t *testing.T) {
	rest := NewRESTClient()

	symbol := "BTC_USDC"

	// test GetAssets
	t.Run("test GetAssets", func(t *testing.T) {
		assets, err := rest.GetMarketAssets()
		if err != nil {
			t.Errorf("GetAssets failed: %v", err)
		} else {
			fmt.Printf("OK: GetAssets, Assets count: %d\n\n", len(assets))
		}
	})

	// test GetCollateral
	t.Run("test GetCollateral", func(t *testing.T) {
		collaterals, err := rest.GetMarketCollateral()
		if err != nil {
			t.Errorf("GetCollateral failed: %v", err)
		} else {
			fmt.Printf("OK: GetCollateral, Collateral count: %d\n\n", len(collaterals))
		}
	})

	// test GetBorrowLendMarkets
	t.Run("test GetBorrowLendMarkets", func(t *testing.T) {
		markets, err := rest.GetBorrowLendMarkets()
		if err != nil {
			t.Errorf("GetBorrowLendMarkets failed: %v", err)
		} else {
			fmt.Printf("OK: GetBorrowLendMarkets, BorrowLendMarkets count: %d\n\n", len(markets))
		}
	})

	// test GetBorrowLendMarketsHistory
	t.Run("test GetBorrowLendMarketsHistory", func(t *testing.T) {
		history, err := rest.GetBorrowLendMarketsHistory(options.OneDay)
		if err != nil {
			t.Errorf("GetBorrowLendMarketsHistory failed: %v", err)
		} else {
			fmt.Printf("OK: GetBorrowLendMarketsHistory, BorrowLendMarketsHistory count: %d\n\n", len(history))
		}
	})

	// test GetMarkets
	t.Run("test GetMarkets", func(t *testing.T) {
		markets, err := rest.GetMarkets()
		if err != nil {
			t.Errorf("GetMarkets failed: %v", err)
		} else {
			fmt.Printf("OK: GetMarkets, Markets count: %d\n\n", len(markets))
		}
	})

	// test GetMarket
	t.Run("test GetMarket", func(t *testing.T) {
		market, err := rest.GetMarket(symbol)
		if err != nil {
			t.Errorf("GetMarket failed: %v", err)
		} else {
			fmt.Printf("OK: GetMarket, %s Market orderbook state: %+v\n\n", market.Symbol, market.OrderBookState)
		}
	})

	// test GetTicker
	t.Run("test GetTicker", func(t *testing.T) {
		ticker, err := rest.GetTicker(symbol)
		if err != nil {
			t.Errorf("GetTicker failed: %v", err)
		} else {
			fmt.Printf("OK: GetTicker, %s price: %+v\n\n", ticker.Symbol, ticker.FirstPrice)
		}
	})

	// test GetDepth
	t.Run("test GetDepth", func(t *testing.T) {
		depth, err := rest.GetDepth(symbol)
		if err != nil {
			t.Errorf("GetDepth failed: %v", err)
		} else {
			fmt.Printf("OK: GetDepth, depth last update id: %s\n\n", depth.LastUpdateID)
		}
	})

	// test GetKlines
	t.Run("test GetKlines", func(t *testing.T) {
		endTime := time.Now()
		startTime := endTime.Add(-24 * time.Hour) // 过去24小时
		klines, err := rest.GetKlines(symbol, options.KLineInterval1h, startTime, endTime)
		if err != nil {
			t.Errorf("GetKlines failed: %v", err)
		} else {
			fmt.Printf("OK: GetKlines, klines count: %d\n\n", len(klines))
		}
	})

	// test GetOpenInterest
	t.Run("test GetOpenInterest", func(t *testing.T) {
		interests, err := rest.GetOpenInterest()
		if err != nil {
			t.Errorf("GetOpenInterest failed: %v", err)
		} else {
			fmt.Printf("OK: GetOpenInterest, open interest count: %+v\n\n", len(interests))
		}
	})

	// test GetFundingRates
	t.Run("test GetFundingRates", func(t *testing.T) {
		headers, rates, err := rest.GetFundingRates(symbol + "_PERP")
		if err != nil {
			t.Errorf("GetFundingRates failed: %v", err)
		} else {
			fmt.Printf("OK: GetFundingRates, total: %+v, count: %d\n\n", headers.Total, len(rates))
		}
	})

	// test GetTickers
	t.Run("test GetTickers", func(t *testing.T) {
		tickers, err := rest.GetTickers()
		if err != nil {
			t.Errorf("GetTickers failed: %v", err)
		} else {
			fmt.Printf("OK: GetTickers, tickers count: %d\n\n", len(tickers))
		}
	})

	// test GetMarkPrices
	t.Run("test GetMarkPrices", func(t *testing.T) {
		prices, err := rest.GetMarkPrices()
		if err != nil {
			t.Errorf("GetMarkPrices failed: %v", err)
		} else {
			fmt.Printf("OK: GetMarkPrices, mark prices count: %d\n\n", len(prices))
		}
	})

	// test Ping
	t.Run("test Ping", func(t *testing.T) {
		err := rest.Ping()
		if err != nil {
			t.Errorf("Ping failed: %v", err)
		} else {
			fmt.Printf("OK: Ping\n\n")
		}
	})

	// test GetTime
	t.Run("test GetTime", func(t *testing.T) {
		time, err := rest.GetTime()
		if err != nil {
			t.Errorf("GetTime failed: %v", err)
		} else {
			fmt.Printf("OK: GetTime, Server time: %v\n\n", time)
		}
	})

	// test GetStatus
	t.Run("test GetStatus", func(t *testing.T) {
		status, err := rest.GetStatus()
		if err != nil {
			t.Errorf("GetStatus failed: %v", err)
		} else {
			fmt.Printf("OK: GetStatus, Status: %+v\n\n", status)
		}
	})

	// test GetTrades
	t.Run("test GetTrades", func(t *testing.T) {
		trades, err := rest.GetTrades(symbol, options.LimitOffset{Limit: 100})
		if err != nil {
			t.Errorf("GetTrades failed: %v", err)
		} else {
			fmt.Printf("OK: GetTrades, trades count: %d\n\n", len(trades))
		}
	})

	// test GetTradesHistory
	t.Run("test GetTradesHistory", func(t *testing.T) {
		trades, err := rest.GetTradesHistory(symbol, 100)
		if err != nil {
			t.Errorf("GetTradesHistory failed: %v", err)
		} else {
			fmt.Printf("OK: GetTradesHistory, trades count: %d\n\n", len(trades))
		}
	})
}
