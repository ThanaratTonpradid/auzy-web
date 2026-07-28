package repository

import (
	"auzy-api/model"
)

type StaffWithRole struct {
	model.Staff
	RoleLabel string `gorm:"column:role_label"`
}

func (h Handler) CreateStaff(entity *model.Staff) error {
	return h.mysql.DB.Create(&entity).Error
}

func (h Handler) FindOneStaffByID(staffID uint32) (model.Staff, error) {
	entity := model.Staff{}
	if err := h.mysql.DB.
		Where(&model.Staff{
			ID: staffID,
		}, "ID").
		Where("deleted_at IS NULL").
		First(&entity).Error; err != nil {
		return model.Staff{}, err
	}
	return entity, nil
}

func (h Handler) FindOneStaffByUsername(username string) (model.Staff, error) {
	entity := model.Staff{}
	if err := h.mysql.DB.
		Where(&model.Staff{
			Username: username,
		}, "Username").
		Where("deleted_at IS NULL").
		First(&entity).Error; err != nil {
		return model.Staff{}, err
	}
	return entity, nil
}

func (h Handler) FindAllStaffs() ([]StaffWithRole, error) {
	var entities []StaffWithRole
	err := h.mysql.DB.Table("staffs").
		Select("staffs.*, roles.label as role_label").
		Joins("LEFT JOIN roles ON roles.id = staffs.roles_id").
		Where("staffs.deleted_at IS NULL").
		Order("staffs.id ASC").
		Find(&entities).Error
	return entities, err
}

func (h Handler) CountStaffs() (int64, error) {
	var count int64
	err := h.mysql.DB.Model(&model.Staff{}).
		Where("deleted_at IS NULL").
		Count(&count).Error
	return count, err
}

func (h Handler) UpdateStaff(staffID uint32, updates map[string]interface{}) error {
	return h.mysql.DB.Model(&model.Staff{}).
		Where("id = ? AND deleted_at IS NULL", staffID).
		Updates(updates).Error
}

func (h Handler) SoftDeleteStaffByID(staffID uint32) error {
	unixTimeNow := GetUnixTimestamp()
	return h.mysql.DB.Model(&model.Staff{}).
		Where("id = ? AND deleted_at IS NULL", staffID).
		Updates(map[string]interface{}{
			"deleted_at": unixTimeNow,
			"updated_at": unixTimeNow,
			"is_active":  false,
		}).Error
}

func (h Handler) UpdateStaffLastLoginByID(staffID uint32, ip string) error {
	unixTimeNow := GetUnixTimestamp()
	return h.mysql.DB.
		Where(&model.Staff{
			ID: staffID,
		}, "ID").
		Updates(model.Staff{
			LastIP:    &ip,
			LastLogin: &unixTimeNow,
			UpdatedAt: unixTimeNow,
		}).Error
}
