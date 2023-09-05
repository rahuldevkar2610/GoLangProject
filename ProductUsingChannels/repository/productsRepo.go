package repository

import (
	"ProductsUsingChannels/models"
)

type ProductRepo interface {
	FetchProductsHttp() (<-chan models.Product, error)
	ProductReaders() (<-chan models.Product, error)
	ProductWriteToJsonFile(<-chan models.Product) string
}

type ProductDataMemory struct{}
