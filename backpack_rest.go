package backpackgo

import (
	"fmt"
	"maps"
	"net/url"
	"strconv"
	"time"

	"github.com/feeeei/backpack-go/constants"
	"github.com/feeeei/backpack-go/models"
	"github.com/feeeei/backpack-go/rest"
	"github.com/feeeei/backpack-go/utils"
	json "github.com/json-iterator/go"
	"github.com/liamylian/jsontime/v3"
	"resty.dev/v3"
)

type BackpackREST struct {
	BaseURL   string
	APIKey    string
	APISecret string
	Windows   time.Duration
	client    *resty.Client
}

func init() {
	json.RegisterExtension(jsontime.NewCustomTimeExtension())
}

func NewRESTClient(options ...rest.Options) *BackpackREST {
	opts := rest.DefaultRESTOptions()
	client := resty.New().
		SetBaseURL(opts.BaseURL).
		SetResponseMiddlewares(handleError).
		SetAllowMethodDeletePayload(true)

	for _, option := range options {
		option(opts, client)
	}

	return &BackpackREST{
		BaseURL:   opts.BaseURL,
		APIKey:    opts.APIKey,
		APISecret: opts.APISecret,
		Windows:   opts.Windows,
		client:    client,
	}
}

func (b *BackpackREST) GetMarketAssets() ([]*models.MarketAssets, error) {
	path := "/api/v1/assets"
	return Response[[]*models.MarketAssets](Request(b, "GET", path, nil))
}

func (b *BackpackREST) GetMarketCollateral() ([]*models.MarketCollateral, error) {
	path := "/api/v1/collateral"
	return Response[[]*models.MarketCollateral](Request(b, "GET", path, nil))
}

func (b *BackpackREST) GetBorrowLendMarkets() ([]*models.BorrowLendMarket, error) {
	path := "/api/v1/borrowLend/markets"
	return Response[[]*models.BorrowLendMarket](Request(b, "GET", path, nil))
}

// symbol is optional, ex USDT、USDC、SOL...
func (b *BackpackREST) GetBorrowLendMarketsHistory(interval models.BorrowLendMarketHistoryInterval, symbol ...string) ([]*models.BorrowLendMarketHistory, error) {
	path := "/api/v1/borrowLend/markets/history"
	params := map[string]string{"interval": string(interval)}
	if len(symbol) > 0 {
		params["symbol"] = symbol[0]
	}
	return Response[[]*models.BorrowLendMarketHistory](Request(b, "GET", path, params))
}

func (b *BackpackREST) GetMarkets() ([]*models.Market, error) {
	path := "/api/v1/markets"
	return Response[[]*models.Market](Request(b, "GET", path, nil))
}

func (b *BackpackREST) GetMarket(symbol string) (*models.Market, error) {
	path := "/api/v1/market"
	params := map[string]string{"symbol": symbol}
	return Response[*models.Market](Request(b, "GET", path, params))
}

// interval is optional
func (b *BackpackREST) GetTickers(interval ...models.TickerInterval) ([]*models.Ticker, error) {
	path := "/api/v1/tickers"
	params := map[string]string{}
	if len(interval) > 0 {
		params["interval"] = string(interval[0])
	}
	return Response[[]*models.Ticker](Request(b, "GET", path, params))
}

// interval is optional
func (b *BackpackREST) GetTicker(symbol string, interval ...models.TickerInterval) (*models.Ticker, error) {
	path := "/api/v1/ticker"
	params := map[string]string{"symbol": symbol}
	if len(interval) > 0 {
		params["interval"] = string(interval[0])
	}
	return Response[*models.Ticker](Request(b, "GET", path, params))
}

func (b *BackpackREST) GetDepth(symbol string) (*models.Depth, error) {
	path := "/api/v1/depth"
	params := map[string]string{"symbol": symbol}
	return Response[*models.Depth](Request(b, "GET", path, params))
}

