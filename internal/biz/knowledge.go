package biz

import (
	"context"
	"errors"
	"github.com/go-kratos/kratos/v2/log"
	pb "github.com/richingm/knowledge/api/knowledge/v1"
)

type KnowledgeUseCase struct {
	repo KnowledgeRepoIf
}

func NewKnowledgeUseCase(repo KnowledgeRepoIf, logger log.Logger) *KnowledgeUseCase {
	return &KnowledgeUseCase{
		repo: repo,
	}
}

func (s *KnowledgeUseCase) CreateKnowledge(ctx context.Context, do KnowledgeDo) (*KnowledgeDo, error) {
	po := &KnowledgePo{}
	do.ConvertToPo(po)
	po, err := s.repo.Create(ctx, *po)
	if err != nil {
		return nil, err
	}
	po.ConvertToDo(&do)
	return &do, nil
}

func (s *KnowledgeUseCase) UpdateKnowledge(ctx context.Context, do KnowledgeDo) (*KnowledgeDo, error) {
	po := &KnowledgePo{}
	do.ConvertToPo(po)
	poOld, err := s.repo.Find(ctx, do.Id)
	if err != nil {
		return nil, err
	}
	if poOld == nil {
		return nil, errors.New("不存在")
	}
	po.CreatedAt = poOld.CreatedAt
	po, err = s.repo.Update(ctx, *po)
	if err != nil {
		return nil, err
	}
	po.ConvertToDo(&do)
	return &do, nil
}

func (s *KnowledgeUseCase) DeleteKnowledge(ctx context.Context, id int64) error {
	err := s.repo.Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func (s *KnowledgeUseCase) GetKnowledge(ctx context.Context, id int64) (*KnowledgeDo, error) {
	po, err := s.repo.Find(ctx, id)
	if err != nil {
		return nil, err
	}
	do := &KnowledgeDo{}
	po.ConvertToDo(do)
	return do, nil
}

func (s *KnowledgeUseCase) ListKnowledge(ctx context.Context, req *pb.ListKnowledgeRequest) ([]KnowledgeDo, error) {
	return nil, nil
}
