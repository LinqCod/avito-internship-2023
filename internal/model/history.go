package model

import (
	"time"
)

const (
	AddType    = "ADD"
	RemoveType = "REMOVE"
)

type UserSegmentHistory struct {
	UserId     int64     `json:"user_id"`
	Slug       string    `json:"slug"`
	ActionType string    `json:"action_type"`
	ActionTime time.Time `json:"action_time"`
}
