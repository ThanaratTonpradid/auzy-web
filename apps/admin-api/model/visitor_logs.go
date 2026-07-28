package model

const TableNameVisitorLog = "visitor_logs"

// LocationMetadata is stored as JSON in visitor_logs.metadata.
type LocationMetadata struct {
	Country   string   `json:"country,omitempty"`
	Region    string   `json:"region,omitempty"`
	City      string   `json:"city,omitempty"`
	Latitude  *float64 `json:"latitude,omitempty"`
	Longitude *float64 `json:"longitude,omitempty"`
	Source    string   `json:"source,omitempty"`
}

// VisitorLog stores visitor IP and location metadata.
type VisitorLog struct {
	ID        uint32            `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	IP        string            `gorm:"column:ip;size:64;not null;index" json:"ip"`
	Metadata  *LocationMetadata `gorm:"column:metadata;type:json;serializer:json" json:"metadata"`
	UserAgent *string           `gorm:"column:user_agent;size:512" json:"user_agent"`
	Path      *string           `gorm:"column:path;size:255" json:"path"`
	Referer   *string           `gorm:"column:referer;size:512" json:"referer"`
	CreatedAt uint32            `gorm:"column:created_at;not null;index" json:"created_at"`
}

func (*VisitorLog) TableName() string {
	return TableNameVisitorLog
}
