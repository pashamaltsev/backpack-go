module github.com/UnipayFI/backpack-go

go 1.24.0

require (
	github.com/go-json-experiment/json v0.0.0-00010101000000-000000000000
	github.com/gorilla/websocket v1.5.3
	resty.dev/v3 v3.0.0-beta.2
)

require golang.org/x/net v0.38.0 // indirect

replace github.com/go-json-experiment/json => github.com/UnipayFI/json v0.0.0-20250427052530-1311881935d4
