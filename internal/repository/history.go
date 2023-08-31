package repository

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/linqcod/avito-internship-2023/internal/model"
	"time"
)

const (
	SaveUserSegmentHistoryRecordQuery = `
							INSERT INTO users_segments_history (user_id, slug, action_type, action_time) 
							VALUES ($1, $2, $3, $4);`
	GetUserSegmentHistoryQuery = `
							SELECT 
							    user_id, 
							    slug, 
							    action_type, 
							    action_time
							FROM users_segments_history
							WHERE user_id = $1 
							  AND EXTRACT(MONTH FROM action_time) = $2 
							  AND EXTRACT(YEAR FROM action_time) = $3;`
)

type HistoryRepository interface {
	SaveUserSegmentHistoryRecord(userId int64, slug string, actionType string) error
	GetUserSegmentHistory(userId int64, month int64, year int64) ([]*model.UserSegmentHistory, error)
}

type HistoryRepositoryImpl struct {
	ctx context.Context
	db  *sql.DB
}

func NewHistoryRepository(ctx context.Context, db *sql.DB) HistoryRepository {
	return &HistoryRepositoryImpl{
		ctx: ctx,
		db:  db,
	}
}

func (r HistoryRepositoryImpl) SaveUserSegmentHistoryRecord(userId int64, slug string, actionType string) error {
	_, err := r.db.ExecContext(r.ctx, SaveUserSegmentHistoryRecordQuery, userId, slug, actionType, time.Now())
	if err != nil {
		return fmt.Errorf("error while saving user segment history record: %v", err)
	}

	return nil
}

func (r HistoryRepositoryImpl) GetUserSegmentHistory(userId int64, month int64, year int64) ([]*model.UserSegmentHistory, error) {
	var history []*model.UserSegmentHistory

	rows, err := r.db.QueryContext(r.ctx, GetUserSegmentHistoryQuery, userId, month, year)
	if err != nil {
		return nil, err
	}
	if rows.Err() != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var historyEntry model.UserSegmentHistory
		if err := rows.Scan(
			&historyEntry.UserId,
			&historyEntry.Slug,
			&historyEntry.ActionType,
			&historyEntry.ActionTime,
		); err != nil {
			return nil, err
		}
		history = append(history, &historyEntry)
	}

	return history, nil
}