// endTime is optional
func (b *BackpackREST) GetKlines(symbol string, interval models.KlineInterval, startTime time.Time, endTime ...time.Time) ([]*models.Kline, error) {
	path := "/api/v1/klines"
	params := map[string]string{"symbol": symbol, "interval": string(interval), "startTime": fmt.Sprintf("%d", startTime.UTC().Unix())}
	if len(endTime) > 0 {
		params["endTime"] = fmt.Sprintf("%d", endTime[0].UTC().Unix())
	}
	return Response[[]*models.Kline](Request(b, "GET", path, params))
}

// symbol is optional
func (b *BackpackREST) GetMarkPrices(symbol ...string) ([]*models.MarkPrice, error) {
	path := "/api/v1/markPrices"
	params := map[string]string{}
	if len(symbol) > 0 {
		params["interval"] = string(symbol[0])
	}
	return Response[[]*models.MarkPrice](Request(b, "GET", path, params))
}

func (b *BackpackREST) GetOpenInterest(symbol ...string) ([]*models.OpenInterest, error) {
	path := "/api/v1/openInterest"
	params := map[string]string{}
	if len(symbol) > 0 {
		params["symbol"] = string(symbol[0])
	}
	return Response[[]*models.OpenInterest](Request(b, "GET", path, params))
}

func (b *BackpackREST) GetFundingRates(symbol string) (*models.PageHeaders, []*models.FundingRate, error) {
	path := "/api/v1/fundingRates"
	params := map[string]string{"symbol": symbol}
	resp, err := Request(b, "GET", path, params)
	rates, err := Response[[]*models.FundingRate](resp, err)
	if err != nil {
		return nil, nil, err
	}
	headers := models.ParseFundingRateHeaders(resp.Header())
	return headers, rates, err
}

func (b *BackpackREST) GetStatus() (*models.Status, error) {
	path := "/api/v1/status"
	return Response[*models.Status](Request(b, "GET", path, nil))
}

func (b *BackpackREST) Ping() error {
	path := "/api/v1/ping"
	_, err := Request(b, "GET", path, nil)
	return err
}

func (b *BackpackREST) GetTime() (*time.Time, error) {
	path := "/api/v1/time"
	resp, err := Request(b, "GET", path, nil)
	if err != nil {
		return nil, err
	}
	timestamp, err := strconv.ParseInt(string(resp.Bytes()), 10, 64)
	if err != nil {
		return nil, err
	}
	return utils.Ptr(time.UnixMilli(timestamp)), nil
}

// limitoffset is optional
func (b *BackpackREST) GetTrades(symbol string, limitoffset ...models.LimitOffset) ([]*models.Trade, error) {
	path := "/api/v1/trades"
	params := map[string]string{"symbol": symbol}
	if limitoffset != nil {
		maps.Copy(params, utils.StructToMap[map[string]string](limitoffset[0]))
	}
	return Response[[]*models.Trade](Request(b, "GET", path, params))
}

func (b *BackpackREST) GetTradesHistory(symbol string, limit ...int) ([]*models.Trade, error) {
	path := "/api/v1/trades/history"
	params := map[string]string{"symbol": symbol}
	if len(limit) > 0 {
		params["limit"] = strconv.Itoa(limit[0])
	}
	return Response[[]*models.Trade](Request(b, "GET", path, params))
}

func (b *BackpackREST) GetAccount() (*models.Account, error) {
	path := "/api/v1/account"
	return Response[*models.Account](RequestWithAuth(b, "GET", path, "accountQuery", nil))
}

func (b *BackpackREST) UpdateAccount(account *models.AccountUpdateble) error {
	path := "/api/v1/account"
	_, err := RequestWithAuth(b, "PATCH", path, "accountUpdate", account)
	return err
}

