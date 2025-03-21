package backpackgo

import (
	"time"

	"github.com/feeeei/backpack-go/constants"
	"resty.dev/v3"
)

type options struct {
	BaseURL   string
	APIKey    string
	APISecret string
	Windows   time.Duration
}

type Options func(*options, *resty.Client)

func WithAPIToken(apiKey, apiSecret string) Options {
	return func(o *options, c *resty.Client) {
		o.APIKey = apiKey
		o.APISecret = apiSecret
	}
}

func WithBaseURL(baseURL string) Options {
	return func(o *options, c *resty.Client) {
		o.BaseURL = baseURL
	}
}

func WithWindows(windows time.Duration) Options {
	return func(o *options, c *resty.Client) {
		o.Windows = windows
	}
}

func WithTimeout(timeout time.Duration) Options {
	return func(o *options, c *resty.Client) {
		c.SetTimeout(timeout)
	}
}

func WithRetry(retry int) Options {
	return func(o *options, c *resty.Client) {
		c.SetRetryCount(retry)
	}
}

func WithProxy(proxy string) Options {
	return func(o *options, c *resty.Client) {
		c.SetProxy(proxy)
	}
}

func defaultRESTOptions() *options {
	return &options{
		BaseURL: constants.RESTBaseURL,
		Windows: constants.DefaultWindows,
	}
}

func defaultWebSocketOptions() *options {
	return &options{
		BaseURL: constants.WebSocketBaseURL,
	}
}
