package websocket

import (
	"net/http"
	"net/url"
	"time"

	"github.com/feeeei/backpack-go/constants"
	"github.com/gorilla/websocket"
)

type options struct {
	BaseURL   string
	APIKey    string
	APISecret string
	Windows   time.Duration
}

type Options func(*options, *websocket.Dialer)

func WithBaseURL(baseURL string) Options {
	return func(o *options, d *websocket.Dialer) {
		o.BaseURL = baseURL
	}
}

func WithAPIToken(apiKey, apiSecret string) Options {
	return func(o *options, d *websocket.Dialer) {
		o.APIKey = apiKey
		o.APISecret = apiSecret
	}
}

func WithWindows(windows time.Duration) Options {
	return func(o *options, d *websocket.Dialer) {
		o.Windows = windows
	}
}

func WithProxy(proxy string) Options {
	return func(o *options, d *websocket.Dialer) {
		d.Proxy = func(req *http.Request) (*url.URL, error) {
			return url.Parse(proxy)
		}
	}
}

func DefaultWebSocketOptions() *options {
	return &options{
		BaseURL: constants.WebSocketBaseURL,
		Windows: constants.DefaultWindows,
	}
}
