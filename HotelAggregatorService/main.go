package main

import (
	"HotelAggregatorService/config"
	"HotelAggregatorService/handlers"
	"fmt"
	"net/http"

	"github.com/spf13/viper"
)

func main() {
	fmt.Println("Welcome to Hotel Assignment......")

	config.AppConfig()

	r := handlers.SetupRoutes()

	http.Handle("/", r)
	err := http.ListenAndServe(":"+viper.GetString("port"), nil)

	if err != nil {
		panic(err)
	}
}
