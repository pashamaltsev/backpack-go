package backpackgo

import (
	"time"

	"github.com/feeeei/backpack-go/constants"
)

type options struct {
	BaseURL   string
	APIKey    string
	APISecret string
	Windows   time.Duration
}

type Options func(*options)

func WithAPIToken(apiKey, apiSecret string) Options {
	return func(o *options) {
		o.APIKey = apiKey
		o.APISecret = apiSecret
	}
}

func WithBaseURL(baseURL string) Options {
	return func(o *options) {
		o.BaseURL = baseURL
	}
}

func WithWindows(windows time.Duration) Options {
	return func(o *options) {
		o.Windows = windows
	}
}

func defaultRESTOptions() *options {
	return &options{
		BaseURL: constants.RESTBaseURL,
		Windows: constants.Windows,
	}
}

func defaultWebSocketOptions() *options {
	return &options{
		BaseURL: constants.WebSocketBaseURL,
	}
}
