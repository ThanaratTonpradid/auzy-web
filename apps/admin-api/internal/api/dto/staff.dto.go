package dto

type (
	GetStaffByIdResponse struct {
		StaffId     uint32   `json:"id" example:"1"`
		Username    string   `json:"username" example:"username"`
		Fullname    string   `json:"fullname" example:"fullname"`
		RolesID     uint32   `json:"roleId" example:"1"`
		RoleLabel   string   `json:"roleLabel" example:"ADMIN"`
		IsActive    bool     `json:"isActive" example:"true"`
		IsAdmin     bool     `json:"isAdmin" example:"false"`
		Permissions []string `json:"permissions"`
	}

	StaffListItem struct {
		ID        uint32 `json:"id"`
		Username  string `json:"username"`
		Fullname  string `json:"fullname"`
		RoleID    uint32 `json:"roleId"`
		RoleLabel string `json:"roleLabel"`
		IsActive  bool   `json:"isActive"`
		IsAdmin   bool   `json:"isAdmin"`
		LastLogin *uint32 `json:"lastLogin,omitempty"`
		LastIP    *string `json:"lastIp,omitempty"`
	}

	StaffListResponse struct {
		Items []StaffListItem `json:"items"`
		Total int64           `json:"total"`
	}

	CreateStaffRequest struct {
		Username string `json:"username" validate:"required,min=3,max=50"`
		Password string `json:"password" validate:"required,min=6"`
		Fullname string `json:"fullname" validate:"required"`
		RoleID   uint32 `json:"roleId" validate:"required"`
		IsActive *bool  `json:"isActive"`
		IsAdmin  bool   `json:"isAdmin"`
	}

	UpdateStaffRequest struct {
		Fullname *string `json:"fullname"`
		Password *string `json:"password" validate:"omitempty,min=6"`
		RoleID   *uint32 `json:"roleId"`
		IsActive *bool   `json:"isActive"`
		IsAdmin  *bool   `json:"isAdmin"`
	}
)
