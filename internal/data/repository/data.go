package infrastructure

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/richingm/knowledge/internal/biz"
	"github.com/richingm/knowledge/internal/conf"
	"github.com/richingm/knowledge/migrations"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

type Data struct {
	db *gorm.DB
}

// 用来承载事务的上下文
type contextTxKey struct{}

// NewTransaction .
func NewTransaction(d *Data) biz.Transaction {
	return d
}

// ExecTx gorm Transaction
func (d *Data) ExecTx(ctx context.Context, fn func(ctx context.Context) error) error {
	return d.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		ctx = context.WithValue(ctx, contextTxKey{}, tx)
		return fn(ctx)
	})
}

// DB 根据此方法来判断当前的 db 是不是使用 事务的 DB
func (d *Data) DB(ctx context.Context) *gorm.DB {
	tx, ok := ctx.Value(contextTxKey{}).(*gorm.DB)
	if ok {
		return tx
	}
	return d.db
}

func NewGormDB(c *conf.Data, logger log.Logger) (*gorm.DB, error) {
	dsn := c.Database.Source
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.NewHelper(logger).Info("closing the data resources")
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	sqlDB.SetMaxIdleConns(50)
	sqlDB.SetMaxOpenConns(150)
	sqlDB.SetConnMaxLifetime(time.Second * 25)
	return db, err
}

func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	d := &Data{}
	db, err := NewGormDB(c, logger)
	if err != nil {
		return d, cleanup, err
	}
	d.db = db

	var ctx context.Context
	err = migrations.InitMigrate(ctx, db)
	if err != nil {
		return d, cleanup, nil
	}

	return d, cleanup, nil
}

var ProviderSet = wire.NewSet(NewData, NewGormDB, NewKnowledgeRepo)
