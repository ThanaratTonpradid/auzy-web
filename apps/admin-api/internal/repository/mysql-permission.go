package repository

import (
	"auzy-api/model"
)

func (h Handler) CreatePermission(entity *model.Permission) error {
	return h.mysql.DB.Create(&entity).Error
}

func (h Handler) FindOnePermissionByID(permissionID uint32) (model.Permission, error) {
	entity := model.Permission{}
	err := h.mysql.DB.
		Where(&model.Permission{
			ID: permissionID,
		}, "ID").
		Where("deleted_at IS NULL").
		Take(&entity).Error
	return entity, err
}

func (h Handler) FindOnePermissionByCodeName(codeName string) (model.Permission, error) {
	entity := model.Permission{}
	err := h.mysql.DB.
		Where(&model.Permission{
			CodeName: codeName,
		}, "CodeName").
		Where("deleted_at IS NULL").
		Take(&entity).Error
	return entity, err
}

func (h Handler) FindAllPermissions() ([]model.Permission, error) {
	var entities []model.Permission
	err := h.mysql.DB.
		Where("deleted_at IS NULL").
		Order("code_name ASC").
		Find(&entities).Error
	return entities, err
}

func (h Handler) FindPermissionCodeNamesByRoleID(roleID uint32) ([]string, error) {
	var codes []string
	err := h.mysql.DB.Table("permissions").
		Select("permissions.code_name").
		Joins("INNER JOIN roles_has_permissions ON roles_has_permissions.permissions_id = permissions.id").
		Where("roles_has_permissions.roles_id = ?", roleID).
		Where("permissions.deleted_at IS NULL").
		Pluck("permissions.code_name", &codes).Error
	return codes, err
}

func (h Handler) FindPermissionsByRoleID(roleID uint32) ([]model.Permission, error) {
	var entities []model.Permission
	err := h.mysql.DB.Table("permissions").
		Select("permissions.*").
		Joins("INNER JOIN roles_has_permissions ON roles_has_permissions.permissions_id = permissions.id").
		Where("roles_has_permissions.roles_id = ?", roleID).
		Where("permissions.deleted_at IS NULL").
		Order("permissions.code_name ASC").
		Find(&entities).Error
	return entities, err
}
