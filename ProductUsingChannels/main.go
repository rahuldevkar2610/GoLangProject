package main

import (
	"ProductsUsingChannels/repository"
	"ProductsUsingChannels/services"
	"fmt"
)

func main() {
	fmt.Println("Welcome to Product Assignment Using Channels.....")

	productData := repository.ProductDataMemory{}
	dataOfProductFromFile, err := productData.ProductReaders()
	productDataFromJsonServer, err := productData.FetchProductsHttp()

	if err != nil {
		panic(err)
	}
	for {

		fmt.Println("\nMenu List For Users:")
		fmt.Println("1. Read Product Data from product json file")
		fmt.Println("2. Read Product Data from json server")
		fmt.Println("3. Display Electronic products Before and After Discount")
		fmt.Println("4. Display Grocery products Before and After Discount")
		fmt.Println("5. Display Kitchen products Before and After Discount")
		fmt.Println("6. Display Clothing products Before and After Discount")
		fmt.Println("7. Display Product List After Discount")
		fmt.Println("8. Store Discounted Products in db.json File")
		fmt.Println("9. Exit")
		fmt.Print("Enter your choice: ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:

			fmt.Println("---------------------------Products Before Discount(Reading From the Product.json file) --------------------------------- ")
			for product := range dataOfProductFromFile {
				fmt.Println(product)
			}
			fmt.Println("--------------------------------------------------------------------------------------")
		case 2:
			fmt.Println("---------------------------Products Before Discount(Reading From the Json Server) --------------------------------- ")
			for product := range productDataFromJsonServer {
				fmt.Println(product)
			}
			fmt.Println("--------------------------------------------------------------------------------------")
		case 3:
			Type := "Electronic"
			err := services.BeforeDiscount(&productData, Type)
			if err != nil {
				panic(err)
			}
		case 4:
			Type := "Grocery"
			err := services.BeforeDiscount(&productData, Type)
			if err != nil {
				panic(err)
			}
		case 5:
			Type := "Kitchen"
			err := services.BeforeDiscount(&productData, Type)
			if err != nil {
				panic(err)
			}
		case 6:
			Type := "Clothing"
			err := services.BeforeDiscount(&productData, Type)
			if err != nil {
				panic(err)
			}
		case 7:
			discountedProductChannel, err := services.DiscountOnProductsChannel(&productData)

			if err != nil {
				panic(err)
			}
			for discountedProduct := range discountedProductChannel {
				fmt.Println(discountedProduct)
			}
		case 8:
			discountedChannel, err := services.DiscountOnProductsChannel(&productData)

			if err != nil {
				panic(err)
			}
			message := productData.ProductWriteToJsonFile(discountedChannel)
			fmt.Println(message)
		case 9:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice")
		}

	}
}
