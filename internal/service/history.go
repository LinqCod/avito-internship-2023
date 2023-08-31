package service

import (
	"encoding/csv"
	"fmt"
	"github.com/linqcod/avito-internship-2023/internal/model"
	"os"
	"path/filepath"
	"strconv"
)

type HistoryRepository interface {
	SaveUserSegmentHistoryRecord(userId int64, slug string, actionType string) error
	GetUserSegmentHistory(userId int64, month int64, year int64) ([]*model.UserSegmentHistory, error)
}

type HistoryService struct {
	repo HistoryRepository
}

func NewHistoryService(repo HistoryRepository) *HistoryService {
	return &HistoryService{
		repo: repo,
	}
}

func (s HistoryService) GetUserSegmentHistory(userId int64, month int64, year int64) error {
	history, err := s.repo.GetUserSegmentHistory(userId, month, year)
	if err != nil {
		return err
	}

	if err = createCSVHistoryFile(userId, history); err != nil {
		return err
	}

	return nil
}

func createCSVHistoryFile(userId int64, history []*model.UserSegmentHistory) error {
	historyRecords := make([][]string, 0)
	for _, r := range history {
		historyRecords = append(
			historyRecords,
			[]string{strconv.Itoa(int(r.UserId)),
				r.Slug,
				r.ActionType,
				r.ActionTime.String()},
		)
	}

	path, err := filepath.Abs("")
	if err != nil {
		return err
	}
	filename := fmt.Sprintf("%d.csv", userId)

	f, err := os.Create(path + "/" + filename)
	defer f.Close()

	if err != nil {
		return err
	}

	w := csv.NewWriter(f)
	defer w.Flush()

	for _, record := range historyRecords {
		if err := w.Write(record); err != nil {
			return err
		}
	}

	return nil
}
