package dto

type Session struct {
	ID           string   `json:"id" validate:"required"`
	Username     string   `json:"username" validate:"required"`
	StaffID      uint32   `json:"staffId" validate:"required"`
	RoleID       uint32   `json:"roleId"`
	RoleLabel    string   `json:"roleLabel"`
	IsAdmin      bool     `json:"isAdmin"`
	Permissions  []string `json:"permissions"`
	CreatedAt    int64    `json:"createdAt" validate:"required"`
	RefreshToken string   `json:"refreshToken,omitempty"`
}
