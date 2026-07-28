package dto

import "auzy-api/model"

type (
	RecordVisitRequest struct {
		Path    string `json:"path" example:"/"`
		Referer string `json:"referer" example:"https://example.com"`
	}

	VisitorLogItem struct {
		ID        uint32                 `json:"id"`
		IP        string                 `json:"ip"`
		Metadata  *model.LocationMetadata `json:"metadata"`
		UserAgent *string                `json:"userAgent"`
		Path      *string                `json:"path"`
		Referer   *string                `json:"referer"`
		CreatedAt uint32                 `json:"createdAt"`
	}

	VisitorLogListResponse struct {
		Items []VisitorLogItem `json:"items"`
		Total int64            `json:"total"`
		Page  int              `json:"page"`
		Limit int              `json:"limit"`
	}
)