func (b *BackpackREST) GetAccountMaxBorrow(asset string) (*models.AccountBorrowLimit, error) {
	path := "/api/v1/account/limits/borrow"
	params := map[string]string{"symbol": asset}
	return Response[*models.AccountBorrowLimit](RequestWithAuth(b, "GET", path, "maxBorrowQuantity", params))
}

func (b *BackpackREST) GetAccountMaxOrder(symbol string, side models.Side, args ...models.AccountOrderLimitOptions) (*models.AccountOrderLimit, error) {
	path := "/api/v1/account/limits/order"
	params := map[string]string{"symbol": symbol, "side": string(side)}
	if len(args) > 0 {
		maps.Copy(params, utils.StructToMap[map[string]string](args[0]))
	}
	return Response[*models.AccountOrderLimit](RequestWithAuth(b, "GET", path, "maxOrderQuantity", params))
}

func (b *BackpackREST) GetAccountMaxWithdrawal(asset string, args ...models.AccountWithdrawalLimitOptions) (*models.AccountWithdrawalLimit, error) {
	path := "/api/v1/account/limits/withdrawal"
	params := map[string]string{"symbol": asset}
	if len(args) > 0 {
		maps.Copy(params, utils.StructToMap[map[string]string](args[0]))
	}
	return Response[*models.AccountWithdrawalLimit](RequestWithAuth(b, "GET", path, "maxWithdrawalQuantity", params))
}

func (b *BackpackREST) GetBorrowLendPositions() ([]*models.BorrowLend, error) {
	path := "/api/v1/borrowLend/positions"
	return Response[[]*models.BorrowLend](RequestWithAuth(b, "GET", path, "borrowLendPositionQuery", nil))
}

func (b *BackpackREST) ExecuteBorrowLend(asset string, side models.BorrowLendSide, quantity float64) error {
	path := "/api/v1/borrowLend"
	params := map[string]string{"symbol": asset, "side": string(side), "quantity": fmt.Sprintf("%f", quantity)}
	_, err := RequestWithAuth(b, "POST", path, "borrowLendExecute", params)
	return err
}

func (b *BackpackREST) GetBalance() (map[string]*models.AssetBalance, error) {
	path := "/api/v1/capital"
	return Response[map[string]*models.AssetBalance](RequestWithAuth(b, "GET", path, "balanceQuery", nil))
}

func (b *BackpackREST) GetAccountCollateral() (*models.Collateral, error) {
	path := "/api/v1/capital/collateral"
	return Response[*models.Collateral](RequestWithAuth(b, "GET", path, "collateralQuery", nil))
}

func (b *BackpackREST) GetDeposits(filter ...models.DateFilter) ([]*models.Deposit, error) {
	path := "/wapi/v1/capital/deposits"
	return Response[[]*models.Deposit](RequestWithAuth(b, "GET", path, "depositQueryAll", nil))
}

func (b *BackpackREST) GetDepositAddress(blockchain string) (*models.DepositAddress, error) {
	path := "/wapi/v1/capital/deposit/address"
	params := map[string]string{"blockchain": blockchain}
	return Response[*models.DepositAddress](RequestWithAuth(b, "GET", path, "depositAddressQuery", params))
}

func (b *BackpackREST) GetWithdrawals(filter ...models.DateFilter) ([]*models.Withdrawal, error) {
	path := "/wapi/v1/capital/withdrawals"
	return Response[[]*models.Withdrawal](RequestWithAuth(b, "GET", path, "withdrawalQueryAll", nil))
}

func (b *BackpackREST) RequestWithdrawal(asset string, quantity float64, address, blockchain string, options ...models.WithdrawalOptions) (*models.Withdrawal, error) {
	path := "/wapi/v1/capital/withdrawals"
	params := map[string]any{"symbol": asset, "quantity": fmt.Sprintf("%f", quantity), "address": address, "blockchain": blockchain}
	if len(options) > 0 {
		maps.Copy(params, utils.StructToMap[map[string]any](options[0]))
	}
	return Response[*models.Withdrawal](RequestWithAuth(b, "POST", path, "withdraw", params))
}

