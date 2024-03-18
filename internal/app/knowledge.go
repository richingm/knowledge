package app

import (
	"context"
	"github.com/richingm/knowledge/internal/biz"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"

	pb "github.com/richingm/knowledge/api/knowledge/v1"
)

type KnowledgeApplication struct {
	pb.UnimplementedKnowledgeServer
	uc *biz.KnowledgeUseCase
}

func NewKnowledgeApplication(uc *biz.KnowledgeUseCase) *KnowledgeApplication {
	return &KnowledgeApplication{
		uc: uc,
	}
}

func (s *KnowledgeApplication) CreateKnowledge(ctx context.Context, req *pb.CreateKnowledgeRequest) (*pb.CreateKnowledgeReply, error) {
	knowledgeDo, err := s.uc.CreateKnowledge(ctx, biz.KnowledgeDo{
		CreatedAt:       time.Time{},
		UpdatedAt:       time.Time{},
		Pid:             req.Pid,
		Name:            req.Name,
		ImportanceLevel: req.ImportanceLevel,
		Notes:           req.Notes,
	})
	if err != nil {
		return nil, err
	}
	reply := pb.CreateKnowledgeReply{
		Id:              knowledgeDo.Id,
		CreatedAt:       timestamppb.New(knowledgeDo.CreatedAt),
		UpdatedAt:       timestamppb.New(knowledgeDo.UpdatedAt),
		Pid:             knowledgeDo.Pid,
		Name:            knowledgeDo.Name,
		ImportanceLevel: knowledgeDo.ImportanceLevel,
		Notes:           knowledgeDo.Notes,
	}
	return &reply, nil
}
func (s *KnowledgeApplication) UpdateKnowledge(ctx context.Context, req *pb.UpdateKnowledgeRequest) (*pb.UpdateKnowledgeReply, error) {
	knowledgeDo, err := s.uc.UpdateKnowledge(ctx, biz.KnowledgeDo{
		Id:              req.Id,
		CreatedAt:       time.Time{},
		UpdatedAt:       time.Time{},
		Pid:             req.Pid,
		Name:            req.Name,
		ImportanceLevel: req.ImportanceLevel,
		Notes:           req.Notes,
	})
	if err != nil {
		return nil, err
	}
	repy := pb.UpdateKnowledgeReply{
		Id:              knowledgeDo.Id,
		CreatedAt:       timestamppb.New(knowledgeDo.CreatedAt),
		UpdatedAt:       timestamppb.New(knowledgeDo.UpdatedAt),
		Pid:             knowledgeDo.Pid,
		Name:            knowledgeDo.Name,
		ImportanceLevel: knowledgeDo.ImportanceLevel,
		Notes:           knowledgeDo.Notes,
	}
	return &repy, nil
}
func (s *KnowledgeApplication) DeleteKnowledge(ctx context.Context, req *pb.DeleteKnowledgeRequest) (*pb.DeleteKnowledgeReply, error) {
	err := s.uc.DeleteKnowledge(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.DeleteKnowledgeReply{}, nil
}
func (s *KnowledgeApplication) GetKnowledge(ctx context.Context, req *pb.GetKnowledgeRequest) (*pb.GetKnowledgeReply, error) {
	do, err := s.uc.GetKnowledge(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.GetKnowledgeReply{Id: do.Id, Name: do.Name}, nil
}
func (s *KnowledgeApplication) ListKnowledge(ctx context.Context, req *pb.ListKnowledgeRequest) (*pb.ListKnowledgeReply, error) {
	return &pb.ListKnowledgeReply{}, nil
}
