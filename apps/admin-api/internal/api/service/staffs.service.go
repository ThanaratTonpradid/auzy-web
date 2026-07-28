package service

import (
	"github.com/dollarsignteam/go-logger"

	"mini-api/config"
	"mini-api/helper"
	"mini-api/internal/api/constant"
	"mini-api/internal/api/dto"
	"mini-api/internal/repository"
	"mini-api/lib"
	"mini-api/model"
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
	// Check if admin staff already exists
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
		return NewCommonErrorSomethingWentWrong(dto.GetStaffByIdResponse{}, err)
	}
	fullname := ""
	if staff.Fullname != nil {
		fullname = *staff.Fullname
	}
	return dto.GetStaffByIdResponse{
		StaffId:  staff.ID,
		Username: staff.Username,
		Fullname: fullname,
		RolesID:  staff.RolesID,
		IsActive: staff.IsActive,
	}, nil
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
	entity := model.Staff{
		RolesID:   role.ID,
		Username:  req.Username,
		Password:  password,
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
