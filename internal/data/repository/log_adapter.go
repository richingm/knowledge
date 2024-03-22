package infrastructure

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/utils"
	"time"
)

type GormLogAdapter struct {
	Logger log.Logger
}

func NewGormLogAdapter(logger log.Logger) logger.Interface {
	return &GormLogAdapter{Logger: logger}
}

func (g *GormLogAdapter) LogMode(logger.LogLevel) logger.Interface {
	// Kratos的log.Logger接口没有提供设置日志级别的方法
	// 可以在这里创建一个新的GormLogAdapter实例，如果需要根据不同的日志级别做特殊处理
	return g
}

func (g *GormLogAdapter) Info(ctx context.Context, msg string, data ...interface{}) {
	log.NewHelper(g.Logger).Info(fmt.Sprintf(msg, data...))
}

func (g *GormLogAdapter) Warn(ctx context.Context, msg string, data ...interface{}) {
	log.NewHelper(g.Logger).Warn(fmt.Sprintf(msg, data...))
}

func (g *GormLogAdapter) Error(ctx context.Context, msg string, data ...interface{}) {
	log.NewHelper(g.Logger).Error(fmt.Sprintf(msg, data...))
}

func (g *GormLogAdapter) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	elapsed := time.Since(begin)
	sql, rows := fc()
	tracedMsg := fmt.Sprintf("%s [%s]", utils.FileWithLineNum(), sql)
	if err != nil {
		log.NewHelper(g.Logger).Error(tracedMsg, "err", err, "elapsed", float64(elapsed.Nanoseconds())/1e6, "rows", rows)
	} else {
		log.NewHelper(g.Logger).Info(tracedMsg, "elapsed", float64(elapsed.Nanoseconds())/1e6, "rows", rows)
	}
}
