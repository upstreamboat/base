package logger

import (
	"time"

	"github.com/upstreamboat/base/pkg/logger/internal"
)

// Option 日志配置项
type Option func(*internal.Zap)

// WithLevel 设置日志级别
// 默认 info, 可选 debug, info, warn, error, dpanic, panic, fatal
func WithLevel(level string) Option {
	return func(z *internal.Zap) {
		if level == "" {
			return
		}
		z.Level = level
	}
}

// WithPrefix 设置日志前缀
// 默认 ""
func WithPrefix(prefix string) Option {
	return func(z *internal.Zap) {
		if prefix == "" {
			return
		}
		z.Prefix = prefix
	}
}

// WithFormat 输出格式
// 默认 console
func WithFormat(format string) Option {
	return func(z *internal.Zap) {
		if format == "" {
			return
		}
		z.Format = format
	}
}

// WithDirector 日志目录
// 默认 logs
func WithDirector(director string) Option {
	return func(z *internal.Zap) {
		if director == "" {
			return
		}
		z.Director = director
	}
}

// WithEncodeLevel 日志编码级别
// 默认 LowercaseColorLevelEncoder
func WithEncodeLevel(encodeLevel string) Option {
	return func(z *internal.Zap) {
		if encodeLevel == "" {
			return
		}
		z.EncodeLevel = encodeLevel
	}
}

// WithStacktraceKey 栈名
// 默认 "stacktrace"
func WithStacktraceKey(stacktraceKey string) Option {
	return func(z *internal.Zap) {
		if stacktraceKey == "" {
			return
		}
		z.StacktraceKey = stacktraceKey
	}
}

// WithShowLine 显示行号
// 默认 true
func WithShowLine(showLine *bool) Option {
	return func(z *internal.Zap) {
		if showLine == nil {
			return
		}
		z.ShowLine = *showLine
	}
}

// WithLogInConsole 是否输出到控制台
// 默认 true
func WithLogInConsole(logInConsole *bool) Option {
	return func(z *internal.Zap) {
		if logInConsole == nil {
			return
		}
		z.LogInConsole = *logInConsole
	}
}

// WithRetentionDay 日志保留天数
// 默认 30, 小于0(比如-1)为永久保留
func WithRetentionDay(retentionDay int) Option {
	return func(z *internal.Zap) {
		if retentionDay == 0 {
			return
		}
		z.RetentionDay = retentionDay
	}
}

func WithSqlLevel(sqlLevel string) Option {
	return func(z *internal.Zap) {
		if sqlLevel == "" {
			return
		}
		z.SqlLevel = sqlLevel
	}
}

func WithSqlSlowTime(slowTime time.Duration) Option {
	return func(z *internal.Zap) {
		if slowTime == 0 {
			return
		}
		z.SqlSlowTime = slowTime
	}
}
