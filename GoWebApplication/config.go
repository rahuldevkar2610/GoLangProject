package main

import (
	"flag"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

const SERVER = "server"
const PORT = "port"

func AppConfig() {
	flag.String(SERVER, "appserver", "name of server")
	flag.String(PORT, "8080", "specify port on which server runs")

	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)

	viper.AutomaticEnv()
}
