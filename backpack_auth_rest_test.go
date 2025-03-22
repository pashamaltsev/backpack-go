package backpackgo

import (
	"fmt"
	"os"
	"testing"

	"github.com/feeeei/backpack-go/models"
	"github.com/feeeei/backpack-go/rest"
	"github.com/feeeei/backpack-go/utils"
)

func TestBackpackAuthREST(t *testing.T) {
	rest := NewRESTClient(rest.WithAPIToken(os.Getenv("API_KEY"), os.Getenv("API_SECRET")))

	symbol := "BTC_USDC"
	asset := "USDC"
	blockchain := "Solana"

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
		orderLimit, err := rest.GetAccountMaxOrder(symbol, models.Ask)
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

	// test ExecuteBorrowLend
	t.Run("test ExecuteBorrowLend", func(t *testing.T) {
		err := rest.ExecuteBorrowLend(asset, models.Lend, 1)
		if err != nil {
			t.Errorf("ExecuteBorrowLend failed: %v", err)
		} else {
			fmt.Printf("OK: ExecuteBorrowLend\n\n")
		}
	})

	// test GetBalance
	t.Run("test GetBalance", func(t *testing.T) {
		balance, err := rest.GetBalance()
		if err != nil {
			t.Errorf("GetBalance failed: %v", err)
		} else {
			fmt.Printf("OK: GetBalance, balance: %+v\n\n", balance["USDC"].Available)
		}
	})

	// test GetCollateral
	t.Run("test GetCollateral", func(t *testing.T) {
		collateral, err := rest.GetAccountCollateral()
		if err != nil {
			t.Errorf("GetCollateral failed: %v", err)
		} else {
			fmt.Printf("OK: GetCollateral, collateral count: %+v\n\n", len(collateral.Collateral))
		}
	})

	// test GetDeposits
	t.Run("test GetDeposits", func(t *testing.T) {
		deposits, err := rest.GetDeposits()
		if err != nil {
			t.Errorf("GetDeposits failed: %v", err)
		} else {
			fmt.Printf("OK: GetDeposits, deposits count: %+v\n\n", len(deposits))
		}
	})

	// test GetDepositAddress
	t.Run("test GetDepositAddress", func(t *testing.T) {
		address, err := rest.GetDepositAddress(blockchain)
		if err != nil {
			t.Errorf("GetDepositAddress failed: %v", err)
		} else {
			fmt.Printf("OK: GetDepositAddress, address: %+v\n\n", address)
		}
	})

	// test GetWithdrawals
	t.Run("test GetWithdrawals", func(t *testing.T) {
		withdrawals, err := rest.GetWithdrawals()
		if err != nil {
			t.Errorf("GetWithdrawals failed: %v", err)
		} else {
			fmt.Printf("OK: GetWithdrawals, withdrawals count: %+v\n\n", len(withdrawals))
		}
	})

	// test RequestWithdrawal
	t.Run("test Withdraw", func(t *testing.T) {
		withdrawal, err := rest.RequestWithdrawal(asset, 100, "6mi1SR95VbyKCzSymR85d2yikLcdLWR6YRWnsuv12hVC", blockchain)
		if err != nil {
			t.Errorf("Withdraw failed: %v", err)
		} else {
			fmt.Printf("OK: Withdraw, withdrawal: %+v\n\n", withdrawal)
		}
	})

	// test GetPositions
	t.Run("test GetPositions", func(t *testing.T) {
		positions, err := rest.GetPositions()
		if err != nil {
			t.Errorf("GetPositions failed: %v", err)
		} else {
			fmt.Printf("OK: GetPositions, positions count: %+v\n\n", len(positions))
		}
	})

	// test GetBorrowLendHistory
	t.Run("test GetBorrowLendHistory", func(t *testing.T) {
		history, err := rest.GetBorrowLendHistory()
		if err != nil {
			t.Errorf("GetBorrowLendHistory failed: %v", err)
		} else {
			fmt.Printf("OK: GetBorrowLendHistory, history count: %+v\n\n", len(history))
		}
	})

	// test GetInterestHistory
	t.Run("test GetInterestHistory", func(t *testing.T) {
		history, err := rest.GetInterestHistory()
		if err != nil {
			t.Errorf("GetInterestHistory failed: %v", err)
		} else {
			fmt.Printf("OK: GetInterestHistory, history count: %+v\n\n", len(history))
		}
	})

	// TODO: waiting for API stable
	// // test GetBorrowPositionsHistory
	// t.Run("test GetBorrowPositionsHistory", func(t *testing.T) {
	// 	history, err := rest.GetBorrowPositionsHistory()
	// 	if err != nil {
	// 		t.Errorf("GetBorrowPositionsHistory failed: %v", err)
	// 	} else {
	// 		fmt.Printf("OK: GetBorrowPositionsHistory, history count: %+v\n\n", len(history))
	// 	}
	// })

	// test GetFillHistory
	t.Run("test GetFillHistory", func(t *testing.T) {
		history, err := rest.GetFillHistory()
		if err != nil {
			t.Errorf("GetFillHistory failed: %v", err)
		} else {
			fmt.Printf("OK: GetFillHistory, history count: %+v\n\n", len(history))
		}
	})

	// test GetFundingHistory
	t.Run("test GetFundingHistory", func(t *testing.T) {
		history, err := rest.GetFundingHistory()
		if err != nil {
			t.Errorf("GetFundingHistory failed: %v", err)
		} else {
			fmt.Printf("OK: GetFundingHistory, history count: %+v\n\n", len(history))
		}
	})

	// test GetOrdersHistory
	t.Run("test GetOrdersHistory", func(t *testing.T) {
		history, err := rest.GetOrdersHistory()
		if err != nil {
			t.Errorf("GetOrdersHistory failed: %v", err)
		} else {
			fmt.Printf("OK: GetOrdersHistory, history count: %+v\n\n", len(history))
		}
	})

	// test GetPnlHistory
	t.Run("test GetPnlHistory", func(t *testing.T) {
		history, err := rest.GetPnlHistory()
		if err != nil {
			t.Errorf("GetPnlHistory failed: %v", err)
		} else {
			fmt.Printf("OK: GetPnlHistory, history count: %+v\n\n", len(history))
		}
	})

	// test GetSettlementHistory
	t.Run("test GetSettlementHistory", func(t *testing.T) {
		history, err := rest.GetSettlementHistory()
		if err != nil {
			t.Errorf("GetSettlementHistory failed: %v", err)
		} else {
			fmt.Printf("OK: GetSettlementHistory, history count: %+v\n\n", len(history))
		}
	})
}
