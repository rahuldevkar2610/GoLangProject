package handlers

import (
	"HotelAggregatorService/config"
	"HotelAggregatorService/middlewares"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Home handler of API"))
}

func SetupRoutes() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/", HomeHandler)

	p := r.PathPrefix("/hotels").Subrouter()

	SetupHotelsRoutes(p)

	return r
}

// "/hotels"
func SetupHotelsRoutes(hRouter *mux.Router) {
	repo, err := config.LoadHotelsRepo()
	auth := config.LoadAuthService()
	fmt.Println("Inside the SetupHotelsRoutes...")

	if err != nil {
		fmt.Println(err)
		return
	}

	// 1) http://<server>:<port>/hotels
	// => above should give all hotels combined

	// gethotelshandler := GetHotelsHandler(repo)

	//using middleware
	gethotelshandler := middlewares.AuthMiddleware(auth, middlewares.LoggingMiddleware(GetHotelsHandler(repo)))
	hRouter.HandleFunc("", gethotelshandler).Methods(http.MethodGet)

	// 2) http://<server>:<port>/hotels?orderby=price
	// => above should give all hotels in ascending order of price
	// getHhtelsbypriceasc := GetHotelsByPriceAscHandler(repo)

	//using middleware
	gethotelsbypriceasc := middlewares.AuthMiddleware(auth, middlewares.LoggingMiddleware(GetHotelsByPriceAscHandler(repo)))
	hRouter.Queries("orderby", "").Handler(gethotelsbypriceasc)

	//----------------------------------------------------------------------------------------

	// 3) http://<server>:<port>/hotels?source=silver
	// => above should give all hotels from silver source. If source=gold then only gold source data. This
	// is to filter room by data source.

	// getHotelsbysource := GetHotelsBySource(repo)

	//using middleware
	getHotelsbysource := middlewares.AuthMiddleware(auth, middlewares.LoggingMiddleware(GetHotelsBySource(repo)))
	hRouter.Queries("source", "").Handler(getHotelsbysource)

	//----------------------------------------------------------------------------------------

	// 4) http://<server>:<port>/hotels?roomtype=LARGE
	// => above should give all hotels with room type LARGE. Also can give other options like
	// roomtype=X-LARGE or roomtype=SMALL. This is to filter rooms by type.

	// getHotelsbyroomtype := GetHotelsByRoomType(repo)

	// using middleware
	getHotelsbyroomtype := middlewares.AuthMiddleware(auth, middlewares.LoggingMiddleware(GetHotelsByRoomType(repo)))
	hRouter.Queries("roomtype", "").Handler(getHotelsbyroomtype)

	//----------------------------------------------------------------------------------------

	// 5) Queries and filters should work in combination like:
	// http://<server>:<port>/hotels?roomtype=LARGE&roomtype=X-LARGE&orderby=price
	// getHotelsbyroomtypes := GetHotelsByRoomTypes(repo)

	//using middleware
	gethotelsbyroomtypesandorderbyprice := middlewares.AuthMiddleware(auth, middlewares.LoggingMiddleware(GetHotelsByRoomTypesAndOrderByPrice(repo)))
	hRouter.Queries("roomtype", "", "orderby", "").Handler(gethotelsbyroomtypesandorderbyprice)

	//----------------------------------------------------------------------------------------

}
