package dto

import "github.com/linqcod/avito-internship-2023/internal/model"

type CreateSegmentDTO struct {
	Slug        string `json:"slug"`
	Description string `json:"description,omitempty"`
}

type CreateSegmentResponse struct {
	Slug string `json:"slug"`
}

type ChangeUserSegmentsDTO struct {
	SegmentsToAdd    []SegmentWithTTL `json:"segments_to_add"`
	SegmentsToRemove []string         `json:"segments_to_remove"`
}

type SegmentWithTTL struct {
	Slug string `json:"slug"`
	TTL  string `json:"ttl,omitempty"`
}

type ActiveUserSegmentsDTO struct {
	UserId   int64             `json:"user_id"`
	Segments []*SegmentWithTTL `json:"segments"`
}

func ConvertUserSegmentToActiveUserSegments(userId int64, userSegments []*model.UserSegment) *ActiveUserSegmentsDTO {
	var segments []*SegmentWithTTL

	for _, segm := range userSegments {
		segment := &SegmentWithTTL{
			Slug: segm.Slug,
			TTL:  segm.TTL.String,
		}
		segments = append(segments, segment)
	}

	// TODO: fix?
	if userSegments == nil {
		segments = []*SegmentWithTTL{}
	}

	return &ActiveUserSegmentsDTO{
		UserId:   userId,
		Segments: segments,
	}
}
