package backpackgo

import (
	"fmt"
	"os"
	"testing"

	"github.com/feeeei/backpack-go/options"
	"github.com/feeeei/backpack-go/rest"
)

func TestBackpackExecuteOrder(t *testing.T) {
	rest := NewRESTClient(rest.WithAPIToken(os.Getenv("API_KEY"), os.Getenv("API_SECRET")))

	symbol := "SOL_USDC"

	// test ExecuteMarketOrder
	t.Run("test ExecuteMarketOrder", func(t *testing.T) {
		var err error
		order, err := rest.ExecuteMarketOrder(symbol+"_PERP", options.Buy, 0.1, options.WithAutoLendRedeem(true), options.WithTimeInForce(options.TimeInForceFOK))
		if err != nil {
			t.Errorf("ExecuteOrder failed: %v", err)
		} else {
			fmt.Printf("OK: ExecuteOrder, order id: %+v\n\n", order.ID)
		}
	})

	// test ExecuteLimitOrder
	t.Run("test basic ExecuteLimitOrder", func(t *testing.T) {
		order, err := rest.ExecuteLimitOrder(
			symbol,
			options.Sell,
			500,
			0.11,
			options.WithAutoLendRedeem(true),
		)
		if err != nil {
			t.Errorf("ExecuteOrder failed: %v", err)
		} else {
			rest.CancelOrderByOrderID(symbol, order.ID)
			fmt.Printf("OK: ExecuteOrder, order id: %+v\n\n", order.ID)
		}
	})

	// test ExecuteConditionalLimitOrder
	t.Run("test ExecuteConditionalLimitOrder", func(t *testing.T) {
		order, err := rest.ExecuteConditionalLimitOrder(
			symbol,
			options.Sell,
			500,
			490,
			0.1,
			options.WithAutoLendRedeem(true),
		)
		if err != nil {
			t.Errorf("ExecuteOrder failed: %v", err)
		} else {
			rest.CancelOrderByOrderID(symbol, order.ID)
			fmt.Printf("OK: ExecuteOrder, order id: %+v\n\n", order.ID)
		}
	})
}
