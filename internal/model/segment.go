package model

import "database/sql"

type Segment struct {
	Slug        string `json:"slug"`
	Description string `json:"description,omitempty"`
}

type UserSegment struct {
	UserId int64          `json:"user_id"`
	Slug   string         `json:"slug"`
	TTL    sql.NullString `json:"ttl"`
}
