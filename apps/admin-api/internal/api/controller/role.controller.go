package controller

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	"auzy-api/helper"
	"auzy-api/internal/api/dto"
	"auzy-api/internal/api/service"
)

type RoleController struct {
	roleService       service.RolesService
	permissionService service.PermissionsService
}

func NewRoleController(
	roleService service.RolesService,
	permissionService service.PermissionsService,
) RoleController {
	return RoleController{
		roleService:       roleService,
		permissionService: permissionService,
	}
}

//	@Tags		Roles
//	@Summary	List roles
//	@Security	Bearer
//	@Produce	json
//	@Success	200	{object}	dto.RoleListResponse	"OK"
//	@Router		/api/roles [get]
func (ctrl RoleController) ListRoles(c echo.Context) error {
	resp, err := ctrl.roleService.ListRoles()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, resp)
}

//	@Tags		Roles
//	@Summary	Get role by id
//	@Security	Bearer
//	@Produce	json
//	@Param		id	path		int						true	"Role ID"
//	@Success	200	{object}	dto.RoleDetailResponse	"OK"
//	@Router		/api/roles/{id} [get]
func (ctrl RoleController) GetRole(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return NewCommonErrorBadRequest(err)
	}
	resp, err := ctrl.roleService.GetRoleByID(uint32(id))
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, resp)
}

//	@Tags		Roles
//	@Summary	Update role permissions
//	@Security	Bearer
//	@Accept		json
//	@Produce	json
//	@Param		id		path		int									true	"Role ID"
//	@Param		data	body		dto.UpdateRolePermissionsRequest	true	"Request payload"
//	@Success	200		{object}	dto.RoleDetailResponse				"OK"
//	@Router		/api/roles/{id}/permissions [put]
func (ctrl RoleController) UpdateRolePermissions(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return NewCommonErrorBadRequest(err)
	}
	req := new(dto.UpdateRolePermissionsRequest)
	if err := c.Bind(req); err != nil {
		return NewCommonErrorBadRequest(err)
	}
	if err := helper.ValidateStruct(req); err != nil {
		return NewCommonErrorBadRequest(err)
	}
	resp, err := ctrl.roleService.UpdateRolePermissions(uint32(id), req)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, resp)
}

//	@Tags		Permissions
//	@Summary	List permissions
//	@Security	Bearer
//	@Produce	json
//	@Success	200	{object}	dto.PermissionListResponse	"OK"
//	@Router		/api/permissions [get]
func (ctrl RoleController) ListPermissions(c echo.Context) error {
	resp, err := ctrl.permissionService.ListPermissions()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, resp)
}
