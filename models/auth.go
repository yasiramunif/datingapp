package models

type RegisterRequest struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	ImgPath   string `json:"imgPath"`
	Fullname  string `json:"fullname"`
	Gender    string `json:"gender"`
	BirthDate string `json:"birthDate"`
}

type AuthResponse struct {
	User *User `json:"user"`
	Token string `json:"token"`
}

func NewAuthResponse(user *User, token string) *AuthResponse {
	return &AuthResponse{
		User: user,
    Token: token,
	}
}

type LoginRequest struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
}