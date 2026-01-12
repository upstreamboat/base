package logger

import (
	"sync"

	"github.com/upstreamboat/base/pkg/logger/internal"
	"go.uber.org/zap"
)

var (
	log      *zap.Logger
	initOnce sync.Once
	cfg      *internal.Zap
)

// InitLog 初始化日志
// opts 可修改日志配置, 不修改则使用默认配置. 配置详情见 logger.Option
func InitLog(opts ...Option) *zap.Logger {
	initOnce.Do(func() {
		cfg = internal.NewConfig()
		for _, opt := range opts {
			opt(cfg)
		}
		log = internal.NewZap(cfg)
		zap.ReplaceGlobals(log)
	})

	return log
}

// L 获取日志实例
func L() *zap.Logger {
	return log
}

func C() *internal.Zap {
	return cfg
}
