package service

import (
	"errors"
	"net/http"

	"github.com/dollarsignteam/go-logger"

	"auzy-api/internal/api/constant"
	"auzy-api/internal/api/dto"
	"auzy-api/internal/repository"
	"auzy-api/lib"
	"auzy-api/model"
)

type RoleDefault struct {
	Label       string
	Permissions []string
}

type RolesService struct {
	logger     *logger.Logger
	repository *repository.Handler
	jwtHandler *lib.JWTHandler
}

func NewRolesService(
	logger *logger.Logger,
	repository *repository.Handler,
	jwtHandler *lib.JWTHandler,
) RolesService {
	return RolesService{
		logger:     logger,
		repository: repository,
		jwtHandler: jwtHandler,
	}
}

func (svc RolesService) InitRoles() {
	existingRole, err := svc.repository.FindOneRoleByLabel(constant.RoleAdmin)
	if err == nil && existingRole.ID > 0 {
		svc.logger.Info("Roles already initialized, skipping...")
		return
	}

	roles := []RoleDefault{
		{
			Label: constant.RoleAdmin,
			Permissions: []string{
				constant.RolesRead,
				constant.RolesCreate,
				constant.RolesUpdate,
				constant.RolesDelete,
				constant.StaffsRead,
				constant.StaffsCreate,
				constant.StaffsUpdate,
				constant.StaffsDelete,
			},
		},
		{
			Label: constant.RoleStaff,
			Permissions: []string{
				constant.RolesRead,
				constant.StaffsRead,
			},
		},
		{
			Label: constant.RoleMember,
			Permissions: []string{
				constant.RolesRead,
				constant.StaffsRead,
			},
		},
	}
	svc.logger.Info("Start init role")
	for _, r := range roles {
		svc.logger.Infof("Insert: %s", r.Label)
		createdRole, _ := svc.CreateRole(r.Label)
		for _, p := range r.Permissions {
			findPermission, _ := svc.repository.FindOnePermissionByCodeName(p)
			svc.CreateRoleHasPermissions(createdRole.ID, findPermission.ID)
		}
	}
	svc.logger.Info("Init role complete")
}

func (svc RolesService) ListRoles() (dto.RoleListResponse, error) {
	roles, err := svc.repository.FindAllRoles()
	if err != nil {
		svc.logger.Error(err)
		return NewCommonErrorSomethingWentWrong(dto.RoleListResponse{}, err)
	}
	items := make([]dto.RoleListItem, 0, len(roles))
	for _, role := range roles {
		permissions, err := svc.repository.FindPermissionCodeNamesByRoleID(role.ID)
		if err != nil {
			permissions = []string{}
		}
		if permissions == nil {
			permissions = []string{}
		}
		items = append(items, dto.RoleListItem{
			ID:          role.ID,
			Label:       role.Label,
			Permissions: permissions,
		})
	}
	return dto.RoleListResponse{Items: items}, nil
}

func (svc RolesService) GetRoleByID(roleID uint32) (dto.RoleDetailResponse, error) {
	role, err := svc.repository.FindOneRoleByID(roleID)
	if err != nil {
		return dto.RoleDetailResponse{}, lib.CommonError{
			StatusCode:    http.StatusNotFound,
			ErrorCode:     constant.ErrCodeNotFound,
			ErrorInstance: err,
		}
	}
	permissions, err := svc.repository.FindPermissionsByRoleID(roleID)
	if err != nil {
		svc.logger.Error(err)
		return NewCommonErrorSomethingWentWrong(dto.RoleDetailResponse{}, err)
	}
	codes := make([]string, 0, len(permissions))
	items := make([]dto.PermissionItem, 0, len(permissions))
	for _, p := range permissions {
		codes = append(codes, p.CodeName)
		items = append(items, dto.PermissionItem{ID: p.ID, CodeName: p.CodeName})
	}
	return dto.RoleDetailResponse{
		ID:              role.ID,
		Label:           role.Label,
		Permissions:     codes,
		PermissionItems: items,
	}, nil
}

func (svc RolesService) UpdateRolePermissions(roleID uint32, req *dto.UpdateRolePermissionsRequest) (dto.RoleDetailResponse, error) {
	if _, err := svc.repository.FindOneRoleByID(roleID); err != nil {
		return dto.RoleDetailResponse{}, lib.CommonError{
			StatusCode:    http.StatusNotFound,
			ErrorCode:     constant.ErrCodeNotFound,
			ErrorInstance: err,
		}
	}

	uniqueIDs := map[uint32]struct{}{}
	for _, id := range req.PermissionIDs {
		uniqueIDs[id] = struct{}{}
	}

	for id := range uniqueIDs {
		if _, err := svc.repository.FindOnePermissionByID(id); err != nil {
			return dto.RoleDetailResponse{}, lib.CommonError{
				StatusCode:    http.StatusBadRequest,
				ErrorCode:     constant.ErrCodeBadRequest,
				ErrorInstance: errors.New("invalid permission id"),
			}
		}
	}

	if err := svc.repository.DeleteRolePermissionsByRoleID(roleID); err != nil {
		svc.logger.Error(err)
		return NewCommonErrorSomethingWentWrong(dto.RoleDetailResponse{}, err)
	}

	for id := range uniqueIDs {
		if _, err := svc.CreateRoleHasPermissions(roleID, id); err != nil {
			return NewCommonErrorSomethingWentWrong(dto.RoleDetailResponse{}, err)
		}
	}
	return svc.GetRoleByID(roleID)
}

func (svc RolesService) FindOneRoleByID(roleID uint32) (model.Role, error) {
	role, err := svc.repository.FindOneRoleByID(roleID)
	if err != nil {
		svc.logger.Error(err)
		return NewCommonErrorSomethingWentWrong(model.Role{}, err)
	}
	return role, nil
}

func (svc RolesService) FindOneRoleByLabel(label string) (model.Role, error) {
	role, err := svc.repository.FindOneRoleByLabel(label)
	if err != nil {
		svc.logger.Error(err)
		return NewCommonErrorSomethingWentWrong(model.Role{}, err)
	}
	return role, nil
}

func (svc RolesService) CreateRole(roleName string) (model.Role, error) {
	unixTimeNow := GetUnixTimestamp()
	entity := model.Role{
		Label:     roleName,
		CreatedAt: unixTimeNow,
		UpdatedAt: unixTimeNow,
	}
	if err := svc.repository.CreateRole(&entity); err != nil {
		svc.logger.Error(err)
		return NewCommonErrorSomethingWentWrong(model.Role{}, err)
	}
	return svc.FindOneRoleByID(entity.ID)
}

func (svc RolesService) CreateRoleHasPermissions(createdRoleID uint32, permissionID uint32) (model.Role, error) {
	entity := model.RolesHasPermission{
		RolesID:       createdRoleID,
		PermissionsID: permissionID,
	}
	if err := svc.repository.CreateRoleHasPermissions(&entity); err != nil {
		svc.logger.Error(err)
		return NewCommonErrorSomethingWentWrong(model.Role{}, err)
	}
	return svc.FindOneRoleByID(createdRoleID)
}
