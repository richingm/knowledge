package biz

import "context"

// KnowledgeRepoIf is a repo.
type KnowledgeRepoIf interface {
	Create(context.Context, KnowledgePo) (*KnowledgePo, error)
	Update(context.Context, KnowledgePo) (*KnowledgePo, error)
	Delete(context.Context, KnowledgePo) (*KnowledgePo, error)
	Find(context.Context, int64) (*KnowledgePo, error)
	ListAll(context.Context) ([]KnowledgePo, error)
}
