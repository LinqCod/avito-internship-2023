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
	GetUserIdsByPercentageQuery = `
									SELECT id FROM users
									ORDER BY random()
									LIMIT (SELECT count(DISTINCT id) FROM users) * $1 / 100;
	`
	AddSegmentToUserQuery = `INSERT INTO users_segments (user_id, slug) VALUES ($1, $2);`

	CreateSegmentQuery    = `INSERT INTO segments (slug, description) VALUES ($1, $2) RETURNING slug;`
	DeleteSegmentQuery    = `DELETE FROM segments WHERE slug=$1;`
	GetSegmentBySlugQuery = `SELECT slug, description FROM segments WHERE slug=$1;`
)

type SegmentRepository struct {
	ctx         context.Context
	db          *sql.DB
	historyRepo HistoryRepository
}

func NewSegmentRepository(ctx context.Context, db *sql.DB, historyRepo HistoryRepository) *SegmentRepository {
	return &SegmentRepository{
		ctx:         ctx,
		db:          db,
		historyRepo: historyRepo,
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

		if segment.Percentage != 0 {
			var userIds []int64

			rows, err := s.db.QueryContext(s.ctx, GetUserIdsByPercentageQuery, segment.Percentage)
			if err != nil {
				return slug, err
			}
			if rows.Err() != nil {
				return slug, err
			}
			defer rows.Close()

			for rows.Next() {
				var userId int64
				if err := rows.Scan(&userId); err != nil {
					return slug, err
				}
				userIds = append(userIds, userId)
			}

			for _, id := range userIds {
				if _, err := s.db.ExecContext(s.ctx, AddSegmentToUserQuery, id, slug); err != nil {
					return slug, err
				}

				if err = s.historyRepo.SaveUserSegmentHistoryRecord(id, slug, model.AddType); err != nil {
					return slug, err
				}
			}
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
