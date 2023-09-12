package config

import (
	"HotelAggregatorService/auth"
	"HotelAggregatorService/repos"
	"flag"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

const SERVER = "server"
const PORT = "port"

func AppConfig() {
	flag.String(SERVER, "appserver", "name of server")
	flag.String(PORT, "8090", "specify port on which server runs")
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)

	viper.AutomaticEnv()
}

func LoadHotelsRepo() (repos.HotelsRepo, error) {

	if viper.Get("env") == "prod" {
		//if prod is set we will return actual service not the inmemory dummy one
		//just replace below with that
		hotelrepo, err := repos.GetHotelRepo()
		if err != nil {
			return nil, err
		}
		return &hotelrepo, nil
	} else {
		hotelrepo, err := repos.GetHotelRepo()
		if err != nil {
			return nil, err
		}
		return &hotelrepo, nil
	}
}

func LoadAuthService() auth.Auth {
	if viper.Get("env") == "prod" {
		//if prod is set we will return actual service not the inmemory dummy one
		//just replace below with that
		authsrv := auth.NewAuthServiceMemory()
		return &authsrv
	} else {
		authsrv := auth.NewAuthServiceMemory()
		return &authsrv
	}
}
