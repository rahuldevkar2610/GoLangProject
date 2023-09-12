package repos

import "HotelAggregatorService/models"

type HotelsRepo interface {
	// GetHotel(id string) (*models.Hotel, error)
	GetHotels() ([]models.Hotel, error)
	// GetHotelsByPriceAsc() ([]models.Hotel, error)
	// AddHotel(product models.Hotel) (string, error)
	// UpdateHotel(product models.Hotel) error
	// DeleteHotel(id string) error
}
