package biz

import (
	"time"
)

type KnowledgeDo struct {
	Id              int64
	CreatedAt       time.Time
	UpdatedAt       time.Time
	Pid             int64
	Name            string
	ImportanceLevel string
	Notes           string
}

func (K KnowledgeDo) ConvertToDo() KnowledgePo {
	return KnowledgePo{}
}
