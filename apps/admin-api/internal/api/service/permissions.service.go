package service

import (
	"github.com/dollarsignteam/go-logger"

	"auzy-api/internal/api/constant"
	"auzy-api/internal/api/dto"
	"auzy-api/internal/repository"
	"auzy-api/lib"
	"auzy-api/model"
)

type PermissionsService struct {
	logger     *logger.Logger
	repository *repository.Handler
	jwtHandler *lib.JWTHandler
}

func NewPermissionsService(
	logger *logger.Logger,
	repository *repository.Handler,
	jwtHandler *lib.JWTHandler,
) PermissionsService {
	return PermissionsService{
		logger:     logger,
		repository: repository,
		jwtHandler: jwtHandler,
	}
}

func (svc PermissionsService) InitPermissions() {
	permissions := []string{
		constant.RolesRead,
		constant.RolesCreate,
		constant.RolesUpdate,
		constant.RolesDelete,
		constant.StaffsRead,
		constant.StaffsCreate,
		constant.StaffsUpdate,
		constant.StaffsDelete,
		constant.VisitorsRead,
	}
	svc.logger.Info("Start ensure permissions")
	for _, code := range permissions {
		if _, err := svc.repository.FindOnePermissionByCodeName(code); err == nil {
			continue
		}
		svc.logger.Infof("Insert: %s", code)
		svc.CreatePermission(code)
	}
	svc.logger.Info("Ensure permissions complete")
}

func (svc PermissionsService) FindOnePermissionByID(permissionID uint32) (model.Permission, error) {
	permission, err := svc.repository.FindOnePermissionByID(permissionID)
	if err != nil {
		svc.logger.Error(err)
		return NewCommonErrorSomethingWentWrong(model.Permission{}, err)
	}
	return permission, nil
}

func (svc PermissionsService) CreatePermission(code string) (model.Permission, error) {
	unixTimeNow := GetUnixTimestamp()
	entity := model.Permission{
		CodeName:  code,
		CreatedAt: unixTimeNow,
		UpdatedAt: unixTimeNow,
	}
	if err := svc.repository.CreatePermission(&entity); err != nil {
		svc.logger.Error(err)
		return NewCommonErrorSomethingWentWrong(model.Permission{}, err)
	}
	return svc.FindOnePermissionByID(entity.ID)
}

func (svc PermissionsService) ListPermissions() (dto.PermissionListResponse, error) {
	permissions, err := svc.repository.FindAllPermissions()
	if err != nil {
		svc.logger.Error(err)
		return NewCommonErrorSomethingWentWrong(dto.PermissionListResponse{}, err)
	}
	items := make([]dto.PermissionItem, 0, len(permissions))
	for _, p := range permissions {
		items = append(items, dto.PermissionItem{
			ID:       p.ID,
			CodeName: p.CodeName,
		})
	}
	return dto.PermissionListResponse{Items: items}, nil
}
