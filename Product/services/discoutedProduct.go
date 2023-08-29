package services

import "fmt"

func DiscountOnProductsType(productrepo ProductRepo) ([]Product, error) {
	product, err := productrepo.FetchProductsHttp()

	productMap := map[string]float64{
		"Electronic": 10,
		"Grocery":    12.5,
		"Clothing":   20.5,
		"Kitchen":    10.5,
	}

	products := make([]Product, 0, 10)

	fmt.Println("---------------------------------------------")
	fmt.Println("Products Rate Before Discount - ", product)
	fmt.Println("---------------------------------------------")

	for i := range product {
		if discount, ok := productMap[product[i].Type]; ok {
			product[i].Price = product[i].Price - uint64(float64(product[i].Price)*discount/100)
		}
		products = append(products, product[i])
	}
	fmt.Println("---------------------------------------------")
	fmt.Println("Products At Discounted Rate - ")
	fmt.Println("---------------------------------------------")

	return products, err
}
