package biz

import "time"

type KnowledgeDo struct {
	Id              int64
	CreatedAt       time.Time
	UpdatedAt       time.Time
	Pid             int64
	Name            string
	ImportanceLevel string
	Notes           string
}

func (k *KnowledgeDo) ConvertToPo(po *KnowledgePo) {
	po.Id = k.Id
	po.CreatedAt = k.CreatedAt
	po.UpdatedAt = k.UpdatedAt
	po.Pid = k.Pid
	po.Name = k.Name
	po.ImdortanceLevel = k.ImportanceLevel
	po.Notes = k.Notes
}
