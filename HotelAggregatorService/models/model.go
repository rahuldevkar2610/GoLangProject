package models

type Hotel struct {
	Source    string  `json:"source"`
	HotelName string  `json:"name"`
	RoomType  string  `json:"room-type"`
	Price     float64 `json:"price,string"`
}

func NewProduct(source, hotelname, roomtype string, price float64) Hotel {

	return Hotel{
		Source:    source,
		HotelName: hotelname,
		RoomType:  roomtype,
		Price:     price,
	}
}
