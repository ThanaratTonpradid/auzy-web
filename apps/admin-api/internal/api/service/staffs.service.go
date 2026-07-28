package service

import (
	"errors"
	"net/http"

	"github.com/dollarsignteam/go-logger"

	"auzy-api/config"
	"auzy-api/helper"
	"auzy-api/internal/api/constant"
	"auzy-api/internal/api/dto"
	"auzy-api/internal/repository"
	"auzy-api/lib"
	"auzy-api/model"
)

type StaffInit struct {
	RolesName string
	Username  string
	Password  string
	Fullname  string
	IsAdmin   bool
}

type StaffsService struct {
	logger     *logger.Logger
	repository *repository.Handler
	jwtHandler *lib.JWTHandler
	cfg        *config.APIConfig
}

func NewStaffsService(
	logger *logger.Logger,
	repository *repository.Handler,
	jwtHandler *lib.JWTHandler,
	cfg *config.APIConfig,
) StaffsService {
	return StaffsService{
		logger:     logger,
		repository: repository,
		jwtHandler: jwtHandler,
		cfg:        cfg,
	}
}

func (svc StaffsService) InitStaffs() {
	existingStaff, err := svc.repository.FindOneStaffByUsername("admin")
	if err == nil && existingStaff.ID > 0 {
		svc.logger.Info("Admin staff already exists, skipping...")
		return
	}

	svc.logger.Info("Creating admin staff...")
	staff := StaffInit{
		RolesName: constant.RoleAdmin,
		Username:  "admin",
		Password:  svc.cfg.DefaultDevPassword,
		Fullname:  "Administrator",
		IsAdmin:   true,
	}
	_, err = svc.CreateStaff(staff)
	if err != nil {
		svc.logger.Errorf("Failed to create admin staff: %v", err)
	} else {
		svc.logger.Info("Admin staff created successfully")
	}
}

func (svc StaffsService) GetStaffById(staffId uint32) (dto.GetStaffByIdResponse, error) {
	staff, err := svc.repository.FindOneStaffByID(staffId)
	if err != nil {
		svc.logger.Error(err)
		return dto.GetStaffByIdResponse{}, lib.CommonError{
			StatusCode:    http.StatusNotFound,
			ErrorCode:     constant.ErrCodeNotFound,
			ErrorInstance: err,
		}
	}
	return svc.toStaffResponse(staff)
}

func (svc StaffsService) ListStaffs() (dto.StaffListResponse, error) {
	staffs, err := svc.repository.FindAllStaffs()
	if err != nil {
		svc.logger.Error(err)
		return NewCommonErrorSomethingWentWrong(dto.StaffListResponse{}, err)
	}
	total, err := svc.repository.CountStaffs()
	if err != nil {
		svc.logger.Error(err)
		return NewCommonErrorSomethingWentWrong(dto.StaffListResponse{}, err)
	}

	items := make([]dto.StaffListItem, 0, len(staffs))
	for _, staff := range staffs {
		fullname := ""
		if staff.Fullname != nil {
			fullname = *staff.Fullname
		}
		items = append(items, dto.StaffListItem{
			ID:        staff.ID,
			Username:  staff.Username,
			Fullname:  fullname,
			RoleID:    staff.RolesID,
			RoleLabel: staff.RoleLabel,
			IsActive:  staff.IsActive,
			IsAdmin:   staff.IsAdmin,
			LastLogin: staff.LastLogin,
			LastIP:    staff.LastIP,
		})
	}
	return dto.StaffListResponse{Items: items, Total: total}, nil
}

func (svc StaffsService) CreateStaffFromRequest(req *dto.CreateStaffRequest) (dto.GetStaffByIdResponse, error) {
	if _, err := svc.repository.FindOneStaffByUsername(req.Username); err == nil {
		return dto.GetStaffByIdResponse{}, lib.CommonError{
			StatusCode:    http.StatusBadRequest,
			ErrorCode:     constant.ErrCodeBadRequest,
			ErrorInstance: errors.New("username already exists"),
		}
	}

	role, err := svc.repository.FindOneRoleByID(req.RoleID)
	if err != nil {
		return dto.GetStaffByIdResponse{}, lib.CommonError{
			StatusCode:    http.StatusBadRequest,
			ErrorCode:     constant.ErrCodeBadRequest,
			ErrorInstance: errors.New("role not found"),
		}
	}

	isActive := true
	if req.IsActive != nil {
		isActive = *req.IsActive
	}

	password, err := helper.GenerateHashPassword(req.Password)
	if err != nil {
		svc.logger.Error(err)
		return NewCommonErrorSomethingWentWrong(dto.GetStaffByIdResponse{}, err)
	}

	unixTimeNow := GetUnixTimestamp()
	fullname := req.Fullname
	entity := model.Staff{
		RolesID:   role.ID,
		Username:  req.Username,
		Password:  password,
		Fullname:  &fullname,
		IsActive:  isActive,
		IsAdmin:   req.IsAdmin,
		CreatedAt: unixTimeNow,
		UpdatedAt: unixTimeNow,
	}
	if err := svc.repository.CreateStaff(&entity); err != nil {
		svc.logger.Error(err)
		return NewCommonErrorSomethingWentWrong(dto.GetStaffByIdResponse{}, err)
	}
	return svc.GetStaffById(entity.ID)
}

