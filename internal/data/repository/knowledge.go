package infrastructure

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/richingm/knowledge/internal/biz"
)

type knowledgeRepo struct {
	data *Data
	log  *log.Helper
}

// NewKnowledgeRepo .
func NewKnowledgeRepo(data *Data, logger log.Logger) biz.KnowledgeRepoIf {
	return &knowledgeRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *knowledgeRepo) Create(context.Context, biz.KnowledgePo) (*biz.KnowledgePo, error) {
	return nil, nil
}
func (r *knowledgeRepo) Update(context.Context, biz.KnowledgePo) (*biz.KnowledgePo, error) {
	return nil, nil
}
func (r *knowledgeRepo) Delete(context.Context, biz.KnowledgePo) (*biz.KnowledgePo, error) {
	return nil, nil
}
func (r *knowledgeRepo) Find(context.Context, int64) (*biz.KnowledgePo, error) {
	return nil, nil
}
func (r *knowledgeRepo) ListAll(context.Context) ([]biz.KnowledgePo, error) {
	return nil, nil
}
