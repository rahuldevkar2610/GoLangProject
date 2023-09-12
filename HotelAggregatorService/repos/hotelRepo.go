package repos

import (
	"HotelAggregatorService/models"
	"encoding/json"
	"net/http"
	"sync"
)

type HotelMemoryRepo struct {
	hotelsdb []models.Hotel
	lock     sync.RWMutex
}

func GetHotelRepo() (HotelMemoryRepo, error) {
	//silver data
	res, err := http.Get("http://localhost:8085/inventory")
	if err != nil {
		return HotelMemoryRepo{}, err
	}

	defer res.Body.Close()

	hotels1 := make([]models.Hotel, 0, 10)

	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&hotels1)
	if err != nil {
		panic(err)
	}

	//golden data
	res, err = http.Get("http://localhost:8086/inventory")
	if err != nil {
		return HotelMemoryRepo{}, err
	}
	defer res.Body.Close()

	hotels2 := make([]models.Hotel, 0, 10)

	decoder = json.NewDecoder(res.Body)
	err = decoder.Decode(&hotels2)
	if err != nil {
		panic(err)
	}

	hotels := make([]models.Hotel, 0, 10)

	for _, h1 := range hotels1 {
		hotels = append(hotels, h1)
	}

	for _, h2 := range hotels2 {
		hotels = append(hotels, h2)
	}

	return HotelMemoryRepo{hotelsdb: hotels}, nil

}

func (hmr *HotelMemoryRepo) GetHotels() ([]models.Hotel, error) {
	hmr.lock.RLock()
	defer hmr.lock.RUnlock()
	hotels := make([]models.Hotel, 0, 10)

	for _, h := range hmr.hotelsdb {
		hotels = append(hotels, h)
	}

	return hotels, nil
}
