package repository

import (
	"fmt"

	"auzy-api/model"
)

func GetKeyGeoIP(ip string) string {
	return fmt.Sprintf("%s:%s", KeyGeoIP, ip)
}

func (h Handler) CreateVisitorLog(entity *model.VisitorLog) error {
	return h.mysql.DB.Create(entity).Error
}

func (h Handler) ListVisitorLogs(offset, limit int) ([]model.VisitorLog, error) {
	var entities []model.VisitorLog
	err := h.mysql.DB.
		Order("id DESC").
		Offset(offset).
		Limit(limit).
		Find(&entities).Error
	return entities, err
}

func (h Handler) CountVisitorLogs() (int64, error) {
	var count int64
	err := h.mysql.DB.Model(&model.VisitorLog{}).Count(&count).Error
	return count, err
}

func (h Handler) DeleteVisitorLogsOlderThan(unixTimestamp uint32) (int64, error) {
	result := h.mysql.DB.
		Where("created_at < ?", unixTimestamp).
		Delete(&model.VisitorLog{})
	return result.RowsAffected, result.Error
}
