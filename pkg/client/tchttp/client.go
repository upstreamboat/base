package tchttp

import (
	"net/http"
	"time"
)

var (
	// DefaultClient 是一个预配置的全局 HTTP 客户端，支持长连接复用
	DefaultClient *http.Client
)

func init() {
	// 初始化默认客户端
	DefaultClient = NewClient()
}

// NewClient 创建一个新的、经过配置的 http.Client
func NewClient(opts ...Option) *http.Client {
	tr := &http.Transport{}

	c := &http.Client{
		Transport: tr,
		Timeout:   10 * time.Second, // 整个请求的绝对超时时间（含读取 Body）
	}

	// 应用自定义配置
	for _, opt := range opts {
		if opt != nil {
			opt(tr, c)
		}
	}

	return c
}
