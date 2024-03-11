package biz

import (
	"context"
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
	s.repo.Create(ctx, KnowledgePo{})
	return nil, nil
}
func (s *KnowledgeUseCase) UpdateKnowledge(ctx context.Context, do KnowledgeDo) (*KnowledgeDo, error) {
	return nil, nil
}
func (s *KnowledgeUseCase) DeleteKnowledge(ctx context.Context, do KnowledgeDo) (*KnowledgeDo, error) {
	return nil, nil
}
func (s *KnowledgeUseCase) GetKnowledge(ctx context.Context, req *pb.GetKnowledgeRequest) (*KnowledgeDo, error) {
	return nil, nil
}
func (s *KnowledgeUseCase) ListKnowledge(ctx context.Context, req *pb.ListKnowledgeRequest) ([]KnowledgeDo, error) {
	return nil, nil
}
