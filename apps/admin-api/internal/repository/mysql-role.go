package repository

import (
	"auzy-api/model"
)

func (h Handler) CreateRole(entity *model.Role) error {
	return h.mysql.DB.Create(&entity).Error
}

func (h Handler) CreateRoleHasPermissions(entity *model.RolesHasPermission) error {
	return h.mysql.DB.Create(&entity).Error
}

func (h Handler) FindOneRoleByID(roleID uint32) (model.Role, error) {
	entity := model.Role{}
	err := h.mysql.DB.
		Where(&model.Role{
			ID: roleID,
		}, "ID").
		Where("deleted_at IS NULL").
		Take(&entity).Error
	return entity, err
}

func (h Handler) FindOneRoleByLabel(label string) (model.Role, error) {
	entity := model.Role{}
	err := h.mysql.DB.
		Where(&model.Role{
			Label: label,
		}, "Label").
		Where("deleted_at IS NULL").
		Take(&entity).Error
	return entity, err
}

func (h Handler) FindAllRoles() ([]model.Role, error) {
	var entities []model.Role
	err := h.mysql.DB.
		Where("deleted_at IS NULL").
		Order("id ASC").
		Find(&entities).Error
	return entities, err
}

func (h Handler) DeleteRolePermissionsByRoleID(roleID uint32) error {
	return h.mysql.DB.
		Where("roles_id = ?", roleID).
		Delete(&model.RolesHasPermission{}).Error
}
