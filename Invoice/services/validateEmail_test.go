package services_test

import (
	"fmt"
	"invoice/services"
	"testing"
)

func TestEmails(t *testing.T) {

	emails := []struct {
		in       string
		expected bool
	}{
		{"shubham@gmail.com", true}, {"rahul@gmai.com", false}, {"abhijit@gmail.com", true}, {"amit@gmail.com", false},
	}

	for _, input := range emails {
		t.Run(fmt.Sprintf("validating %v", input.in), func(ct *testing.T) {
			output, err := services.ValidateEmail(input.in)
			if err != nil {
				ct.Error(err)
				return
			}
			if output != true {
				ct.Errorf("Expected %v got %v", true, output)
			}
		})
	}
}

func BenchmarkEmailValidator(b *testing.B) {
	for i := 0; i < b.N; i++ {
		services.ValidateEmail("rahul@gmail.com")
	}
}
