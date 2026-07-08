package model

type Product struct {
	Id    uint`json:"id_product"`
	Name  string`json:"name"`
	Price float64`json:"price"`
}
