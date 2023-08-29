package services_test

import (
	"fmt"
	"invoice/services"
	"testing"
	"time"
)

func TestInvoice(t *testing.T) {

	Date := time.Now().Format("Mon Jan 2 2006")
	input := services.NewInvoice(7, "Rahul", "Pune")
	const expected = `
Date: %s

Invoice No: 7 is confirmed.
Following are the details

Name : Rahul,

Address: Pune
`
	text := fmt.Sprintf(expected, Date)

	output, err := services.GenerateInvoiceTmpl(input)

	if err != nil {
		t.Error(err)
		return
	}

	if text != output {
		t.Errorf("Expected %v got %v", text, output)
	}

}

func BenchmarkInvoice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		input := services.NewInvoice(7, "rahul", "Pune")
		services.GenerateInvoiceTmpl(input)
	}
}
