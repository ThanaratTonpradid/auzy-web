package dto

type (
	LoginRequest struct {
		Username string `json:"username" validate:"required" example:"username"`
		Password string `json:"password" validate:"required" example:"password"`
	}
	LoginResponse struct {
		Token        string `json:"token" example:"Y29weSB0byBjbGlwYm9hcmQuCg=="`
		RefreshToken string `json:"refreshToken" example:"cmVmcmVzaCB0b2tlbiBleGFtcGxl"`
		ExpiresIn    int64  `json:"expiresIn" example:"3600"`
	}
	RefreshTokenRequest struct {
		RefreshToken string `json:"refreshToken" validate:"required" example:"cmVmcmVzaCB0b2tlbiBleGFtcGxl"`
	}
	RefreshTokenResponse struct {
		Token        string `json:"token" example:"Y29weSB0byBjbGlwYm9hcmQuCg=="`
		RefreshToken string `json:"refreshToken" example:"cmVmcmVzaCB0b2tlbiBleGFtcGxl"`
		ExpiresIn    int64  `json:"expiresIn" example:"3600"`
	}
)