func (b *BackpackREST) GetPositions() ([]*models.Position, error) {
	path := "/api/v1/position"
	return Response[[]*models.Position](RequestWithAuth(b, "GET", path, "positionQuery", nil))
}

func (b *BackpackREST) GetBorrowLendHistory(options ...models.BorrowHistoryOptions) ([]*models.BorrowLendHistory, error) {
	path := "/wapi/v1/history/borrowLend"
	params := map[string]string{}
	if len(options) > 0 {
		params = utils.StructToMap[map[string]string](options[0])
	}
	return Response[[]*models.BorrowLendHistory](RequestWithAuth(b, "GET", path, "borrowHistoryQueryAll", params))
}

func (b *BackpackREST) GetInterestHistory(options ...models.InterestHistoryOptions) ([]*models.InterestHistory, error) {
	path := "/wapi/v1/history/interest"
	params := map[string]string{}
	if len(options) > 0 {
		params = utils.StructToMap[map[string]string](options[0])
	}
	return Response[[]*models.InterestHistory](RequestWithAuth(b, "GET", path, "interestHistoryQueryAll", params))
}

// TODO: Can't yet query that with a signed request, unknown instruction
// func (b *BackpackREST) GetBorrowPositionsHistory(options ...models.BorrowPostionHistoryOptions) ([]*models.BorrowPositionHistory, error) {
// 	path := "/wapi/v1/history/borrowLend/positions"
// 	params := map[string]string{}
// 	if len(options) > 0 {
// 		params = utils.StructToMap[map[string]string](options[0])
// 	}
// 	return Response[[]*models.BorrowPositionHistory](RequestWithAuth(b, "GET", path, "borrowLendPositionsQueryAll", params))
// }

func (b *BackpackREST) GetFillHistory(options ...models.FillHistoryOptions) ([]*models.FillHistory, error) {
	path := "/wapi/v1/history/fills"
	params := map[string]string{}
	if len(options) > 0 {
		params = utils.StructToMap[map[string]string](options[0])
	}
	return Response[[]*models.FillHistory](RequestWithAuth(b, "GET", path, "fillHistoryQueryAll", params))
}

func (b *BackpackREST) GetFundingHistory(options ...models.FundingHistoryOptions) ([]*models.FundingHistory, error) {
	path := "/wapi/v1/history/funding"
	params := map[string]string{}
	if len(options) > 0 {
		params = utils.StructToMap[map[string]string](options[0])
	}
	return Response[[]*models.FundingHistory](RequestWithAuth(b, "GET", path, "fundingHistoryQueryAll", params))
}

func (b *BackpackREST) GetOrdersHistory(options ...models.OrderHistoryOptions) ([]*models.OrderHistory, error) {
	path := "/wapi/v1/history/orders"
	params := map[string]string{}
	if len(options) > 0 {
		params = utils.StructToMap[map[string]string](options[0])
	}
	return Response[[]*models.OrderHistory](RequestWithAuth(b, "GET", path, "orderHistoryQueryAll", params))
}

func (b *BackpackREST) GetPnlHistory(options ...models.PnlHistoryOptions) ([]*models.PnlHistory, error) {
	path := "/wapi/v1/history/pnl"
	params := map[string]string{}
	if len(options) > 0 {
		params = utils.StructToMap[map[string]string](options[0])
	}
	return Response[[]*models.PnlHistory](RequestWithAuth(b, "GET", path, "pnlHistoryQueryAll", params))
}

func (b *BackpackREST) GetSettlementHistory(options ...models.SettlementHistoryOptions) ([]*models.SettlementHistory, error) {
	path := "/wapi/v1/history/settlement"
	params := map[string]string{}
	if len(options) > 0 {
		params = utils.StructToMap[map[string]string](options[0])
	}
	return Response[[]*models.SettlementHistory](RequestWithAuth(b, "GET", path, "settlementHistoryQueryAll", params))
}

