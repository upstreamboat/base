package logger

import (
	"strings"
	"time"

	"gorm.io/gorm/logger"
)

type GormOption func(*GormZap)

// WithGormLevel 设置日志级别
// 默认 warn. 可选: silent, info, warn, error
//
// 扩展识别:
//
//	silent: "silent", "off", "disable", "no"
//	info: "info", "debug"
//	warn: "warn", "warning"
//	error: "error"
func WithGormLevel(level string) GormOption {
	return func(l *GormZap) {
		l.level = parseLevel(level)
	}
}

// WithGormSlowThreshold 设置慢查询阈值
func WithGormSlowThreshold(d time.Duration) GormOption {
	return func(l *GormZap) {
		l.slowThreshold = d
	}
}

func parseLevel(level string) logger.LogLevel {
	switch strings.ToLower(level) {
	case "silent", "off", "disable", "no":
		return logger.Silent
	case "error":
		return logger.Error
	case "warn", "warning":
		return logger.Warn
	case "info", "debug":
		return logger.Info
	default:
		// ⚠️ 非法值兜底策略
		return logger.Warn
	}
}
