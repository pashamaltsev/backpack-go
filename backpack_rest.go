package backpackgo

import (
	"fmt"
	"maps"
	"net/url"
	"strconv"
	"time"

	"github.com/feeeei/backpack-go/constants"
	"github.com/feeeei/backpack-go/models"
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

func NewRESTClient(options ...Options) *BackpackREST {
	opts := defaultRESTOptions()
	for _, option := range options {
		option(opts)
	}
	restyClient := resty.New().
		SetBaseURL(opts.BaseURL).
		SetResponseMiddlewares(handleError)

	return &BackpackREST{
		BaseURL:   opts.BaseURL,
		APIKey:    opts.APIKey,
		APISecret: opts.APISecret,
		Windows:   opts.Windows,
		client:    restyClient,
	}
}

func (b *BackpackREST) GetMarketAssets() ([]*models.MarketAssets, error) {
	url, _ := url.JoinPath(b.BaseURL + "/api/v1/assets")
	return response[[]*models.MarketAssets](request(b, "GET", url, nil))
}

func (b *BackpackREST) GetMarketCollateral() ([]*models.MarketCollateral, error) {
	url, _ := url.JoinPath(b.BaseURL + "/api/v1/collateral")
	return response[[]*models.MarketCollateral](request(b, "GET", url, nil))
}

func (b *BackpackREST) GetBorrowLendMarkets() ([]*models.BorrowLendMarket, error) {
	url, _ := url.JoinPath(b.BaseURL + "/api/v1/borrowLend/markets")
	return response[[]*models.BorrowLendMarket](request(b, "GET", url, nil))
}

// symbol is optional, ex USDT、USDC、SOL...
func (b *BackpackREST) GetBorrowLendMarketsHistory(interval models.BorrowLendMarketHistoryInterval, symbol ...string) ([]*models.BorrowLendMarketHistory, error) {
	url, _ := url.JoinPath(b.BaseURL + "/api/v1/borrowLend/markets/history")
	params := map[string]string{"interval": string(interval)}
	if len(symbol) > 0 {
		params["symbol"] = symbol[0]
	}
	return response[[]*models.BorrowLendMarketHistory](request(b, "GET", url, params))
}

func (b *BackpackREST) GetMarkets() ([]*models.Market, error) {
	url, _ := url.JoinPath(b.BaseURL + "/api/v1/markets")
	return response[[]*models.Market](request(b, "GET", url, nil))
}

func (b *BackpackREST) GetMarket(symbol string) (*models.Market, error) {
	url, _ := url.JoinPath(b.BaseURL + "/api/v1/market")
	params := map[string]string{"symbol": symbol}
	return response[*models.Market](request(b, "GET", url, params))
}

// interval is optional
func (b *BackpackREST) GetTickers(interval ...models.TickerInterval) ([]*models.Ticker, error) {
	url, _ := url.JoinPath(b.BaseURL + "/api/v1/tickers")
	params := map[string]string{}
	if len(interval) > 0 {
		params["interval"] = string(interval[0])
	}
	return response[[]*models.Ticker](request(b, "GET", url, params))
}

// interval is optional
func (b *BackpackREST) GetTicker(symbol string, interval ...models.TickerInterval) (*models.Ticker, error) {
	url, _ := url.JoinPath(b.BaseURL + "/api/v1/ticker")
	params := map[string]string{"symbol": symbol}
	if len(interval) > 0 {
		params["interval"] = string(interval[0])
	}
	return response[*models.Ticker](request(b, "GET", url, params))
}

func (b *BackpackREST) GetDepth(symbol string) (*models.Depth, error) {
	url, _ := url.JoinPath(b.BaseURL + "/api/v1/depth")
	params := map[string]string{"symbol": symbol}
	return response[*models.Depth](request(b, "GET", url, params))
}

// endTime is optional
func (b *BackpackREST) GetKlines(symbol string, interval models.KlineInterval, startTime time.Time, endTime ...time.Time) ([]*models.Kline, error) {
	url, _ := url.JoinPath(b.BaseURL + "/api/v1/klines")
	params := map[string]string{"symbol": symbol, "interval": string(interval), "startTime": fmt.Sprintf("%d", startTime.UTC().Unix())}
	if len(endTime) > 0 {
		params["endTime"] = fmt.Sprintf("%d", endTime[0].UTC().Unix())
	}
	return response[[]*models.Kline](request(b, "GET", url, params))
}

// symbol is optional
func (b *BackpackREST) GetMarkPrices(symbol ...string) ([]*models.MarkPrice, error) {
	url, _ := url.JoinPath(b.BaseURL + "/api/v1/markPrices")
	params := map[string]string{}
	if len(symbol) > 0 {
		params["interval"] = string(symbol[0])
	}
	return response[[]*models.MarkPrice](request(b, "GET", url, params))
}

func (b *BackpackREST) GetOpenInterest(symbol ...string) ([]*models.OpenInterest, error) {
	url, _ := url.JoinPath(b.BaseURL + "/api/v1/openInterest")
	params := map[string]string{}
	if len(symbol) > 0 {
		params["symbol"] = string(symbol[0])
	}
	return response[[]*models.OpenInterest](request(b, "GET", url, params))
}

func (b *BackpackREST) GetFundingRates(symbol string) (*models.PageHeaders, []*models.FundingRate, error) {
	url, _ := url.JoinPath(b.BaseURL + "/api/v1/fundingRates")
	params := map[string]string{"symbol": symbol}
	resp, err := request(b, "GET", url, params)
	rates, err := response[[]*models.FundingRate](resp, err)
	if err != nil {
		return nil, nil, err
	}
	headers := models.ParseFundingRateHeaders(resp.Header())
	return headers, rates, err
}

func (b *BackpackREST) GetStatus() (*models.Status, error) {
	url, _ := url.JoinPath(b.BaseURL + "/api/v1/status")
	return response[*models.Status](request(b, "GET", url, nil))
}

func (b *BackpackREST) Ping() error {
	url, _ := url.JoinPath(b.BaseURL + "/api/v1/ping")
	_, err := request(b, "GET", url, nil)
	return err
}

func (b *BackpackREST) GetTime() (*time.Time, error) {
	url, _ := url.JoinPath(b.BaseURL + "/api/v1/time")
	resp, err := request(b, "GET", url, nil)
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
	url, _ := url.JoinPath(b.BaseURL + "/api/v1/trades")
	params := map[string]string{"symbol": symbol}
	if limitoffset != nil {
		maps.Copy(params, utils.StructToMap[map[string]string](limitoffset[0]))
	}
	return response[[]*models.Trade](request(b, "GET", url, params))
}

func (b *BackpackREST) GetTradesHistory(symbol string, limit ...int) ([]*models.Trade, error) {
	url, _ := url.JoinPath(b.BaseURL + "/api/v1/trades/history")
	params := map[string]string{"symbol": symbol}
	if len(limit) > 0 {
		params["limit"] = strconv.Itoa(limit[0])
	}
	return response[[]*models.Trade](request(b, "GET", url, params))
}

func (b *BackpackREST) GetAccount() (*models.Account, error) {
	url, _ := url.JoinPath(b.BaseURL + "/api/v1/account")
	return response[*models.Account](requestWithAuth(b, "GET", url, "accountQuery", nil))
}

func (b *BackpackREST) UpdateAccount(account *models.AccountUpdateble) error {
	url, _ := url.JoinPath(b.BaseURL + "/api/v1/account")
	_, err := requestWithAuth(b, "PATCH", url, "accountUpdate", account)
	return err
}

func (b *BackpackREST) GetAccountMaxBorrow(asset string) (*models.AccountBorrowLimit, error) {
	url, _ := url.JoinPath(b.BaseURL + "/api/v1/account/limits/borrow")
	params := map[string]string{"symbol": asset}
	return response[*models.AccountBorrowLimit](requestWithAuth(b, "GET", url, "maxBorrowQuantity", params))
}

func (b *BackpackREST) GetAccountMaxOrder(symbol string, side models.Side, args ...models.AccountOrderLimitOptions) (*models.AccountOrderLimit, error) {
	url, _ := url.JoinPath(b.BaseURL + "/api/v1/account/limits/order")
	params := map[string]string{"symbol": symbol, "side": string(side)}
	if len(args) > 0 {
		maps.Copy(params, utils.StructToMap[map[string]string](args[0]))
	}
	return response[*models.AccountOrderLimit](requestWithAuth(b, "GET", url, "maxOrderQuantity", params))
}

func (b *BackpackREST) GetAccountMaxWithdrawal(asset string, args ...models.AccountWithdrawalLimitOptions) (*models.AccountWithdrawalLimit, error) {
	url, _ := url.JoinPath(b.BaseURL + "/api/v1/account/limits/withdrawal")
	params := map[string]string{"symbol": asset}
	if len(args) > 0 {
		maps.Copy(params, utils.StructToMap[map[string]string](args[0]))
	}
	fmt.Println(utils.StructToMap[map[string]string](args[0]))
	fmt.Println("params", params["autoBorrow"])
	return response[*models.AccountWithdrawalLimit](requestWithAuth(b, "GET", url, "maxWithdrawalQuantity", params))
}

func (b *BackpackREST) GetBorrowLendPositions() ([]*models.BorrowLend, error) {
	url, _ := url.JoinPath(b.BaseURL + "/api/v1/borrowLend/positions")
	return response[[]*models.BorrowLend](requestWithAuth(b, "GET", url, "borrowLendPositionQuery", nil))
}

func (b *BackpackREST) ExecuteBorrowLend(asset string, side models.BorrowLendSide, quantity float64) error {
	url, _ := url.JoinPath(b.BaseURL + "/api/v1/borrowLend/market")
	params := map[string]string{"symbol": asset, "side": string(side), "quantity": fmt.Sprintf("%f", quantity)}
	_, err := requestWithAuth(b, "POST", url, "borrowLendMarketQuery", params)
	return err
}

func (b *BackpackREST) GetBalance() (map[string]*models.AssetBalance, error) {
	url, _ := url.JoinPath(b.BaseURL + "/api/v1/capital")
	return response[map[string]*models.AssetBalance](requestWithAuth(b, "GET", url, "balanceQuery", nil))
}

func (b *BackpackREST) GetCollateral() (*models.Collateral, error) {
	url, _ := url.JoinPath(b.BaseURL + "/api/v1/capital/collateral")
	return response[*models.Collateral](requestWithAuth(b, "GET", url, "collateralQuery", nil))
}

func (b *BackpackREST) GetDeposits(filter ...models.DateFilter) (*models.Deposit, error) {
	url, _ := url.JoinPath(b.BaseURL + "/wapi/v1/capital/deposits")
	return response[*models.Deposit](requestWithAuth(b, "GET", url, "depositQueryAll", nil))
}

func (b *BackpackREST) GetDepositAddress(blockchain string) (*models.DepositAddress, error) {
	url, _ := url.JoinPath(b.BaseURL + "/wapi/v1/capital/deposit/address")
	params := map[string]string{"blockchain": blockchain}
	return response[*models.DepositAddress](requestWithAuth(b, "GET", url, "depositAddressQuery", params))
}

func (b *BackpackREST) GetWithdrawals(filter ...models.DateFilter) ([]*models.Withdrawal, error) {
	url, _ := url.JoinPath(b.BaseURL + "/wapi/v1/capital/withdrawals")
	return response[[]*models.Withdrawal](requestWithAuth(b, "GET", url, "withdrawalQueryAll", nil))
}

func (b *BackpackREST) RequestWithdrawal(asset string, quantity float64, address, blockchain string, options ...models.WithdrawalOptions) (*models.Withdrawal, error) {
	url, _ := url.JoinPath(b.BaseURL + "/wapi/v1/capital/withdrawals")
	params := map[string]any{"symbol": asset, "quantity": fmt.Sprintf("%f", quantity), "address": address}
	if len(options) > 0 {
		maps.Copy(params, utils.StructToMap[map[string]any](options[0]))
	}
	return response[*models.Withdrawal](requestWithAuth(b, "POST", url, "withdraw", params))
}

func (b *BackpackREST) GetPositions() ([]*models.Position, error) {
	url, _ := url.JoinPath(b.BaseURL + "/api/v1/positions")
	return response[[]*models.Position](requestWithAuth(b, "GET", url, "positionQuery", nil))
}

func (b *BackpackREST) GetBorrowLendHistory(options ...models.BorrowHistoryOptions) ([]*models.BorrowLendHistory, error) {
	url, _ := url.JoinPath(b.BaseURL + "/wapi/v1/history/borrowLend")
	params := map[string]string{}
	if len(options) > 0 {
		params = utils.StructToMap[map[string]string](options[0])
	}
	return response[[]*models.BorrowLendHistory](requestWithAuth(b, "GET", url, "borrowHistoryQueryAll", params))
}

func (b *BackpackREST) GetInterestHistory(options ...models.InterestHistoryOptions) ([]*models.InterestHistory, error) {
	url, _ := url.JoinPath(b.BaseURL + "/wapi/v1/history/interest")
	params := map[string]string{}
	if len(options) > 0 {
		params = utils.StructToMap[map[string]string](options[0])
	}
	return response[[]*models.InterestHistory](requestWithAuth(b, "GET", url, "interestHistoryQueryAll", params))
}

func (b *BackpackREST) GetBorrowPositionsHistory(options ...models.BorrowPostionHistoryOptions) ([]*models.BorrowPositionHistory, error) {
	url, _ := url.JoinPath(b.BaseURL + "/wapi/v1/history/borrowLend/positions")
	params := map[string]string{}
	if len(options) > 0 {
		params = utils.StructToMap[map[string]string](options[0])
	}
	return response[[]*models.BorrowPositionHistory](requestWithAuth(b, "GET", url, "borrowLendPositionHistoryQueryAll", params))
}

func (b *BackpackREST) GetFillHistory(options ...models.FillHistoryOptions) ([]*models.FillHistory, error) {
	url, _ := url.JoinPath(b.BaseURL + "/wapi/v1/history/fills")
	params := map[string]string{}
	if len(options) > 0 {
		params = utils.StructToMap[map[string]string](options[0])
	}
	return response[[]*models.FillHistory](requestWithAuth(b, "GET", url, "fillHistoryQueryAll", params))
}

func (b *BackpackREST) GetFundingHistory(options ...models.FundingHistoryOptions) ([]*models.FundingHistory, error) {
	url, _ := url.JoinPath(b.BaseURL + "/wapi/v1/history/funding")
	params := map[string]string{}
	if len(options) > 0 {
		params = utils.StructToMap[map[string]string](options[0])
	}
	return response[[]*models.FundingHistory](requestWithAuth(b, "GET", url, "fundingHistoryQueryAll", params))
}

func (b *BackpackREST) GetOrdersHistory(options ...models.OrderHistoryOptions) ([]*models.OrderHistory, error) {
	url, _ := url.JoinPath(b.BaseURL + "/wapi/v1/history/orders")
	params := map[string]string{}
	if len(options) > 0 {
		params = utils.StructToMap[map[string]string](options[0])
	}
	return response[[]*models.OrderHistory](requestWithAuth(b, "GET", url, "orderHistoryQueryAll", params))
}

func (b *BackpackREST) GetPnlHistory(options ...models.PnlHistoryOptions) ([]*models.PnlHistory, error) {
	url, _ := url.JoinPath(b.BaseURL + "/wapi/v1/history/pnl")
	params := map[string]string{}
	if len(options) > 0 {
		params = utils.StructToMap[map[string]string](options[0])
	}
	return response[[]*models.PnlHistory](requestWithAuth(b, "GET", url, "pnlHistoryQueryAll", params))
}

func (b *BackpackREST) GetSettlementHistory(options ...models.SettlementHistoryOptions) ([]*models.SettlementHistory, error) {
	url, _ := url.JoinPath(b.BaseURL + "/wapi/v1/history/settlement")
	params := map[string]string{}
	if len(options) > 0 {
		params = utils.StructToMap[map[string]string](options[0])
	}
	return response[[]*models.SettlementHistory](requestWithAuth(b, "GET", url, "settlementHistoryQueryAll", params))
}

func (b *BackpackREST) GetOrderByClientID(symbol string, clientID int) (*models.Order, error) {
	url, _ := url.JoinPath(b.BaseURL + "/api/v1/order")
	params := map[string]string{"symbol": symbol, "clientId": strconv.Itoa(clientID)}
	return response[*models.Order](requestWithAuth(b, "GET", url, "orderQuery", params))
}

func (b *BackpackREST) GetOrderByOrderID(symbol, orderID string) (*models.Order, error) {
	url, _ := url.JoinPath(b.BaseURL + "/api/v1/order")
	params := map[string]string{"symbol": symbol, "orderId": orderID}
	return response[*models.Order](requestWithAuth(b, "GET", url, "orderQuery", params))
}

func (b *BackpackREST) ExecuteMarketOrder(symbol string, side models.Side, quantity float64, options ...models.OrderOptions) (*models.Order, error) {
	url, _ := url.JoinPath(b.BaseURL + "/api/v1/order")
	params := map[string]string{"symbol": symbol, "side": string(side), "quantity": fmt.Sprintf("%f", quantity)}
	if len(options) > 0 {
		maps.Copy(params, utils.StructToMap[map[string]string](options[0]))
	}
	return response[*models.Order](requestWithAuth(b, "POST", url, "orderExecute", params))
}

func (b *BackpackREST) ExecuteLimitOrder(symbol string, side models.Side, quantity float64, price float64, options ...models.OrderOptions) (*models.Order, error) {
	url, _ := url.JoinPath(b.BaseURL + "/api/v1/order")
	params := map[string]string{"symbol": symbol, "side": string(side), "quantity": fmt.Sprintf("%f", quantity), "price": fmt.Sprintf("%f", price)}
	if len(options) > 0 {
		maps.Copy(params, utils.StructToMap[map[string]string](options[0]))
	}
	return response[*models.Order](requestWithAuth(b, "POST", url, "orderExecute", params))
}

func (b *BackpackREST) ExecuteStopMarketOrder(symbol string, side models.Side, quantity float64, price float64, options ...models.OrderOptions) (*models.Order, error) {
	url, _ := url.JoinPath(b.BaseURL + "/api/v1/order")
	params := map[string]string{"symbol": symbol, "side": string(side), "quantity": fmt.Sprintf("%f", quantity), "price": fmt.Sprintf("%f", price)}
	if len(options) > 0 {
		maps.Copy(params, utils.StructToMap[map[string]string](options[0]))
	}
	return response[*models.Order](requestWithAuth(b, "POST", url, "orderExecute", params))
}

func (b *BackpackREST) ExecuteStopLimitOrder(symbol string, side models.Side, quantity float64, price, triggerPrice float64, options ...models.OrderOptions) (*models.Order, error) {
	url, _ := url.JoinPath(b.BaseURL + "/api/v1/order")
	params := map[string]string{"symbol": symbol, "side": string(side), "quantity": fmt.Sprintf("%f", quantity), "price": fmt.Sprintf("%f", price), "triggerPrice": fmt.Sprintf("%f", triggerPrice)}
	if len(options) > 0 {
		maps.Copy(params, utils.StructToMap[map[string]string](options[0]))
	}
	return response[*models.Order](requestWithAuth(b, "POST", url, "orderExecute", params))
}

func (b *BackpackREST) CancelOrderByOrderID(symbol string, orderID string) (*models.Order, error) {
	url, _ := url.JoinPath(b.BaseURL + "/api/v1/order")
	params := map[string]string{"symbol": symbol, "orderId": orderID}
	return response[*models.Order](requestWithAuth(b, "DELETE", url, "orderCancel", params))
}

func (b *BackpackREST) CancelOrderByClientID(symbol string, clientID int) (*models.Order, error) {
	url, _ := url.JoinPath(b.BaseURL + "/api/v1/order")
	params := map[string]string{"symbol": symbol, "clientId": strconv.Itoa(clientID)}
	return response[*models.Order](requestWithAuth(b, "DELETE", url, "orderCancel", params))
}

// symbol is optional
// marketType is optional
func (b *BackpackREST) GetOrders(symbol *string, marketType *models.MarketType) ([]*models.Order, error) {
	url, _ := url.JoinPath(b.BaseURL + "/api/v1/orders")
	params := map[string]string{}
	if symbol != nil {
		params["symbol"] = *symbol
	}
	if marketType != nil {
		params["marketType"] = string(*marketType)
	}
	return response[[]*models.Order](requestWithAuth(b, "GET", url, "orderQueryAll", params))
}

func (b *BackpackREST) CancelOrders(symbol string, marketType *models.OrderType) ([]*models.Order, error) {
	url, _ := url.JoinPath(b.BaseURL + "/api/v1/orders")
	params := map[string]string{"symbol": symbol}
	if marketType != nil {
		params["marketType"] = string(*marketType)
	}
	return response[[]*models.Order](requestWithAuth(b, "DELETE", url, "orderCancelAll", params))
}

func (b *BackpackREST) RequestForQuote(rfqId string, bidPrice, askPrice float64, clientID *int) (*models.Quote, error) {
	url, _ := url.JoinPath(b.BaseURL + "/api/v1/rfq/quote")
	params := map[string]string{"rfqId": rfqId, "bidPrice": fmt.Sprintf("%f", bidPrice), "askPrice": fmt.Sprintf("%f", askPrice)}
	if clientID != nil {
		params["clientId"] = strconv.Itoa(*clientID)
	}
	return response[*models.Quote](requestWithAuth(b, "POST", url, "quoteSubmit", params))
}

func requestWithAuth(api *BackpackREST, method, url string, instruction string, params any) (resp *resty.Response, err error) {
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
	return request(api, method, url, params, headers)
}

func request(api *BackpackREST, method, url string, params any, headers ...map[string]string) (*resty.Response, error) {
	request := api.client.R().SetHeader("User-Agent", constants.UserAgent)
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
		fn = request.Delete
	}
	return fn(url)
}

func response[T any](response *resty.Response, e error) (result T, err error) {
	if e != nil {
		return result, e
	}
	err = json.Unmarshal(response.Bytes(), &result)
	return
}

func handleError(_ *resty.Client, res *resty.Response) error {
	if res.IsError() {
		fmt.Println(res.String())
		backpackError := &BackpackError{}
		if err := json.Unmarshal(res.Bytes(), backpackError); err != nil {
			return err
		}
		return backpackError
	}
	return nil
}
