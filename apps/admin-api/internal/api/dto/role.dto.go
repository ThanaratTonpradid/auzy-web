package dto

type (
	PermissionItem struct {
		ID       uint32 `json:"id"`
		CodeName string `json:"codeName"`
	}

	RoleListItem struct {
		ID          uint32   `json:"id"`
		Label       string   `json:"label"`
		Permissions []string `json:"permissions"`
	}

	RoleDetailResponse struct {
		ID              uint32           `json:"id"`
		Label           string           `json:"label"`
		Permissions     []string         `json:"permissions"`
		PermissionItems []PermissionItem `json:"permissionItems"`
	}

	RoleListResponse struct {
		Items []RoleListItem `json:"items"`
	}

	UpdateRolePermissionsRequest struct {
		PermissionIDs []uint32 `json:"permissionIds" validate:"required"`
	}

	PermissionListResponse struct {
		Items []PermissionItem `json:"items"`
	}
)
