package models

type ErrorResponse struct {
	Error   string `json:"error"`
	Details string `json:"details,omitempty"` // se omite si está vacío
}
