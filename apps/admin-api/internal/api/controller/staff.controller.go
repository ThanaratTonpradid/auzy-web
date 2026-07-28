package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"mini-api/internal/api/service"
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

//	@Tags		Srtaff
//	@Summary	Get staff by id
//	@Security	Bearer
//	@Produce	json
//	@Success	200	{object}	dto.GetStaffByIdResponse	"OK"
//	@Failure	400	{object}	dto.ErrorValidationResponse	"Bad Request"
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
