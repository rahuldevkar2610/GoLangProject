package services

import (
	"ProductsUsingChannels/models"
	"ProductsUsingChannels/repository"
	"fmt"
)

func DiscountOnProductsChannel(productRepo repository.ProductRepo) (<-chan models.Product, error) {
	productsDataChannel, err := productRepo.ProductReaders()
	if err != nil {
		return nil, err
	}

	productMap := map[string]float64{
		"Electronic": 10,
		"Grocery":    12.5,
		"Clothing":   20.5,
		"Kitchen":    10.5,
	}

	discountedProductsChannel := make(chan models.Product)
	go func() {

		defer close(discountedProductsChannel)
		for product := range productsDataChannel {
			if discount, ok := productMap[product.Type]; ok {
				product.Price = product.Price - uint64(float64(product.Price)*discount/100)
			}
			discountedProductsChannel <- product
		}

	}()

	return discountedProductsChannel, nil
}

func BeforeDiscount(productRepo repository.ProductRepo, Type string) error {
	var choice int
	if Type == "Electronic" {
		choice = 1
	} else if Type == "Grocery" {
		choice = 2
	} else if Type == "Kitchen" {
		choice = 3
	} else if Type == "Clothing" {
		choice = 4
	}
	productsDataChannel, err := productRepo.FetchProductsHttp()
	generalProductChannel := make(chan models.Product)
	if err != nil {
		return err
	}
	switch choice {
	case 1:

		go func() {
			defer close(generalProductChannel)
			for product := range productsDataChannel {
				if product.Type == Type {
					generalProductChannel <- product
				}

			}
		}()

		for product := range generalProductChannel {
			fmt.Println("-----------------------------------------------Before Discount-------------------------------------------------")
			fmt.Println(product)
			AfterDiscount(product)
		}
		fmt.Println("----------------------------------------------------------------------------------------------------------------")

		return nil
	case 2:
		go func() {
			defer close(generalProductChannel)
			for product := range productsDataChannel {
				if product.Type == Type {
					generalProductChannel <- product
				}

			}
		}()

		for product := range generalProductChannel {
			fmt.Println("-----------------------------------------------Before Discount-------------------------------------------------")
			fmt.Println(product)
			AfterDiscount(product)
		}
		fmt.Println("----------------------------------------------------------------------------------------------------------------")

		return nil
	case 3:
		go func() {
			defer close(generalProductChannel)
			for product := range productsDataChannel {
				if product.Type == Type {
					generalProductChannel <- product
				}

			}
		}()

		for product := range generalProductChannel {
			fmt.Println("-----------------------------------------------Before Discount-------------------------------------------------")
			fmt.Println(product)
			AfterDiscount(product)
		}
		fmt.Println("----------------------------------------------------------------------------------------------------------------")
		return nil
	case 4:
		go func() {
			defer close(generalProductChannel)
			for product := range productsDataChannel {
				if product.Type == Type {
					generalProductChannel <- product
				}

			}
		}()

		for product := range generalProductChannel {
			fmt.Println("-----------------------------------------------Before Discount-------------------------------------------------")
			fmt.Println(product)
			AfterDiscount(product)
		}
		fmt.Println("----------------------------------------------------------------------------------------------------------------")
		return nil

	}

	return nil
}

func AfterDiscount(product models.Product) {
	productMap := map[string]float64{
		"Electronic": 10,
		"Grocery":    12.5,
		"Clothing":   20.5,
		"Kitchen":    10.5,
	}

	if discount, ok := productMap[product.Type]; ok {
		product.Price = product.Price - uint64(float64(product.Price)*discount/100)
	}

	fmt.Println("-----------------------------------------------After Discount --------------------------------------------------")
	fmt.Println(product)
	fmt.Println("----------------------------------------------------------------------------------------------------------------")
}
