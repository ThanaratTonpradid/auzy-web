package service

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/dollarsignteam/go-logger"

	"mini-api/helper"
	"mini-api/internal/api/constant"
	"mini-api/internal/api/dto"
	"mini-api/internal/repository"
	"mini-api/lib"
	"mini-api/model"
)

type AuthService struct {
	logger     *logger.Logger
	repository *repository.Handler
	jwtHandler *lib.JWTHandler
}

func NewAuthService(
	logger *logger.Logger,
	repository *repository.Handler,
	jwtHandler *lib.JWTHandler,
) AuthService {
	return AuthService{
		logger:     logger,
		repository: repository,
		jwtHandler: jwtHandler,
	}
}

func (svc AuthService) Login(req *dto.LoginRequest, ip string) (dto.LoginResponse, error) {
	// Sanitize username input
	sanitizedUsername := helper.SanitizeUsername(req.Username)
	if sanitizedUsername == "" || len(sanitizedUsername) < 3 {
		return dto.LoginResponse{}, lib.CommonError{
			StatusCode:    http.StatusUnauthorized,
			ErrorCode:     constant.CodeLoginFailed,
			ErrorInstance: errors.New("invalid username format"),
		}
	}
	
	staff, err := svc.repository.FindOneStaffByUsername(sanitizedUsername)
	if err != nil {
		return dto.LoginResponse{}, lib.CommonError{
			StatusCode:    http.StatusUnauthorized,
			ErrorCode:     constant.CodeLoginFailed,
			ErrorInstance: err,
		}
	}
	isValid := helper.CompareHashPassword(staff.Password, req.Password)
	if !isValid {
		return dto.LoginResponse{}, lib.CommonError{
			StatusCode:    http.StatusUnauthorized,
			ErrorCode:     constant.CodeLoginFailed,
			ErrorInstance: errors.New("invalid credentials"),
		}
	}

	return svc.LoginByStaff(staff, ip)
}

func (svc AuthService) LoginByStaff(staff model.Staff, ip string) (dto.LoginResponse, error) {
	if !staff.IsActive || staff.DeletedAt != nil {
		return dto.LoginResponse{}, lib.CommonError{
			StatusCode:    http.StatusUnauthorized,
			ErrorCode:     constant.CodeLoginFailed,
			ErrorInstance: errors.New("staff is deleted"),
		}
	}
	subject := fmt.Sprintf("%d", staff.ID)
	jwtToken, err := svc.jwtHandler.CreateToken(subject)
	if err != nil {
		svc.logger.Error(GetStaffErrorMessage(staff.ID, err))
		return dto.LoginResponse{}, lib.CommonError{
			StatusCode:    http.StatusUnauthorized,
			ErrorCode:     constant.CodeCreateTokenFailed,
			ErrorInstance: err,
		}
	}
	
	refreshToken, err := svc.jwtHandler.CreateRefreshToken(subject)
	if err != nil {
		svc.logger.Error(GetStaffErrorMessage(staff.ID, err))
		return dto.LoginResponse{}, lib.CommonError{
			StatusCode:    http.StatusUnauthorized,
			ErrorCode:     constant.CodeCreateTokenFailed,
			ErrorInstance: err,
		}
	}
	
	session := dto.Session{
		ID:           jwtToken.ID,
		Username:     staff.Username,
		StaffID:      staff.ID,
		CreatedAt:    jwtToken.IssuedAt,
		RefreshToken: refreshToken.Token,
	}
	if err := svc.repository.SetStaffSession(session); err != nil {
		svc.logger.Error(GetStaffErrorMessage(staff.ID, err))
		return dto.LoginResponse{}, lib.CommonError{
			StatusCode:    http.StatusUnauthorized,
			ErrorCode:     constant.CodeCreateSessionFailed,
			ErrorInstance: err,
		}
	}
	if err := svc.repository.UpdateStaffLastLoginByID(staff.ID, ip); err != nil {
		svc.logger.Warn(GetStaffErrorMessage(staff.ID, err))
	}

	return dto.LoginResponse{
		Token:        jwtToken.Token,
		RefreshToken: refreshToken.Token,
		ExpiresIn:    jwtToken.ExpiresAt - jwtToken.IssuedAt,
	}, nil
}

