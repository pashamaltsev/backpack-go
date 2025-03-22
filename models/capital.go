package models

import (
	"time"
)

type AssetBalance struct {
	Available float64 `json:"available,string"`
	Locked    float64 `json:"locked,string"`
	Staked    float64 `json:"staked,string"`
}

type Collateral struct {
	AssetsValue        float64           `json:"assetsValue,string"`
	BorrowLiability    float64           `json:"borrowLiability,string"`
	Collateral         []CollateralAsset `json:"collateral"`
	Imf                float64           `json:"imf,string"`
	UnsettledEquity    float64           `json:"unsettledEquity,string"`
	LiabilitiesValue   float64           `json:"liabilitiesValue,string"`
	MarginFraction     *float64          `json:"marginFraction,string"`
	Mmf                float64           `json:"mmf,string"`
	NetEquity          float64           `json:"netEquity,string"`
	NetEquityAvailable float64           `json:"netEquityAvailable,string"`
	NetEquityLocked    float64           `json:"netEquityLocked,string"`
	NetExposureFutures float64           `json:"netExposureFutures,string"`
	PnlUnrealized      float64           `json:"pnlUnrealized,string"`
	SubAccountID       *int              `json:"subaccountId"`
	UserID             int               `json:"userId"`
}

type CollateralAsset struct {
	Symbol            string  `json:"symbol"`
	AssetMarkPrice    float64 `json:"assetMarkPrice,string"`
	TotalQuantity     float64 `json:"totalQuantity,string"`
	BalanceNotional   float64 `json:"balanceNotional,string"`
	CollateralWeight  float64 `json:"collateralWeight,string"`
	CollateralValue   float64 `json:"collateralValue,string"`
	OpenOrderQuantity float64 `json:"openOrderQuantity,string"`
	LendQuantity      float64 `json:"lendQuantity,string"`
	AvailableQuantity float64 `json:"availableQuantity,string"`
}

type Deposit struct {
	ID                      int           `json:"id"`
	ToAddress               *string       `json:"toAddress"`
	FromAddress             *string       `json:"fromAddress"`
	ConfirmationBlockNumber *int64        `json:"confirmationBlockNumber"`
	Source                  DepositSource `json:"source"`
	Status                  string        `json:"status"`
	TransactionHash         *string       `json:"transactionHash"`
	Symbol                  string        `json:"symbol"`
	Quantity                string        `json:"quantity"`
	CreatedAt               time.Time     `json:"createdAt" time_format:"2006-01-02T15:04:05.000"`
}

type DepositSource string

const (
	Administrator DepositSource = "administrator"
	Arbitrum      DepositSource = "arbitrum"
	Base          DepositSource = "base"
	Bitcoin       DepositSource = "bitcoin"
	BitcoinCash   DepositSource = "bitcoinCash"
	Bsc           DepositSource = "bsc"
	Cardano       DepositSource = "cardano"
	Dogecoin      DepositSource = "dogecoin"
	Ethereum      DepositSource = "ethereum"
	Litecoin      DepositSource = "litecoin"
	Polygon       DepositSource = "polygon"
	Sui           DepositSource = "sui"
	Solana        DepositSource = "solana"
	Story         DepositSource = "story"
	Xrp           DepositSource = "xRP"
	EqualsMoney   DepositSource = "equalsMoney"
	Nuvei         DepositSource = "nuvei"
	Banxa         DepositSource = "banxa"
	Interac       DepositSource = "ioFinnet"
	Internal      DepositSource = "internal"
)

type DepositStatus string

const (
	DepositStatusCancelled DepositStatus = "cancelled"
	DepositStatusConfirmed DepositStatus = "confirmed"
	DepositStatusDeclined  DepositStatus = "declined"
	DepositStatusExpired   DepositStatus = "expired"
	DepositStatusInitiated DepositStatus = "initiated"
	DepositStatusPending   DepositStatus = "pending"
	DepositStatusRefunded  DepositStatus = "refunded"
)

type DepositAddress struct {
	Address string `json:"address"`
}

type Withdrawal struct {
	ID              int              `json:"id"`
	Blockchain      string           `json:"blockchain"`
	ClientID        *string          `json:"clientId"`
	Identifier      *string          `json:"identifier"`
	Quantity        float64          `json:"quantity,string"`
	Fee             float64          `json:"fee,string"`
	Symbol          string           `json:"symbol"`
	Status          WithdrawalStatus `json:"status"`
	SubAccountID    *int             `json:"subaccountId"`
	ToAddress       string           `json:"toAddress"`
	TransactionHash *string          `json:"transactionHash"`
	CreatedAt       time.Time        `json:"createdAt"  time_format:"2006-01-02T15:04:05.000"`
	IsInternal      bool             `json:"isInternal"`
}

type WithdrawalStatus string

const (
	WithdrawalStatusConfirmed WithdrawalStatus = "confirmed"
	WithdrawalStatusPending   WithdrawalStatus = "pending"
)
