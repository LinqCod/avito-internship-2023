package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/linqcod/avito-internship-2023/internal/handler/dto"
	"github.com/linqcod/avito-internship-2023/internal/model"
)

//TODO: add database errors

const (
	CreateUserQuery  = `INSERT INTO users (username) VALUES ($1) RETURNING id;`
	GetUserByIdQuery = `SELECT id, username FROM users WHERE id=$1;`
	GetAllUsersQuery = `SELECT id, username FROM users;`

	AddUserToSegment      = `INSERT INTO users_segments (user_id, slug, ttl) VALUES ($1, $2, $3);`
	DeleteUserFromSegment = `DELETE FROM users_segments WHERE user_id=$1 AND slug=$2;`
	GetUserActiveSegments = `SELECT user_id, slug, ttl FROM users_segments 
                          		WHERE user_id=$1 AND (ttl IS NULL OR ttl > CURRENT_TIMESTAMP);`
	CheckIfUserHasSegment = `SELECT user_id, slug, ttl FROM users_segments WHERE user_id=$1 AND slug=$2;`
)

type UserRepository struct {
	ctx         context.Context
	db          *sql.DB
	historyRepo HistoryRepository
}

func NewUserRepository(ctx context.Context, db *sql.DB, historyRepo HistoryRepository) *UserRepository {
	return &UserRepository{
		ctx:         ctx,
		db:          db,
		historyRepo: historyRepo,
	}
}

func (u UserRepository) CreateUser(user dto.CreateUserDTO) (int64, error) {
	var id int64

	err := u.db.QueryRowContext(
		u.ctx,
		CreateUserQuery,
		user.Username,
	).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (u UserRepository) GetUserById(id int64) (*model.User, error) {
	var user model.User

	if err := u.db.QueryRowContext(u.ctx, GetUserByIdQuery, id).Scan(
		&user.Id,
		&user.Username,
	); err != nil {
		return nil, err
	}

	return &user, nil
}

func (u UserRepository) GetAllUsers() ([]*model.User, error) {
	var users []*model.User

	rows, err := u.db.QueryContext(u.ctx, GetAllUsersQuery)
	if err != nil {
		return nil, err
	}
	if rows.Err() != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user model.User
		if err := rows.Scan(&user.Id, &user.Username); err != nil {
			return nil, err
		}
		users = append(users, &user)
	}

	return users, nil
}

func (u UserRepository) checkIfUserHasSegment(userId int64, slug string) bool {
	var userSegment model.UserSegment

	err := u.db.QueryRowContext(u.ctx, CheckIfUserHasSegment, userId, slug).Scan(&userSegment.UserId, &userSegment.Slug, &userSegment.TTL)

	return !errors.Is(err, sql.ErrNoRows)
}

func (u UserRepository) AddUserToSegment(userId int64, slug string, ttl string) error {
	var user model.User

	err := u.db.QueryRowContext(u.ctx, GetUserByIdQuery, userId).Scan(&user.Id, &user.Username)
	if err != nil {
		return fmt.Errorf("error: user with specified id not found")
	}

	if ok := u.checkIfUserHasSegment(userId, slug); ok {
		return nil
	}

	_, err = u.db.ExecContext(
		u.ctx,
		AddUserToSegment,
		userId,
		slug,
		ttl,
	)
	if err != nil {
		return err
	}

	if err = u.historyRepo.SaveUserSegmentHistoryRecord(userId, slug, model.AddType); err != nil {
		return err
	}

	return nil
}

func (u UserRepository) DeleteUserFromSegment(userId int64, slug string) error {
	var user model.User

	err := u.db.QueryRowContext(u.ctx, GetUserByIdQuery, userId).Scan(&user.Id, &user.Username)
	if err != nil {
		return fmt.Errorf("error: user with specified id not found")
	}

	if ok := u.checkIfUserHasSegment(userId, slug); !ok {
		return nil
	}

	_, err = u.db.ExecContext(
		u.ctx,
		DeleteUserFromSegment,
		userId,
		slug,
	)
	if err != nil {
		return err
	}

	if err = u.historyRepo.SaveUserSegmentHistoryRecord(userId, slug, model.RemoveType); err != nil {
		return err
	}

	return nil
}

func (u UserRepository) GetUserActiveSegments(userId int64) ([]*model.UserSegment, error) {
	var user model.User

	err := u.db.QueryRowContext(u.ctx, GetUserByIdQuery, userId).Scan(&user.Id, &user.Username)
	if err != nil {
		return nil, fmt.Errorf("error: user with specified id not found")
	}

	var segments []*model.UserSegment

	rows, err := u.db.QueryContext(u.ctx, GetUserActiveSegments, userId)
	if err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var segment model.UserSegment
		if err := rows.Scan(&segment.UserId, &segment.Slug, &segment.TTL); err != nil {
			return nil, err
		}
		segments = append(segments, &segment)
	}

	return segments, nil
}
