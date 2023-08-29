package main

import (
	"Product/services"
	"fmt"
)

func main() {
	fmt.Println("Product Assignment ......")

	productData := services.ProductDataMemory{}

	prodctsAfterDiscount, err := services.DiscountOnProductsType(&productData)
	if err != nil {
		panic(err)
	}

	fmt.Println(prodctsAfterDiscount)
}
