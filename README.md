# backpack-go

A Go SDK for Backpack Exchange, API Docs: <a href="https://docs.backpack.exchange/" target="_blank">backpack API docs</a>

> Support Backpack API version: `2025-04-22`

- ✅ Public REST API
- ✅ Authenticated APIs Example
- ✅ Public Websocket API
- ✅ Authenticated Websocket API

## Installation

To use this SDK in your Go project, you can add it as a dependency with:

```bash
go get -u github.com/pashamaltsev/backpack-go
```

## Usage

### Initialize the REST client

```go
package main

import (
	backpackgo "github.com/pashamaltsev/backpack-go"
	"github.com/pashamaltsev/backpack-go/rest"
)

func main() {
	// Create a client with default settings
	client := backpackgo.NewRESTClient()

	// Or create a client with custom settings
	// support:
	// - WithAPIToken(apiKey, apiSecret string)
	// - WithBaseURL(baseURL string)
	// - WithWindows(windows time.Duration)
	// - WithTimeout(timeout time.Duration)
	// - WithRetry(retry int)
	// - WithProxy(proxy string)
	clientWithAuth := backpackgo.NewRESTClient(
		rest.WithAPIToken("YOUR_API_KEY", "YOUR_API_SECRET"),
	)

	// Now you can use the client to call APIs
}
```

### Public REST APIs Example

```go
// Get market data
markets, err := client.GetMarkets()
if err != nil {
    // Handle error
}
fmt.Printf("Markets count: %d\n", len(markets))

// Get ticker for a specific symbol
ticker, err := client.GetTicker("BTC_USDC")
if err != nil {
    // Handle error
}
fmt.Printf("%s price: %v\n", ticker.Symbol, ticker.FirstPrice)

// Get klines (candlestick) data
endTime := time.Now()
startTime := endTime.Add(-24 * time.Hour) // Last 24 hours
klines, err := client.GetKlines("BTC_USDC", options.KLineInterval1h, startTime, endTime)
if err != nil {
    // Handle error
}
```

### Authenticated REST APIs Example

```go
// Get account information (requires authentication)
account, err := clientWithAuth.GetAccount()
if err != nil {
    // Handle error
}
fmt.Printf("Account: %+v\n", account)

// also support:
// - ExecuteMarketOrder(symbol, side, quantity)
// - ExecuteLimitOrder(symbol, side, price, quantity)
// - ExecuteConditionalLimitOrder(symbol, side, triggerPrice, price, quantity)
order, err := clientWithAuth.ExecuteLimitOrder(
    "BTC_USDC",
    options.Buy,
    0.001, // quantity
    50000, // price
)
if err != nil {
    // Handle error
}
fmt.Printf("Order placed: %+v\n", order)

// Get open orders
symbol := "BTC_USDC"
marketType := options.Spot
orders, err := clientWithAuth.GetOrders(&symbol, &marketType)
if err != nil {
    // Handle error
}
fmt.Printf("Open orders: %d\n", len(orders))
```

### Initialize the Websocket client
```go
package main

import (
    "fmt"
    
    "github.com/pashamaltsev/backpack-go"
)

func main() {
    // Create a client with default settings
    client := backpackgo.NewRESTClient()
    
    // Or create a client with custom settings
    // support:
    // - WithAPIToken(apiKey, apiSecret string)
    // - WithBaseURL(baseURL string)
    // - WithWindows(windows time.Duration)
    // - WithProxy(proxy string)
    clientWithAuth := backpackgo.NewRESTClient(
        rest.WithAPIToken("YOUR_API_KEY", "YOUR_API_SECRET"),
        rest.WithTimeout(10 * time.Second),
    )
    
    // Now you can use the client to call APIs
}
```

### Websocket APIs Example
```go
err = client.SubscribeTrade(symbol, func(trade *models.TradeUpdate) {
    fmt.Println(tade)
})
if err != nil {
    fmt.Errorf("subscribe trade faild: %+v", err)
}

_ = client.Unsubscribe(fmt.Sprintf("trade.%s", symbol+"_PERP"))

// should With API_KEY & API_SECRET
err = client.SubscribeOrderUpdate(func(order *models.OrderUpdate) {
    fmt.Println(order)
})
```

### Advanced: Custom Request

#### Custom REST Request

You can make custom REST requests with your own defined structures if you need to access endpoints that are not covered by the SDK or if you want more control over the request/response handling:

```go
// Define your custom response structure
type Order struct {
    ID        string  `json:"id"`
    Price     float64 `json:"price"`
    Quantity  float64 `json:"quantity"`
    Timestamp int64   `json:"timestamp"`
    // Add any other fields you need
}

// Make a GET request with custom path and response type
customOrderPath := "/api/v1/custom_order"
params := map[string]string{"param1": "value1", "param2": "value2"}
response, err := backpackgo.Request[*Order](client, "GET", customOrderPath, params)
if err != nil {
    // Handle error
}
fmt.Printf("Custom response: %+v\n", response)

// Make a POST request with custom path and payload
postPayload := map[string]any{
    "key1": "value1",
    "key2": 42,
}
postResponse, err := backpackgo.Request[*CustomResponse](client, "POST", customPath, postPayload)
if err != nil {
    // Handle error
}
```

#### Custom WebSocket Request

For WebSocket, you can subscribe to custom stream or handle custom message formats:

```go
// Define your custom message structure
type CustomWSMessage struct {
    EventType string  `json:"e"`
    Symbol    string  `json:"s"`
    Value     float64 `json:"v"`
    // Add any other fields you need
}

// Subscribe to a custom stream with your handler
customStream := "custom.stream.name"
err = client.SubscribePublicStream(customStream, new(CustomWSMessage), func(msg *CustomWSMessage) {
    fmt.Printf("Received custom message: %+v\n", msg)
})
if err != nil {
    fmt.Errorf("subscribe custom stream failed: %+v", err)
}

// For authenticated streams
err = client.SubscribePrivateStream(customStream, new(CustomWSMessage), func(msg *CustomWSMessage) {
    fmt.Printf("Received custom auth message: %+v\n", msg)
})
```

This flexibility allows you to work with new API endpoints or custom implementations without waiting for SDK updates.


## License
feeeei/backpack-go is released under the [MIT License](https://opensource.org/licenses/MIT).