func (svc AuthService) GetSession(staffId uint32) (dto.Session, error) {
	session, err := svc.repository.GetStaffSession(staffId)
	if err != nil {
		return dto.Session{}, err
	}
	if err := helper.ValidateStruct(session); err != nil {
		return dto.Session{}, err
	}
	return session, nil
}

func (svc AuthService) RefreshToken(req *dto.RefreshTokenRequest) (dto.RefreshTokenResponse, error) {
	claims, err := svc.jwtHandler.VerifyToken(req.RefreshToken)
	if err != nil {
		return dto.RefreshTokenResponse{}, lib.CommonError{
			StatusCode:    http.StatusUnauthorized,
			ErrorCode:     constant.CodeLoginFailed,
			ErrorInstance: errors.New("invalid refresh token"),
		}
	}
	
	// Get staff from subject
	var staffID uint32
	fmt.Sscanf(claims.Subject, "%d", &staffID)
	
	// Verify refresh token matches stored session
	session, err := svc.repository.GetStaffSession(staffID)
	if err != nil || session.RefreshToken != req.RefreshToken {
		return dto.RefreshTokenResponse{}, lib.CommonError{
			StatusCode:    http.StatusUnauthorized,
			ErrorCode:     constant.CodeLoginFailed,
			ErrorInstance: errors.New("invalid or expired refresh token"),
		}
	}
	
	// Get staff
	staff, err := svc.repository.FindOneStaffByID(staffID)
	if err != nil {
		return dto.RefreshTokenResponse{}, lib.CommonError{
			StatusCode:    http.StatusUnauthorized,
			ErrorCode:     constant.CodeLoginFailed,
			ErrorInstance: err,
		}
	}
	
	if !staff.IsActive || staff.DeletedAt != nil {
		return dto.RefreshTokenResponse{}, lib.CommonError{
			StatusCode:    http.StatusUnauthorized,
			ErrorCode:     constant.CodeLoginFailed,
			ErrorInstance: errors.New("staff is inactive"),
		}
	}
	
	// Create new tokens
	subject := fmt.Sprintf("%d", staff.ID)
	newJWTToken, err := svc.jwtHandler.CreateToken(subject)
	if err != nil {
		return dto.RefreshTokenResponse{}, lib.CommonError{
			StatusCode:    http.StatusUnauthorized,
			ErrorCode:     constant.CodeCreateTokenFailed,
			ErrorInstance: err,
		}
	}
	
	newRefreshToken, err := svc.jwtHandler.CreateRefreshToken(subject)
	if err != nil {
		return dto.RefreshTokenResponse{}, lib.CommonError{
			StatusCode:    http.StatusUnauthorized,
			ErrorCode:     constant.CodeCreateTokenFailed,
			ErrorInstance: err,
		}
	}
	
	// Update session
	session.ID = newJWTToken.ID
	session.RefreshToken = newRefreshToken.Token
	session.CreatedAt = newJWTToken.IssuedAt
	
	if err := svc.repository.SetStaffSession(session); err != nil {
		return dto.RefreshTokenResponse{}, lib.CommonError{
			StatusCode:    http.StatusUnauthorized,
			ErrorCode:     constant.CodeCreateSessionFailed,
			ErrorInstance: err,
		}
	}
	
	return dto.RefreshTokenResponse{
		Token:        newJWTToken.Token,
		RefreshToken: newRefreshToken.Token,
		ExpiresIn:    newJWTToken.ExpiresAt - newJWTToken.IssuedAt,
	}, nil
}

func (svc AuthService) Logout(session dto.Session) {
	svc.repository.DelStaffSession(session.StaffID)
}
