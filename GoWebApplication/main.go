package main

import (
	"fmt"
	"net/http"

	"github.com/spf13/viper"
)

func Endpoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "From Endpoint")
}

func main() {
	AppConfig()
	app := viper.Get(SERVER)
	port := viper.Get(PORT)

	//setup

	http.HandleFunc("/welcome", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello From Server")
	})

	http.HandleFunc("/endpoint", Endpoint)
	fmt.Printf("server %v starting on port %v", app, port)

	//start the server

	http.ListenAndServe(":"+viper.GetString(PORT), nil)
}
