package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"auzy-api/internal/api/dto"
	"auzy-api/internal/api/service"
)

type VisitorController struct {
	visitorService service.VisitorService
}

func NewVisitorController(
	visitorService service.VisitorService,
) VisitorController {
	return VisitorController{
		visitorService: visitorService,
	}
}

//	@Tags		Public
//	@Summary	Record a public profile visit
//	@Accept		json
//	@Produce	json
//	@Param		data	body	dto.RecordVisitRequest	false	"Visit payload"
//	@Success	204		"No Content"
//	@Router		/api/public/visit [post]
func (ctrl VisitorController) RecordVisit(c echo.Context) error {
	req := new(dto.RecordVisitRequest)
	_ = c.Bind(req)

	ip := c.RealIP()
	ua := c.Request().UserAgent()
	if err := ctrl.visitorService.RecordVisit(ip, ua, req); err != nil {
		// Do not fail the visitor experience for logging errors
		return c.NoContent(http.StatusNoContent)
	}
	return c.NoContent(http.StatusNoContent)
}

//	@Tags		Visitors
//	@Summary	List visitor logs
//	@Security	Bearer
//	@Produce	json
//	@Param		page	query		int	false	"Page number"
//	@Param		limit	query		int	false	"Page size"
//	@Success	200		{object}	dto.VisitorLogListResponse
//	@Router		/api/visitor-logs [get]
func (ctrl VisitorController) ListVisitorLogs(c echo.Context) error {
	page := service.ParsePositiveInt(c.QueryParam("page"), 1)
	limit := service.ParsePositiveInt(c.QueryParam("limit"), 20)
	resp, err := ctrl.visitorService.ListVisits(page, limit)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, resp)
}

//	@Tags		Public
//	@Summary	Lookup location for the current client IP
//	@Produce	json
//	@Success	200	{object}	dto.PublicLocationResponse
//	@Router		/api/public/location [get]
func (ctrl VisitorController) GetPublicLocation(c echo.Context) error {
	ip := c.RealIP()
	return c.JSON(http.StatusOK, ctrl.visitorService.LookupClientLocation(ip))
}
