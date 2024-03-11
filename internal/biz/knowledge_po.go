package biz

import (
	"time"
)

type KnowledgePo struct {
	Id              int64
	CreatedAt       time.Time
	UpdatedAt       time.Time
	Pid             int64
	Name            string
	ImportanceLevel string
	Notes           string
}

func (KnowledgePo) TableName() string {
	return "knowledges"
}

func (k *KnowledgePo) ConvertToPo(do KnowledgeDo) *KnowledgePo {
	return k
}
