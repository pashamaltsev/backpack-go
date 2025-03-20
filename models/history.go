package models

import "time"

type BorrowHistoryOptions struct {
	Type       *BorrowLendType `json:"type"`
	Sources    *string         `json:"sources"`
	PositionID *string         `json:"positionId"`
	Symbol     *int64          `json:"symbol"`
	LimitOffset
}

type BorrowLendHistory struct {
	EventType         BorrowLendType       `json:"eventType"`
	PositionID        string               `json:"positionId"`
	PositionQuantity  *float64             `json:"positionQuantity,string"`
	Quantity          *float64             `json:"quantity,string"`
	Source            BorrowLendSourceType `json:"source"`
	Symbol            string               `json:"symbol"`
	Timestamp         time.Time            `json:"timestamp" time_format:"unix"`
	SpotMarginOrderID string               `json:"spotMarginOrderId"`
}

type BorrowLendType string

const (
	BorrowLendEventTypeBorrow      BorrowLendType = "Borrow"
	BorrowLendEventTypeBorrowRepay BorrowLendType = "BorrowRepay"
	BorrowLendEventTypeLend        BorrowLendType = "Lend"
	BorrowLendEventTypeLendRedeem  BorrowLendType = "LendRedeem"
)

type BorrowLendSourceType string

const (
	BorrowLendSourceTypeAdlProvider         BorrowLendSourceType = "AdlProvider"
	BorrowLendSourceTypeAutoBorrowRepay     BorrowLendSourceType = "AutoBorrowRepay"
	BorrowLendSourceTypeAutoLend            BorrowLendSourceType = "AutoLend"
	BorrowLendSourceTypeBackstopProvider    BorrowLendSourceType = "BackstopProvider"
	BorrowLendSourceTypeInterest            BorrowLendSourceType = "Interest"
	BorrowLendSourceTypeLiquidation         BorrowLendSourceType = "Liquidation"
	BorrowLendSourceTypeLiquidationAdl      BorrowLendSourceType = "LiquidationAdl"
	BorrowLendSourceTypeLiquidationBackstop BorrowLendSourceType = "LiquidationBackstop"
	BorrowLendSourceTypeManual              BorrowLendSourceType = "Manual"
	BorrowLendSourceTypeReconciliation      BorrowLendSourceType = "Reconciliation"
	BorrowLendSourceTypeSpotMargin          BorrowLendSourceType = "SpotMargin"
	BorrowLendSourceTypeWithdrawal          BorrowLendSourceType = "Withdrawal"
)

type InterestHistoryOptions struct {
	Asset      *string                `json:"asset"`
	Symbol     *string                `json:"symbol"`
	PositionID *string                `json:"positionId"`
	Source     *InterestHistorySource `json:"source"`
	LimitOffset
}

type InterestHistorySource string

const (
	InterestHistorySourceUnrealizedPnl InterestHistorySource = "UnrealizedPnl"
	InterestHistorySourceBorrowLend    InterestHistorySource = "BorrowLend"
)

type InterestHistory struct {
	PaymentType  InterestHistoryPaymentType `json:"paymentType"`
	InterestRate float64                    `json:"interestRate,string"`
	Interval     int64                      `json:"interval"`
	MarketSymbol string                     `json:"marketSymbol"`
	PositionID   string                     `json:"positionId"`
	Quantity     float64                    `json:"quantity,string"`
	Symbol       string                     `json:"symbol"`
	Timestamp    time.Time                  `json:"timestamp" time_format:"2006-01-02T15:04:05.000Z"`
}

type InterestHistoryPaymentType string

const (
	InterestHistoryPaymentTypeEntryFee              InterestHistoryPaymentType = "EntryFee"
	InterestHistoryPaymentTypeBorrow                InterestHistoryPaymentType = "Borrow"
	InterestHistoryPaymentTypeLend                  InterestHistoryPaymentType = "Lend"
	InterestHistoryPaymentTypeUnrealizedPositivePnl InterestHistoryPaymentType = "UnrealizedPositivePnl"
	InterestHistoryPaymentTypeUnrealizedNegativePnl InterestHistoryPaymentType = "UnrealizedNegativePnl"
)

type BorrowPostionHistoryOptions struct {
	Symbol *string          `json:"symbol"`
	Side   *BorrowLendSide  `json:"side"`
	State  *BorrowLendState `json:"state"`
	LimitOffset
}

type BorrowLendState string

