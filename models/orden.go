package models

type ProductoItem struct {
	ProductoID string `json:"producto_id"`
	Cantidad   int    `json:"cantidad"`
}

type CreateOrdenRequest struct {
	Cliente   string         `json:"cliente"`
	Productos []ProductoItem `json:"productos"`
}
