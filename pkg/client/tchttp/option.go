package tchttp

import (
	"fmt"
	"net/http"
	"net/url"
	"time"
)

// Option 定义配置函数
type Option func(*http.Transport, *http.Client)

func WithRequestTimeout(seconds int) Option {
	return func(tr *http.Transport, c *http.Client) {
		c.Timeout = time.Duration(seconds) * time.Second
	}
}

func WithMaxIdleConnsPerHost(n int) Option {
	return func(tr *http.Transport, c *http.Client) {
		tr.MaxIdleConnsPerHost = n
	}
}

func WithMaxIdleConns(n int) Option {
	return func(tr *http.Transport, c *http.Client) {
		tr.MaxIdleConns = n
	}
}

func WithIdleConnTimeout(seconds int) Option {
	return func(tr *http.Transport, c *http.Client) {
		tr.IdleConnTimeout = time.Duration(seconds) * time.Second
	}
}

func WithProxy(proxyUrl string) Option {
	if proxyUrl == "" {
		return func(tr *http.Transport, c *http.Client) {}
	}

	proxy, err := url.Parse(proxyUrl)
	if err != nil {
		fmt.Printf("invalid proxy url %s: %v\n", proxyUrl, err)
		return func(tr *http.Transport, c *http.Client) {}
	}

	return func(tr *http.Transport, c *http.Client) {
		tr.Proxy = http.ProxyURL(proxy)
	}
}
