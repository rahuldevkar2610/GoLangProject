package services_test

import (
	"Product/services"
	"testing"
)

type MockProductRepo struct{}

func (mpr MockProductRepo) FetchProductsHttp() ([]services.Product, error) {
	products := []services.Product{
		{Name: "Product 1", Type: "Electronic", Price: 100, Seller: "Seller 1"},
		{Name: "Product 2", Type: "Grocery", Price: 200, Seller: "Seller 2"},
		{Name: "Product 3", Type: "Clothing", Price: 300, Seller: "Seller 3"},
	}
	return products, nil
}

func TestDiscountOnProductsType(t *testing.T) {
	mockProductRepo := MockProductRepo{}
	discountedProducts, err := services.DiscountOnProductsType(mockProductRepo)

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	expectedDiscountedPrices := map[string]uint64{
		"Electronic": 90,
		"Grocery":    175,
		"Clothing":   239,
	}

	for _, product := range discountedProducts {
		expectedPrice := expectedDiscountedPrices[product.Type]
		if product.Price != expectedPrice {
			t.Errorf("Mismatched discount for product %s. Expected: %d, Actual: %d", product.Type, expectedPrice, product.Price)
		}
	}
}
