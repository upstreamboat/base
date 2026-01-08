package logger

import (
	"sync"
	"sync/atomic"

	"github.com/upstreamboat/base/pkg/logger/internal"

	"go.uber.org/zap"
)

var (
	log    *zap.Logger
	config *internal.Zap

	initOnce sync.Once
	isInit   atomic.Bool
)

func InitLog(zapConfig *internal.Zap) *zap.Logger {
	initOnce.Do(func() {
		if zapConfig == nil {
			zapConfig = internal.NewConfig()
		}
		config = zapConfig
		log = internal.NewZap(config)
		zap.ReplaceGlobals(log)
		isInit.Store(true)
	})

	return log
}

func Log() *zap.Logger {
	if isInit.Load() {
		return log
	}

	// 兜底初始化（仅在完全未初始化时）
	initOnce.Do(func() {
		config = internal.NewConfig()
		log = internal.NewZap(config)
		zap.ReplaceGlobals(log)
		isInit.Store(true)
	})

	// 提示未初始化
	log.Warn(
		"未初始化日志, 已回退使用默认配置, 建议调用 InitGlobal() 明确初始化",
		zap.String("pkg", "go-logger"),
	)

	return log
}

func Cfg() *internal.Zap {
	return config
}
