package models

type AuthUser struct {
	IDUser  string `json:"id_user"`
	Email   string `json:"email"`
	RolUser string `json:"rol_user"`
}

type AuthResponse struct {
	Token string `json:"token"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
