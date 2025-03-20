package constants

import "time"

const (
	UserAgent = "backpack-go/0.0.1"

	RESTBaseURL      = `https://api.backpack.exchange`
	WebSocketBaseURL = `wss://ws.backpack.exchange`

	Windows = time.Duration(time.Second * 10)
)
