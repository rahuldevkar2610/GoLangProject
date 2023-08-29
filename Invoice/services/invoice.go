package services

import (
	"bytes"
	"errors"
	"fmt"
	"html/template"
	"time"
)

type Invoice struct {
	OrderId int32
	Name    string
	Address string
	Date    string
}

func NewInvoice(oid int32, name string, add string) Invoice {
	return Invoice{
		OrderId: oid,
		Name:    name,
		Address: add,
		Date:    time.Now().Format("Mon Jan 2 2006"),
	}
}

const templateformat = `
Date: %s

Invoice No: %d is confirmed.
Following are the details

Name : %s,

Address: %s
`

func GenerateInvoice(invoice Invoice) string {
	text := fmt.Sprintf(templateformat, invoice.Date, invoice.OrderId, invoice.Name, invoice.Address)
	return text
}

const textTemplate = `
Date: {{.Date}}

Invoice No: {{.OrderId}} is confirmed.
Following are the details

Name : {{.Name}},

Address: {{.Address}}
`

func GenerateInvoiceTmpl(invoice Invoice) (string, error) {

	if invoice.OrderId <= 0 {
		return "", errors.New("Invalid Order Id")
	}
	if invoice.Name == "" {
		return "", errors.New("name is required")
	}

	if invoice.Address == "" {
		return "", errors.New("address is required")
	}

	if invoice.Date == "" {
		return "", errors.New("date is required")
	}

	tmpl, err := template.New("invoice").Parse(textTemplate)
	if err != nil {
		panic(err)
	}

	var data bytes.Buffer
	err = tmpl.Execute(&data, invoice)
	if err != nil {
		panic(err)
	}
	return data.String(), nil
}
