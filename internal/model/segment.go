package model

type Segment struct {
	Slug        string `json:"slug"`
	Description string `json:"description,omitempty"`
}

type UserSegment struct {
	UserId int64  `json:"user_id"`
	Slug   string `json:"slug"`
}

type CreateSegmentDTO struct {
	Slug        string `json:"slug"`
	Description string `json:"description,omitempty"`
}

type CreateSegmentResponse struct {
	Slug string `json:"slug"`
}

type ChangeUserSegmentsDTO struct {
	SegmentsToAdd    []string `json:"segments_to_add"`
	SegmentsToRemove []string `json:"segments_to_remove"`
}

type ActiveSegment struct {
	Slug string `json:"slug"`
}

type ActiveUserSegmentsDTO struct {
	UserId   int64            `json:"user_id"`
	Segments []*ActiveSegment `json:"segments"`
}

func ConvertUserSegmentToActiveUserSegments(userId int64, userSegments []*UserSegment) *ActiveUserSegmentsDTO {
	var segments []*ActiveSegment

	for _, segm := range userSegments {
		segment := &ActiveSegment{
			Slug: segm.Slug,
		}
		segments = append(segments, segment)
	}

	if userSegments == nil {
		segments = []*ActiveSegment{}
	}

	return &ActiveUserSegmentsDTO{
		UserId:   userId,
		Segments: segments,
	}
}
