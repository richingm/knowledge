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
func (r *knowledgeRepo) Find(ctx context.Context, id int64) (*biz.KnowledgePo, error) {
	var res biz.KnowledgePo
	err := r.data.DB(ctx).Model(&biz.KnowledgePo{}).Where("id = ?", id).First(&res).Error
	if err != nil {
		return nil, err
	}
	return &res, nil
}
func (r *knowledgeRepo) ListAll(context.Context) ([]biz.KnowledgePo, error) {
	return nil, nil
}
