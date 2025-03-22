package backpackgo

import (
	"fmt"
	"os"
	"testing"

	"github.com/feeeei/backpack-go/models"
	"github.com/feeeei/backpack-go/options"
	"github.com/feeeei/backpack-go/rest"
)

func TestBackpackOrderREST(t *testing.T) {
	rest := NewRESTClient(rest.WithAPIToken(os.Getenv("API_KEY"), os.Getenv("API_SECRET")))

	symbol := "SOL_USDC"

	var order *models.Order
	// test ExecuteOrder
	t.Run("test ExecuteOrder", func(t *testing.T) {
		var err error
		order, err = rest.ExecuteLimitOrder(symbol, options.Buy, 1, 10)
		if err != nil {
			t.Errorf("ExecuteOrder failed: %v", err)
		} else {
			fmt.Printf("OK: ExecuteOrder, order id: %+v\n\n", order.ID)
		}
	})

	// test GetOrder
	t.Run("test GetOrder", func(t *testing.T) {
		order, err := rest.GetOrderByOrderID(order.Symbol, order.ID)
		if err != nil {
			t.Errorf("GetOrder failed: %v", err)
		} else {
			fmt.Printf("OK: GetOrder, order id: %+v\n\n", order.ID)
		}
	})

	// test GetOrders
	t.Run("test GetOrders", func(t *testing.T) {
		orders, err := rest.GetOrders(nil, nil)
		if err != nil {
			t.Errorf("GetOrders failed: %v", err)
		} else {
			fmt.Printf("OK: GetOrders, orders count: %+v\n\n", len(orders))
		}
	})

	// test CancelOrder
	t.Run("test CancelOrder", func(t *testing.T) {
		order, err := rest.CancelOrderByOrderID(order.Symbol, order.ID)
		if err != nil {
			t.Errorf("CancelOrder failed: %v", err)
		} else {
			fmt.Printf("OK: CancelOrder, order id: %+v\n\n", order.ID)
		}
	})

	// test CancelAllOrders
	t.Run("test CancelAllOrders", func(t *testing.T) {
		orders, err := rest.CancelOrders(symbol)
		if err != nil {
			t.Errorf("CancelAllOrders failed: %v", err)
		} else {
			fmt.Printf("OK: CancelAllOrders, orders count: %+v\n\n", len(orders))
		}
	})
}
