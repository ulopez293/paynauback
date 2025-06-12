package models

type CreateSucursalRequest struct {
	Nombre    string `json:"nombre"`
	Direccion string `json:"direccion"`
}

type UpdateSucursalRequest struct {
	Nombre    string `json:"nombre"`
	Direccion string `json:"direccion"`
}
