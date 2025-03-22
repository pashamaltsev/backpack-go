module github.com/feeeei/backpack-go

go 1.24.0

require (
	github.com/gorilla/websocket v1.5.3
	github.com/json-iterator/go v1.1.12
	github.com/liamylian/jsontime/v3 v3.0.0
	resty.dev/v3 v3.0.0-beta.2
)

require (
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	golang.org/x/net v0.36.0 // indirect
)

replace github.com/liamylian/jsontime/v3 => github.com/feeeei/jsontime/v3 v3.0.0-20250320154044-5e60ccddd437