func (svc StaffsService) UpdateStaff(staffID uint32, req *dto.UpdateStaffRequest) (dto.GetStaffByIdResponse, error) {
	if _, err := svc.repository.FindOneStaffByID(staffID); err != nil {
		return dto.GetStaffByIdResponse{}, lib.CommonError{
			StatusCode:    http.StatusNotFound,
			ErrorCode:     constant.ErrCodeNotFound,
			ErrorInstance: err,
		}
	}

	updates := map[string]interface{}{
		"updated_at": GetUnixTimestamp(),
	}
	if req.Fullname != nil {
		updates["fullname"] = *req.Fullname
	}
	if req.IsActive != nil {
		updates["is_active"] = *req.IsActive
	}
	if req.IsAdmin != nil {
		updates["is_admin"] = *req.IsAdmin
	}
	if req.RoleID != nil {
		if _, err := svc.repository.FindOneRoleByID(*req.RoleID); err != nil {
			return dto.GetStaffByIdResponse{}, lib.CommonError{
				StatusCode:    http.StatusBadRequest,
				ErrorCode:     constant.ErrCodeBadRequest,
				ErrorInstance: errors.New("role not found"),
			}
		}
		updates["roles_id"] = *req.RoleID
	}
	if req.Password != nil && *req.Password != "" {
		password, err := helper.GenerateHashPassword(*req.Password)
		if err != nil {
			svc.logger.Error(err)
			return NewCommonErrorSomethingWentWrong(dto.GetStaffByIdResponse{}, err)
		}
		updates["password"] = password
	}

	if err := svc.repository.UpdateStaff(staffID, updates); err != nil {
		svc.logger.Error(err)
		return NewCommonErrorSomethingWentWrong(dto.GetStaffByIdResponse{}, err)
	}
	return svc.GetStaffById(staffID)
}

func (svc StaffsService) DeleteStaff(staffID uint32, actorID uint32) error {
	if staffID == actorID {
		return lib.CommonError{
			StatusCode:    http.StatusBadRequest,
			ErrorCode:     constant.ErrCodeBadRequest,
			ErrorInstance: errors.New("cannot delete your own account"),
		}
	}
	if _, err := svc.repository.FindOneStaffByID(staffID); err != nil {
		return lib.CommonError{
			StatusCode:    http.StatusNotFound,
			ErrorCode:     constant.ErrCodeNotFound,
			ErrorInstance: err,
		}
	}
	if err := svc.repository.SoftDeleteStaffByID(staffID); err != nil {
		svc.logger.Error(err)
		return lib.CommonError{
			StatusCode:    http.StatusInternalServerError,
			ErrorCode:     constant.ErrCodeInternalError,
			ErrorInstance: err,
		}
	}
	svc.repository.DelStaffSession(staffID)
	return nil
}

func (svc StaffsService) CreateStaff(req StaffInit) (model.Staff, error) {
	role, err := svc.repository.FindOneRoleByLabel(req.RolesName)
	if err != nil {
		svc.logger.Error(err)
		return NewCommonErrorSomethingWentWrong(model.Staff{}, err)
	}
	unixTimeNow := GetUnixTimestamp()
	password, err := helper.GenerateHashPassword(req.Password)
	if err != nil {
		svc.logger.Error(err)
		return NewCommonErrorSomethingWentWrong(model.Staff{}, err)
	}
	fullname := req.Fullname
	entity := model.Staff{
		RolesID:   role.ID,
		Username:  req.Username,
		Password:  password,
		Fullname:  &fullname,
		IsActive:  true,
		IsAdmin:   req.IsAdmin,
		CreatedAt: unixTimeNow,
		UpdatedAt: unixTimeNow,
	}
	if err := svc.repository.CreateStaff(&entity); err != nil {
		svc.logger.Error(err)
		return NewCommonErrorSomethingWentWrong(model.Staff{}, err)
	}
	return svc.repository.FindOneStaffByID(entity.ID)
}

func (svc StaffsService) toStaffResponse(staff model.Staff) (dto.GetStaffByIdResponse, error) {
	fullname := ""
	if staff.Fullname != nil {
		fullname = *staff.Fullname
	}
	roleLabel := ""
	if role, err := svc.repository.FindOneRoleByID(staff.RolesID); err == nil {
		roleLabel = role.Label
	}
	permissions, err := svc.repository.FindPermissionCodeNamesByRoleID(staff.RolesID)
	if err != nil {
		permissions = []string{}
	}
	if permissions == nil {
		permissions = []string{}
	}
	return dto.GetStaffByIdResponse{
		StaffId:     staff.ID,
		Username:    staff.Username,
		Fullname:    fullname,
		RolesID:     staff.RolesID,
		RoleLabel:   roleLabel,
		IsActive:    staff.IsActive,
		IsAdmin:     staff.IsAdmin,
		Permissions: permissions,
	}, nil
}
