package dto

type RefreshTokenShort struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type LoginResponse struct {
	Role         string `json:"role"`
	Name         string `json:"name"`
	UserID       string `json:"userId"`
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
