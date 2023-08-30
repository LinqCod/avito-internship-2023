package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/linqcod/avito-internship-2023/internal/handler/dto"
	"github.com/linqcod/avito-internship-2023/internal/model"
)

const (
	CreateSegmentQuery    = `INSERT INTO segments (slug, description) VALUES ($1, $2) RETURNING slug;`
	DeleteSegmentQuery    = `DELETE FROM segments WHERE slug=$1;`
	GetSegmentBySlugQuery = `SELECT slug, description FROM segments WHERE slug=$1;`
)

type SegmentRepository struct {
	ctx context.Context
	db  *sql.DB
}

func NewSegmentRepository(ctx context.Context, db *sql.DB) *SegmentRepository {
	return &SegmentRepository{
		ctx: ctx,
		db:  db,
	}
}

func (s SegmentRepository) checkIfSegmentAlreadyExists(slug string) bool {
	var segment model.Segment

	err := s.db.QueryRowContext(s.ctx, GetSegmentBySlugQuery, slug).Scan(&segment.Slug, &segment.Description)

	return !errors.Is(err, sql.ErrNoRows)
}

func (s SegmentRepository) CreateSegment(segment dto.CreateSegmentDTO) (string, error) {
	if exists := s.checkIfSegmentAlreadyExists(segment.Slug); !exists {
		var slug string

		err := s.db.QueryRowContext(
			s.ctx,
			CreateSegmentQuery,
			segment.Slug,
			segment.Description,
		).Scan(&slug)
		if err != nil {
			return "", err
		}

		return slug, nil
	}

	return "", fmt.Errorf("error while creating segment: segment with this slug already exists")
}

func (s SegmentRepository) DeleteSegment(slug string) error {

	if _, err := s.db.ExecContext(s.ctx, DeleteSegmentQuery, slug); err != nil {
		return err
	}

	return nil
}
