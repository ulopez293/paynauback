package models

type ProductoItem struct {
	ProductoID string `json:"producto_id"`
	Cantidad   int    `json:"cantidad"`
}

type CreateOrdenRequest struct {
	Productos []ProductoItem `json:"productos"`
}
