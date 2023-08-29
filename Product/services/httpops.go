package services

import (
	"encoding/json"
	"net/http"
)

func (pd *ProductDataMemory) FetchProductsHttp() ([]Product, error) {
	res, err := http.Get("http://localhost:3000/products")
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	products := make([]Product, 0, 10)
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&products)
	if err != nil {
		panic(err)
	}
	return products, nil
}
