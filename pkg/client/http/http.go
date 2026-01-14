package http

import (
	"net/http"
	"net/url"
	"time"
)

func InitHttp(timeout time.Duration, proxyUrl string) *http.Client {
	if proxyUrl != "" {
		proxy, err := url.Parse(proxyUrl)
		if err != nil {
			panic("failed to parse proxy" + err.Error())
		}

		// 有代理
		return &http.Client{
			Timeout: timeout,
			Transport: &http.Transport{
				MaxIdleConns:    100,
				IdleConnTimeout: 90 * time.Second,
				Proxy:           http.ProxyURL(proxy),
			},
		}
	}

	// 无代理
	return &http.Client{
		Timeout: timeout,
		Transport: &http.Transport{
			MaxIdleConns:    100,
			IdleConnTimeout: 90 * time.Second,
		},
	}
}