func (b *BackpackREST) GetOrderByClientID(symbol string, clientID int) (*models.Order, error) {
	path := "/api/v1/order"
	params := map[string]string{"symbol": symbol, "clientId": strconv.Itoa(clientID)}
	return Response[*models.Order](RequestWithAuth(b, "GET", path, "orderQuery", params))
}

func (b *BackpackREST) GetOrderByOrderID(symbol, orderID string) (*models.Order, error) {
	url, _ := url.JoinPath(b.BaseURL + "/api/v1/order")
	params := map[string]string{"symbol": symbol, "orderId": orderID}
	return Response[*models.Order](RequestWithAuth(b, "GET", url, "orderQuery", params))
}

func (b *BackpackREST) ExecuteMarketOrder(symbol string, side models.Side, quantity float64, options ...models.OrderOptions) (*models.Order, error) {
	path := "/api/v1/order"
	params := map[string]string{
		"orderType": string(models.OrderTypeMarket),
		"symbol":    symbol,
		"side":      string(side),
		"quantity":  fmt.Sprintf("%f", quantity),
	}
	if len(options) > 0 {
		maps.Copy(params, utils.StructToMap[map[string]string](options[0]))
	}
	return Response[*models.Order](RequestWithAuth(b, "POST", path, "orderExecute", params))
}

func (b *BackpackREST) ExecuteLimitOrder(symbol string, side models.Side, quantity float64, price float64, options ...models.OrderOptions) (*models.Order, error) {
	path := "/api/v1/order"
	params := map[string]string{
		"orderType": string(models.OrderTypeLimit),
		"symbol":    symbol,
		"side":      string(side),
		"quantity":  fmt.Sprintf("%f", quantity),
		"price":     fmt.Sprintf("%f", price),
	}
	if len(options) > 0 {
		maps.Copy(params, utils.StructToMap[map[string]string](options[0]))
	}
	return Response[*models.Order](RequestWithAuth(b, "POST", path, "orderExecute", params))
}

func (b *BackpackREST) ExecuteStopMarketOrder(symbol string, side models.Side, quantity float64, price float64, options ...models.OrderOptions) (*models.Order, error) {
	path := "/api/v1/order"
	params := map[string]string{
		"orderType": string(models.OrderTypeStopMarket),
		"symbol":    symbol,
		"side":      string(side),
		"quantity":  fmt.Sprintf("%f", quantity),
		"price":     fmt.Sprintf("%f", price),
	}
	if len(options) > 0 {
		maps.Copy(params, utils.StructToMap[map[string]string](options[0]))
	}
	return Response[*models.Order](RequestWithAuth(b, "POST", path, "orderExecute", params))
}

func (b *BackpackREST) ExecuteStopLimitOrder(symbol string, side models.Side, quantity float64, price, triggerPrice float64, options ...models.OrderOptions) (*models.Order, error) {
	path := "/api/v1/order"
	params := map[string]string{
		"orderType":    string(models.OrderTypeStopLimit),
		"symbol":       symbol,
		"side":         string(side),
		"quantity":     fmt.Sprintf("%f", quantity),
		"price":        fmt.Sprintf("%f", price),
		"triggerPrice": fmt.Sprintf("%f", triggerPrice),
	}
	if len(options) > 0 {
		maps.Copy(params, utils.StructToMap[map[string]string](options[0]))
	}
	return Response[*models.Order](RequestWithAuth(b, "POST", path, "orderExecute", params))
}

func (b *BackpackREST) CancelOrderByOrderID(symbol string, orderID string) (*models.Order, error) {
	path := "/api/v1/order"
	params := map[string]string{"symbol": symbol, "orderId": orderID}
	return Response[*models.Order](RequestWithAuth(b, "DELETE", path, "orderCancel", params))
}

