package dto

type (
	RecordVisitRequest struct {
		Path    string `json:"path" example:"/"`
		Referer string `json:"referer" example:"https://example.com"`
	}

	VisitorLogItem struct {
		ID        uint32   `json:"id"`
		IP        string   `json:"ip"`
		Country   *string  `json:"country"`
		Region    *string  `json:"region"`
		City      *string  `json:"city"`
		Latitude  *float64 `json:"latitude"`
		Longitude *float64 `json:"longitude"`
		UserAgent *string  `json:"userAgent"`
		Path      *string  `json:"path"`
		Referer   *string  `json:"referer"`
		CreatedAt uint32   `json:"createdAt"`
	}

	VisitorLogListResponse struct {
		Items []VisitorLogItem `json:"items"`
		Total int64            `json:"total"`
		Page  int              `json:"page"`
		Limit int              `json:"limit"`
	}
)
