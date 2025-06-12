package models

type UpdateUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password,omitempty"` // opcional
}

type CreateUserRequest struct {
	Email      string `json:"email"`
	Password   string `json:"password"`
	RolUser    string `json:"rol_user"`
	SucursalID string `json:"sucursalId"`
}
