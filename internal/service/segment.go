package service

import (
	"github.com/linqcod/avito-internship-2023/internal/handler/dto"
)

type SegmentRepository interface {
	CreateSegment(segment dto.CreateSegmentDTO) (string, error)
	DeleteSegment(slug string) error
}

type SegmentService struct {
	repo SegmentRepository
}

func NewSegmentService(repo SegmentRepository) *SegmentService {
	return &SegmentService{
		repo: repo,
	}
}

func (s SegmentService) CreateSegment(segment dto.CreateSegmentDTO) (string, error) {
	return s.repo.CreateSegment(segment)
}

func (s SegmentService) DeleteSegment(slug string) error {
	return s.repo.DeleteSegment(slug)
}
