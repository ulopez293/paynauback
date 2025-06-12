package models

type CreateProductoRequest struct {
	Nombre      string  `json:"nombre"`
	Descripcion string  `json:"descripcion"`
	Precio      float64 `json:"precio"`
	Stock       int     `json:"stock"`
}

type UpdateProductoRequest struct {
	Nombre      string  `json:"nombre"`
	Descripcion string  `json:"descripcion"`
	Precio      float64 `json:"precio"`
	Stock       int     `json:"stock"`
}
