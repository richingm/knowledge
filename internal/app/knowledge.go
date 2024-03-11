package app

import (
	"context"
	"github.com/richingm/knowledge/internal/biz"

	pb "github.com/richingm/knowledge/api/knowledge/v1"
)

type KnowledgeApplication struct {
	pb.UnimplementedKnowledgeServer
}

func NewKnowledgeApplication(uc *biz.KnowledgeUseCase) *KnowledgeApplication {
	return &KnowledgeApplication{}
}

func (s *KnowledgeApplication) CreateKnowledge(ctx context.Context, req *pb.CreateKnowledgeRequest) (*pb.CreateKnowledgeReply, error) {
	return &pb.CreateKnowledgeReply{}, nil
}
func (s *KnowledgeApplication) UpdateKnowledge(ctx context.Context, req *pb.UpdateKnowledgeRequest) (*pb.UpdateKnowledgeReply, error) {
	return &pb.UpdateKnowledgeReply{}, nil
}
func (s *KnowledgeApplication) DeleteKnowledge(ctx context.Context, req *pb.DeleteKnowledgeRequest) (*pb.DeleteKnowledgeReply, error) {
	return &pb.DeleteKnowledgeReply{}, nil
}
func (s *KnowledgeApplication) GetKnowledge(ctx context.Context, req *pb.GetKnowledgeRequest) (*pb.GetKnowledgeReply, error) {
	return &pb.GetKnowledgeReply{Id: req.GetId()}, nil
}
func (s *KnowledgeApplication) ListKnowledge(ctx context.Context, req *pb.ListKnowledgeRequest) (*pb.ListKnowledgeReply, error) {
	return &pb.ListKnowledgeReply{}, nil
}
