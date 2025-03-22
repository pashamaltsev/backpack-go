# backpack-go

A Go SDK for Backpack Exchange, API Docs: <a href="https://docs.backpack.exchange/" target="_blank">backpack API docs</a>

> Support Backpack API version: `2025-03-19`

- ✅ Public REST API
- ✅ Authenticated APIs Example
- ✅ Public Websocket API
- ✅ Authenticated Websocket API

## Installation

To use this SDK in your Go project, you can add it as a dependency with:

```bash
go get -u github.com/feeeei/backpack-go
```

## Usage

### Initialize the REST client

```go
package main

import (
    "fmt"
    
    "github.com/feeeei/backpack-go"
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
        backpackgo.WithAPIToken("YOUR_API_KEY", "YOUR_API_SECRET"),
        backpackgo.WithTimeout(10 * time.Second),
        backpackgo.WithRetry(3),
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
klines, err := client.GetKlines("BTC_USDC", models.OneHour, startTime, endTime)
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

// Place a limit order
order, err := clientWithAuth.ExecuteLimitOrder(
    "BTC_USDC",
    models.Buy,
    0.001, // quantity
    50000, // price
)
if err != nil {
    // Handle error
}
fmt.Printf("Order placed: %+v\n", order)

// Get open orders
symbol := "BTC_USDC"
marketType := models.Spot
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
    
    "github.com/feeeei/backpack-go"
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
        backpackgo.WithAPIToken("YOUR_API_KEY", "YOUR_API_SECRET"),
        backpackgo.WithTimeout(10 * time.Second),
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

## License
feeeei/backpack-go is released under the [MIT License](https://opensource.org/licenses/MIT).