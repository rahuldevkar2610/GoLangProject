package repository

import (
	"ProductsUsingChannels/models"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func (pd *ProductDataMemory) ProductReaders() (<-chan models.Product, error) {

	file1, err := ioutil.ReadFile("products.json")
	if err != nil {
		return nil, err
	}
	productChannel := make(chan models.Product)

	products := make([]models.Product, 0, 10)
	_ = json.Unmarshal([]byte(file1), &products)

	go func() {
		defer close(productChannel)
		for _, prod := range products {
			productChannel <- prod
			time.Sleep(time.Millisecond * 100)
		}
	}()
	return productChannel, nil
}

func (pd *ProductDataMemory) FetchProductsHttp() (<-chan models.Product, error) {
	productOutChannel := make(chan models.Product)
	res, err := http.Get("http://localhost:3000/products")
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	products := make([]models.Product, 0, 10)
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&products)
	if err != nil {
		panic(err)
	}

	go func() {
		defer close(productOutChannel)
		for _, prod := range products {
			productOutChannel <- prod
			time.Sleep(time.Millisecond * 100)
		}
	}()

	return productOutChannel, nil
}

func (pd *ProductDataMemory) ProductWriteToJsonFile(data <-chan models.Product) string {

	file, _ := os.OpenFile("db.json", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0755)
	defer file.Close()
	copydp := []models.Product{}
	for prod := range data {
		copydp = append(copydp, prod)
	}
	byteData, _ := json.MarshalIndent(copydp, "", "")
	io.WriteString(file, string(byteData))

	return " Successfully Added Discounted Products to JSON File "
}
