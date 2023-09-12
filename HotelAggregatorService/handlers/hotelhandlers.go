package handlers

import (
	"HotelAggregatorService/models"
	"HotelAggregatorService/repos"
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
)

func GetHotelsHandler(hrepo repos.HotelsRepo) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("inside get all hotels handler")
		fmt.Println(r.URL.Query().Get("orderby"))

		hotels, err := hrepo.GetHotels()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Add("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(hotels)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func GetHotelsByPriceAscHandler(hrepo repos.HotelsRepo) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("inside price asc handler")
		hotels, err := hrepo.GetHotels()

		sort.SliceStable(hotels, func(i, j int) bool {
			return hotels[i].Price < hotels[j].Price
		})

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Add("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(hotels)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func GetHotelsBySource(hrepo repos.HotelsRepo) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("inside source handler")
		hotels, err := hrepo.GetHotels()

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		filterhotel := make([]models.Hotel, 0, 10)

		filtervalue := r.URL.Query().Get("source")

		for _, hotel := range hotels {
			if hotel.Source == filtervalue {
				filterhotel = append(filterhotel, hotel)
			}
		}

		w.Header().Add("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(filterhotel)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func GetHotelsByRoomType(hrepo repos.HotelsRepo) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("inside room type handler")
		hotels, err := hrepo.GetHotels()

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		filterhotel := make([]models.Hotel, 0, 10)

		roomtypesparam := r.URL.Query()
		roomtypes := roomtypesparam["roomtype"]
		fmt.Println(roomtypes)

		//multiple room type
		if len(roomtypes) != 1 {
			fmt.Println("multiple room type")
			if val, ok := roomtypesparam["roomtype"]; ok {
				for j := range val {
					//reference semantics: avoid copy
					for i := range hotels {
						if hotels[i].RoomType == val[j] {
							fmt.Println(val[j])
							filterhotel = append(filterhotel, hotels[i])
						}
					}

				}
			}
		} else {
			//if only one room type parameter
			filtervalue := r.URL.Query().Get("roomtype")
			for i := range hotels {
				if hotels[i].RoomType == filtervalue {
					fmt.Println(filtervalue)
					filterhotel = append(filterhotel, hotels[i])
				}
			}
		}

		w.Header().Add("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(filterhotel)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func GetHotelsByRoomTypesAndOrderByPrice(hrepo repos.HotelsRepo) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("inside room type  and orderby parametr handler")
		hotels, err := hrepo.GetHotels()

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		filterhotel := make([]models.Hotel, 0, 10)

		queryparams := r.URL.Query()

		if roomtype, ok := queryparams["roomtype"]; ok {
			for j := range roomtype {
				//reference semantics: avoid copy
				for i := range hotels {
					if hotels[i].RoomType == roomtype[j] {
						fmt.Println(roomtype[j])
						filterhotel = append(filterhotel, hotels[i])
					}
				}

			}
		}

		if _, ok := queryparams["orderby"]; ok {
			sort.SliceStable(filterhotel, func(i, j int) bool {
				return filterhotel[i].Price < filterhotel[j].Price
			})
		}

		w.Header().Add("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(filterhotel)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
