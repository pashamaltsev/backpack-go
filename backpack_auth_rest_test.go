package backpackgo

import (
	"fmt"
	"os"
	"testing"

	"github.com/feeeei/backpack-go/models"
	"github.com/feeeei/backpack-go/utils"
)

func TestBackpackAuthREST(t *testing.T) {
	rest := NewRESTClient(WithAPIToken(os.Getenv("API_KEY"), os.Getenv("API_SECRET")))

	symbol := "BTC_USDC"
	asset := "USDC"

	// test GetAccount
	t.Run("test GetAccount", func(t *testing.T) {
		account, err := rest.GetAccount()
		if err != nil {
			t.Errorf("GetAccount failed: %v", err)
		} else {
			fmt.Printf("OK: GetAccount, account.AutoLend: %+v\n\n", *account.AutoLend)
		}
	})

	// test UpdateAccount
	t.Run("test UpdateAccount", func(t *testing.T) {
		err := rest.UpdateAccount(&models.AccountUpdateble{
			AutoLend: utils.Ptr(false),
		})
		if err != nil {
			t.Errorf("UpdateAccount failed: %v", err)
		} else {
			fmt.Printf("OK: UpdateAccount\n\n")
		}
	})

	// test GetAccountMaxBorrow
	t.Run("test GetAccountMaxBorrow", func(t *testing.T) {
		borrowLimit, err := rest.GetAccountMaxBorrow(asset)
		if err != nil {
			t.Errorf("GetAccountMaxBorrow failed: %v", err)
		} else {
			fmt.Printf("OK: GetAccountMaxBorrow, max borrow quantity: %+v\n\n", borrowLimit.MaxBorrowQuantity)
		}
	})

	// test GetAccountMaxOrder
	t.Run("test GetAccountMaxOrder", func(t *testing.T) {
		orderLimit, err := rest.GetAccountMaxOrder(symbol, models.SideAsk)
		if err != nil {
			t.Errorf("GetAccountMaxOrder failed: %v", err)
		} else {
			fmt.Printf("OK: GetAccountMaxOrder, max order quantity: %+v\n\n", orderLimit.MaxOrderQuantity)
		}
	})

	// test GetAccountMaxWithdrawal
	t.Run("test GetAccountMaxWithdrawal", func(t *testing.T) {
		withdrawalLimit, err := rest.GetAccountMaxWithdrawal(asset, models.AccountWithdrawalLimitOptions{
			AutoBorrow: utils.Ptr(true),
		})
		if err != nil {
			t.Errorf("GetAccountMaxWithdrawal failed: %v", err)
		} else {
			fmt.Printf("OK: GetAccountMaxWithdrawal, max withdrawal quantity: %+v\n\n", withdrawalLimit.MaxWithdrawalQuantity)
		}
	})

	// test GetBorrowLendPositions
	t.Run("test GetBorrowLendPositions", func(t *testing.T) {
		positions, err := rest.GetBorrowLendPositions()
		if err != nil {
			t.Errorf("GetBorrowLendPositions failed: %v", err)
		} else {
			fmt.Printf("OK: GetBorrowLendPositions, positions count: %d\n\n", len(positions))
		}
	})
}