const (
	BorrowLendStateOpen   BorrowLendState = "Open"
	BorrowLendStateClosed BorrowLendState = "Closed"
)

type BorrowPositionHistory struct {
	PositionID         string               `json:"positionId"`
	Quantity           float64              `json:"quantity,string"`
	Symbol             string               `json:"symbol"`
	Source             BorrowLendSourceType `json:"source"`
	CumulativeInterest float64              `json:"cumulativeInterest,string"`
	AvgInterestRate    float64              `json:"avgInterestRate,string"`
	Side               BorrowLendSide       `json:"side"`
	CreatedAt          time.Time            `json:"createdAt" time_format:"2006-01-02T15:04:05.000Z"`
}

type FillHistoryOptions struct {
	OrderID    *string     `json:"orderId"`
	Symbol     *string     `json:"symbol"`
	FillType   *FillType   `json:"fillType"`
	MarketType *MarketType `json:"marketType"`
	DateFilter
}

type FillType string

const (
	FillTypeUser                                   FillType = "User"
	FillTypeBookLiquidation                        FillType = "BookLiquidation"
	FillTypeAdl                                    FillType = "Adl"
	FillTypeBackstop                               FillType = "Backstop"
	FillTypeLiquidation                            FillType = "Liquidation"
	FillTypeAllLiquidation                         FillType = "AllLiquidation"
	FillTypeCoLiquidation                          FillType = "CoLiquidation"
	FillTypeCollateralConversionAndSpotLiquidation FillType = "CollateralConversionAndSpotLiquidation"
)

type FillHistory struct {
	ClientID        *string             `json:"clientId"`
	Fee             float64             `json:"fee,string"`
	FeeSymbol       string              `json:"feeSymbol"`
	IsMaker         bool                `json:"isMaker"`
	OrderID         string              `json:"orderId"`
	Price           float64             `json:"price,string"`
	Quantity        float64             `json:"quantity,string"`
	Side            Side                `json:"side"`
	Symbol          string              `json:"symbol"`
	SystemOrderType FillSystemOrderType `json:"systemOrderType"`
	Timestamp       time.Time           `json:"timestamp" time_format:"2006-01-02T15:04:05.000Z"`
	TradeID         *int64              `json:"tradeId"`
}

type FillSystemOrderType string

const (
	FillSystemOrderTypeCollateralConversion        FillSystemOrderType = "CollateralConversion"
	FillSystemOrderTypeFutureExpiry                FillSystemOrderType = "FutureExpiry"
	FillSystemOrderTypeLiquidatePositionOnAdl      FillSystemOrderType = "LiquidatePositionOnAdl"
	FillSystemOrderTypeLiquidatePositionOnBook     FillSystemOrderType = "LiquidatePositionOnBook"
	FillSystemOrderTypeLiquidatePositionOnBackstop FillSystemOrderType = "LiquidatePositionOnBackstop"
	FillSystemOrderTypeOrderBookClosed             FillSystemOrderType = "OrderBookClosed"
)

type FundingHistoryOptions struct {
	SubAccountID *int    `json:"subaccountId"`
	Symbol       *string `json:"symbol"`
	LimitOffset
}

type FundingHistory struct {
	UserID               int       `json:"userId"`
	SubAccountID         *int      `json:"subaccountId"`
	Symbol               string    `json:"symbol"`
	Quantity             float64   `json:"quantity,string"`
	IntervalEndTimestamp time.Time `json:"intervalEndTimestamp" time_format:"2006-01-02T15:04:05.000Z"`
	FundingRate          float64   `json:"fundingRate,string"`
}

type OrderHistoryOptions struct {
	OrderID    *string     `json:"orderId"`
	Symbol     *string     `json:"symbol"`
	MarketType *MarketType `json:"marketType"`
	LimitOffset
}

