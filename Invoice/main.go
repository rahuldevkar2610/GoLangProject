package main

import (
	"fmt"
	"invoice/services"
)

func main() {
	fmt.Println("Invoice Assignments....!!!!")

	invoice := services.NewInvoice(1500, "Rahul", "Pune")
	invoicetext := services.GenerateInvoice(invoice)

	fmt.Println(invoicetext)
	fmt.Println("============")

	invoicetext2, err := services.GenerateInvoiceTmpl(invoice)
	if err != nil {
		panic(err)
	}
	fmt.Println(invoicetext2)

	email := "rahul@gmail.com"
	isValidEmail, err := services.ValidateEmail(email)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s is valid or not : %v ", email, isValidEmail)
}
