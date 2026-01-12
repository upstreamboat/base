package logger

import (
	"context"
	"time"

	"go.uber.org/zap"
	"gorm.io/gorm/logger"
)

// NewGormZap 创建 gorm zap 日志, 数据库初始化时使用
func NewGormZap(opts ...GormOption) logger.Interface {
	gz := &GormZap{
		zap:           log,
		level:         logger.Warn,            // 默认生产级
		slowThreshold: 200 * time.Millisecond, // 默认慢 SQL
	}

	for _, opt := range opts {
		opt(gz)
	}

	return gz
}

type GormZap struct {
	zap           *zap.Logger
	level         logger.LogLevel
	slowThreshold time.Duration
}

func (g *GormZap) LogMode(level logger.LogLevel) logger.Interface {
	g.level = level
	return g
}

func (g *GormZap) Info(_ context.Context, msg string, args ...any) {
	if g.level >= logger.Info {
		g.zap.Sugar().Infof(msg, args...)
	}
}

func (g *GormZap) Warn(_ context.Context, msg string, args ...any) {
	if g.level >= logger.Warn {
		g.zap.Sugar().Warnf(msg, args...)
	}
}

func (g *GormZap) Error(_ context.Context, msg string, args ...any) {
	if g.level >= logger.Error {
		g.zap.Sugar().Errorf(msg, args...)
	}
}

func (g *GormZap) Trace(_ context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	if g.level == logger.Silent {
		return
	}

	elapsed := time.Since(begin)
	sql, rows := fc()

	if sql == "SHOW STATUS" {
		return
	}

	fields := []zap.Field{
		zap.String("sql", sql),
		zap.Int64("rows", rows),
		zap.Duration("cost", elapsed),
	}

	switch {
	case err != nil && g.level >= logger.Error:
		g.zap.Error("Gorm error", append(fields, zap.Error(err))...)

	case g.slowThreshold > 0 &&
		elapsed > g.slowThreshold &&
		g.level >= logger.Warn:
		g.zap.Warn("Gorm slow sql", fields...)

	case g.level >= logger.Info:
		g.zap.Debug("Gorm sql", fields...)
	}
}
