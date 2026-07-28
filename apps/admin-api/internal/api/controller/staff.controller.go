package controller

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	"auzy-api/helper"
	"auzy-api/internal/api/dto"
	"auzy-api/internal/api/service"
)

type StaffController struct {
	staffService service.StaffsService
}

func NewStaffController(
	staffService service.StaffsService,
) StaffController {
	return StaffController{
		staffService: staffService,
	}
}

//	@Tags		Staff
//	@Summary	Get current staff profile
//	@Security	Bearer
//	@Produce	json
//	@Success	200	{object}	dto.GetStaffByIdResponse	"OK"
//	@Failure	401	{object}	dto.ErrorResponse			"Unauthorized"
//	@Router		/api/staff/profile [get]
func (ctrl StaffController) GetStaffProfile(c echo.Context) error {
	session := GetSession(c)
	resp, err := ctrl.staffService.GetStaffById(session.StaffID)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, resp)
}

//	@Tags		Staff
//	@Summary	List staff
//	@Security	Bearer
//	@Produce	json
//	@Success	200	{object}	dto.StaffListResponse	"OK"
//	@Failure	401	{object}	dto.ErrorResponse		"Unauthorized"
//	@Failure	403	{object}	dto.ErrorResponse		"Forbidden"
//	@Router		/api/staff [get]
func (ctrl StaffController) ListStaff(c echo.Context) error {
	resp, err := ctrl.staffService.ListStaffs()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, resp)
}

//	@Tags		Staff
//	@Summary	Get staff by id
//	@Security	Bearer
//	@Produce	json
//	@Param		id	path		int							true	"Staff ID"
//	@Success	200	{object}	dto.GetStaffByIdResponse	"OK"
//	@Failure	404	{object}	dto.ErrorResponse			"Not Found"
//	@Router		/api/staff/{id} [get]
func (ctrl StaffController) GetStaff(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return NewCommonErrorBadRequest(err)
	}
	resp, err := ctrl.staffService.GetStaffById(uint32(id))
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, resp)
}

//	@Tags		Staff
//	@Summary	Create staff
//	@Security	Bearer
//	@Accept		json
//	@Produce	json
//	@Param		data	body		dto.CreateStaffRequest		true	"Request payload"
//	@Success	201		{object}	dto.GetStaffByIdResponse	"Created"
//	@Failure	400		{object}	dto.ErrorResponse			"Bad Request"
//	@Router		/api/staff [post]
func (ctrl StaffController) CreateStaff(c echo.Context) error {
	req := new(dto.CreateStaffRequest)
	if err := c.Bind(req); err != nil {
		return NewCommonErrorBadRequest(err)
	}
	if err := helper.ValidateStruct(req); err != nil {
		return NewCommonErrorBadRequest(err)
	}
	resp, err := ctrl.staffService.CreateStaffFromRequest(req)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, resp)
}

//	@Tags		Staff
//	@Summary	Update staff
//	@Security	Bearer
//	@Accept		json
//	@Produce	json
//	@Param		id		path		int							true	"Staff ID"
//	@Param		data	body		dto.UpdateStaffRequest		true	"Request payload"
//	@Success	200		{object}	dto.GetStaffByIdResponse	"OK"
//	@Failure	400		{object}	dto.ErrorResponse			"Bad Request"
//	@Router		/api/staff/{id} [put]
func (ctrl StaffController) UpdateStaff(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return NewCommonErrorBadRequest(err)
	}
	req := new(dto.UpdateStaffRequest)
	if err := c.Bind(req); err != nil {
		return NewCommonErrorBadRequest(err)
	}
	if err := helper.ValidateStruct(req); err != nil {
		return NewCommonErrorBadRequest(err)
	}
	resp, err := ctrl.staffService.UpdateStaff(uint32(id), req)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, resp)
}

//	@Tags		Staff
//	@Summary	Delete staff
//	@Security	Bearer
//	@Produce	json
//	@Param		id	path		int					true	"Staff ID"
//	@Success	200	{object}	dto.SuccessResponse	"OK"
//	@Failure	400	{object}	dto.ErrorResponse	"Bad Request"
//	@Router		/api/staff/{id} [delete]
func (ctrl StaffController) DeleteStaff(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return NewCommonErrorBadRequest(err)
	}
	session := GetSession(c)
	if err := ctrl.staffService.DeleteStaff(uint32(id), session.StaffID); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, dto.SuccessResponse{
		Success: true,
		Message: "Staff deleted",
	})
}
