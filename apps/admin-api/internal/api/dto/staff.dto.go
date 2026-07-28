package dto

type (
	GetStaffByIdResponse struct {
		StaffId  uint32 `json:"id" validate:"required" example:"1"`
		Username string `json:"username" validate:"required" example:"username"`
		Fullname string `json:"fullname" validate:"required" example:"fullname"`
		RolesID  uint32 `json:"roleId" validate:"required" example:"1"`
		IsActive bool   `json:"isActive" validate:"required" example:"true"`
	}
)
