package services

type ProductRepo interface {
	FetchProductsHttp() ([]Product, error)
}

type ProductDataMemory struct{}

type Product struct {
	Name   string `json:"name"`
	Type   string `json:"producttype"`
	Price  uint64 `json:"price"`
	Seller string `json:"seller"`
}

func NewProduct(name string, producttype string, price uint64, seller string) Product {
	return Product{
		Name:   name,
		Type:   producttype,
		Price:  price,
		Seller: seller,
	}
}
