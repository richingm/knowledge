package migrations

import (
	"context"
	"github.com/richingm/knowledge/internal/biz"
	"gorm.io/gorm"
)

func InitMigrate(ctx context.Context, db *gorm.DB) error {
	err := db.AutoMigrate(biz.KnowledgePo{})
	if err != nil {
		return err
	}
	return nil
}
