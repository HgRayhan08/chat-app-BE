package dto

type AuthRequest struct {
	Email       string `json:"email"`
	Password    string `json:"password"`
	PhoneNumber string `json:"phone_number"`
	Username    string `json:"username"`
}

type TokenRequest struct {
	Token string `json:"token"`
}

type User struct {
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Username    string `json:"username"`
}

type RefreshToken struct {
	UserId string `json:"user_id"`
	Token  string `json:"token"`
}

type AuthResponse struct {
	Code         int    `json:"code"`
	Message      string `json:"message"`
	User         User   `json:"user,omitempty"`
	TokenRefresh string `json:"token_refresh,omitempty"`
	TokenJwt     string `json:"token_jwt,omitempty"`
}

type RefreshTokenResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Token   string `json:"token"`
}
