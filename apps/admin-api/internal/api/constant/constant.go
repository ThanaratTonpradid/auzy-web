package constant

import "time"

const TTLJWTExpires = 1 * time.Hour               // Access token expires in 1 hour
const TTLRefreshTokenExpires = 7 * 24 * time.Hour // Refresh token expires in 7 days
const TTLGeoIPCache = 24 * time.Hour
const KeySession = "session"

const (
	ErrCodeSomethingWentWrong = "SOMETHING_WENT_WRONG"
	ErrCodeBadRequest         = "BAD_REQUEST"
	ErrCodeUnauthorized       = "UNAUTHORIZED"
	ErrCodeForbidden          = "FORBIDDEN"
	ErrCodeNotFound           = "NOT_FOUND"
	ErrCodeInternalError      = "INTERNAL_ERROR"
)

const (
	CodeLoginFailed         = "Login failed"
	CodeCreateTokenFailed   = "CreateTokenFailed"
	CodeCreateSessionFailed = "CreateSessionFailed"
	CodeLogoutSuccess       = "LogoutSuccess"
)

const (
	RolesCreate = "ROLES_CREATE"
	RolesRead   = "ROLES_READ"
	RolesUpdate = "ROLES_UPDATE"
	RolesDelete = "ROLES_DELETE"

	StaffsCreate = "STAFFS_CREATE"
	StaffsRead   = "STAFFS_READ"
	StaffsUpdate = "STAFFS_UPDATE"
	StaffsDelete = "STAFFS_DELETE"

	VisitorsRead = "VISITORS_READ"
)

const (
	RoleAdmin  = "ADMIN"
	RoleStaff  = "STAFF"
	RoleMember = "MEMBER"
)
