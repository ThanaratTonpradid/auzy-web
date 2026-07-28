package model

const TableNameVisitorLog = "visitor_logs"

// VisitorLog stores visitor IP and location metadata.
type VisitorLog struct {
	ID        uint32   `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	IP        string   `gorm:"column:ip;size:64;not null;index" json:"ip"`
	Country   *string  `gorm:"column:country;size:100" json:"country"`
	Region    *string  `gorm:"column:region;size:100" json:"region"`
	City      *string  `gorm:"column:city;size:100" json:"city"`
	Latitude  *float64 `gorm:"column:latitude" json:"latitude"`
	Longitude *float64 `gorm:"column:longitude" json:"longitude"`
	UserAgent *string  `gorm:"column:user_agent;size:512" json:"user_agent"`
	Path      *string  `gorm:"column:path;size:255" json:"path"`
	Referer   *string  `gorm:"column:referer;size:512" json:"referer"`
	CreatedAt uint32   `gorm:"column:created_at;not null;index" json:"created_at"`
}

func (*VisitorLog) TableName() string {
	return TableNameVisitorLog
}