func (b *BackpackREST) CancelOrderByClientID(symbol string, clientID int) (*models.Order, error) {
	path := "/api/v1/order"
	params := map[string]string{"symbol": symbol, "clientId": strconv.Itoa(clientID)}
	return Response[*models.Order](RequestWithAuth(b, "DELETE", path, "orderCancel", params))
}

// symbol is optional
// marketType is optional
func (b *BackpackREST) GetOrders(symbol *string, marketType *models.MarketType) ([]*models.Order, error) {
	path := "/api/v1/orders"
	params := map[string]string{}
	if symbol != nil {
		params["symbol"] = *symbol
	}
	if marketType != nil {
		params["marketType"] = string(*marketType)
	}
	return Response[[]*models.Order](RequestWithAuth(b, "GET", path, "orderQueryAll", params))
}

func (b *BackpackREST) CancelOrders(symbol string, marketType ...models.OrderType) ([]*models.Order, error) {
	path := "/api/v1/orders"
	params := map[string]string{"symbol": symbol}
	if len(marketType) > 0 {
		params["marketType"] = string(marketType[0])
	}
	return Response[[]*models.Order](RequestWithAuth(b, "DELETE", path, "orderCancelAll", params))
}

func (b *BackpackREST) RequestForQuote(rfqId string, bidPrice, askPrice float64, clientID *int) (*models.Quote, error) {
	url, _ := url.JoinPath(b.BaseURL + "/api/v1/rfq/quote")
	params := map[string]string{"rfqId": rfqId, "bidPrice": fmt.Sprintf("%f", bidPrice), "askPrice": fmt.Sprintf("%f", askPrice)}
	if clientID != nil {
		params["clientId"] = strconv.Itoa(*clientID)
	}
	return Response[*models.Quote](RequestWithAuth(b, "POST", url, "quoteSubmit", params))
}

func RequestWithAuth(api *BackpackREST, method, url string, instruction string, params any) (resp *resty.Response, err error) {
	if api.APIKey == "" || api.APISecret == "" {
		return nil, ErrInvalidAPIKeyOrSecret
	}
	payload := map[string]any{}
	if params != nil {
		payload = utils.StructToMap[map[string]any](params)
	}
	timestamp := time.Now().UnixNano() / int64(time.Millisecond)
	headers, err := auth(api.APIKey, api.APISecret, payload, instruction, timestamp, api.Windows.Milliseconds())
	if err != nil {
		return nil, err
	}
	return Request(api, method, url, params, headers)
}

func Request(api *BackpackREST, method, path string, params any, headers ...map[string]string) (*resty.Response, error) {
	request := api.client.R().
		SetHeader("User-Agent", constants.UserAgent).
		SetHeader("Content-Type", "application/json")
	if len(headers) > 0 {
		request.SetHeaders(headers[0])
	}
	var fn func(url string) (*resty.Response, error)
	switch method {
	case "GET":
		if params != nil {
			request.SetQueryParams(params.(map[string]string))
		}
		fn = request.Get
	case "POST":
		request.SetBody(params)
		fn = request.Post
	case "PATCH":
		request.SetBody(params)
		fn = request.Patch
	case "DELETE":
		request.SetBody(params)
		fn = request.Delete
	}
	url, _ := url.JoinPath(api.BaseURL, path)
	return fn(url)
}

func Response[T any](response *resty.Response, e error) (result T, err error) {
	if e != nil {
		return result, e
	}

	err = json.Unmarshal(response.Bytes(), &result)
	return
}

func handleError(_ *resty.Client, res *resty.Response) error {
	if res.IsError() {
		backpackError := &BackpackError{}
		if err := json.Unmarshal(res.Bytes(), backpackError); err != nil {
			backpackError.Code = "UNKNOWN"
			backpackError.Message = res.String()
			return backpackError
		}
		return backpackError
	}
	return nil
}
