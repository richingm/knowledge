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

func (k *KnowledgeDo) ConvertToDo(po KnowledgePo) *KnowledgeDo {
	k.Id = po.Id
	k.CreatedAt = po.CreatedAt
	k.UpdatedAt = po.UpdatedAt
	k.Pid = po.Pid
	k.Name = po.Name
	k.ImportanceLevel = po.ImportanceLevel
	k.Notes = po.Notes
	return k
}