type OrderHistory struct {
	ID                     string                  `json:"id"`
	ExecutedQuantity       *float64                `json:"executedQuantity,string"`
	ExecutedQuoteQuantity  *float64                `json:"executedQuoteQuantity,string"`
	ExpiryReason           *OrderExpiryReason      `json:"expiryReason"`
	OrderType              OrderType               `json:"orderType"`
	PostOnly               *bool                   `json:"postOnly"`
	Price                  *float64                `json:"price,string"`
	Quantity               *float64                `json:"quantity,string"`
	QuoteQuantity          *float64                `json:"quoteQuantity,string"`
	SelfTradePrevention    SelfTradePreventionType `json:"selfTradePrevention"`
	Status                 OrderStatus             `json:"status"`
	Side                   Side                    `json:"side"`
	StopLossTriggerPrice   *float64                `json:"stopLossTriggerPrice,string"`
	StopLossLimitPrice     *float64                `json:"stopLossLimitPrice,string"`
	StopLossTriggerBy      *float64                `json:"stopLossTriggerBy,string"`
	Symbol                 string                  `json:"symbol"`
	TakeProfitTriggerPrice *float64                `json:"takeProfitTriggerPrice,string"`
	TakeProfitLimitPrice   *float64                `json:"takeProfitLimitPrice,string"`
	TakeProfitTriggerBy    *float64                `json:"takeProfitTriggerBy,string"`
	TimeInForce            TimeInForce             `json:"timeInForce"`
	TriggerBy              *float64                `json:"triggerBy,string"`
	TriggerPrice           *float64                `json:"triggerPrice,string"`
	TriggerQuantity        *float64                `json:"triggerQuantity,string"`
}

type OrderExpiryReason string

const (
	OrderExpiryReasonAccountTradingSuspended        OrderExpiryReason = "AccountTradingSuspended"
	OrderExpiryReasonFillOrKill                     OrderExpiryReason = "FillOrKill"
	OrderExpiryReasonInsufficientBorrowableQuantity OrderExpiryReason = "InsufficientBorrowableQuantity"
	OrderExpiryReasonInsufficientFunds              OrderExpiryReason = "InsufficientFunds"
	OrderExpiryReasonInsufficientLiquidity          OrderExpiryReason = "InsufficientLiquidity"
	OrderExpiryReasonInvalidPrice                   OrderExpiryReason = "InvalidPrice"
	OrderExpiryReasonInvalidQuantity                OrderExpiryReason = "InvalidQuantity"
	OrderExpiryReasonImmediateOrCancel              OrderExpiryReason = "ImmediateOrCancel"
	OrderExpiryReasonInsufficientMargin             OrderExpiryReason = "InsufficientMargin"
	OrderExpiryReasonLiquidation                    OrderExpiryReason = "Liquidation"
	OrderExpiryReasonPostOnlyTaker                  OrderExpiryReason = "PostOnlyTaker"
	OrderExpiryReasonReduceOnlyNotReduced           OrderExpiryReason = "ReduceOnlyNotReduced"
	OrderExpiryReasonSelfTradePrevention            OrderExpiryReason = "SelfTradePrevention"
	OrderExpiryReasonPriceImpact                    OrderExpiryReason = "PriceImpact"
	OrderExpiryReasonUnknown                        OrderExpiryReason = "Unknown"
	OrderExpiryReasonUserPermissions                OrderExpiryReason = "UserPermissions"
)

type PnlHistoryOptions struct {
	SubAccountID *int    `json:"subaccountId"`
	Symbol       *string `json:"symbol"`
	LimitOffset
}

type PnlHistory struct {
	PnlRealized float64   `json:"pnlRealized,string"`
	Symbol      string    `json:"symbol"`
	Timestamp   time.Time `json:"timestamp" time_format:"2006-01-02T15:04:05.000Z"`
}

type SettlementHistoryOptions struct {
	Source *string `json:"source"`
	LimitOffset
}

type SettlementSourceType string

const (
	SettlementSourceTypeBackstopLiquidation             SettlementSourceType = "BackstopLiquidation"
	SettlementSourceTypeCulledBorrowInterest            SettlementSourceType = "CulledBorrowInterest"
	SettlementSourceTypeCulledRealizePnl                SettlementSourceType = "CulledRealizePnl"
	SettlementSourceTypeCulledRealizePnlBookUtilization SettlementSourceType = "CulledRealizePnlBookUtilization"
	SettlementSourceTypeFundingPayment                  SettlementSourceType = "FundingPayment"
	SettlementSourceTypeRealizePnl                      SettlementSourceType = "RealizePnl"
	SettlementSourceTypeTradingFees                     SettlementSourceType = "TradingFees"
	SettlementSourceTypeTradingFeesSystem               SettlementSourceType = "TradingFeesSystem"
)

type SettlementHistory struct {
	Quantity     float64   `json:"quantity,string"`
	Source       string    `json:"source"`
	SubaccountID *int      `json:"subaccountId"`
	Timestamp    time.Time `json:"timestamp" time_format:"2006-01-02T15:04:05.000Z"`
	UserID       int       `json:"userId"`
}
