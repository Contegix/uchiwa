package main

import (
	"flag"

	"github.com/palourde/auth"
	"github.com/palourde/logger"
	"github.com/sensu/uchiwa/uchiwa"
)

func main() {
	configFile := flag.String("c", "./config.json", "Full or relative path to the configuration file")
	publicPath := flag.String("p", "public", "Full or relative path to the public directory")
	flag.Parse()

	config, err := uchiwa.LoadConfig(*configFile)
	if err != nil {
		logger.Fatal(err)
	}

	uchiwa.New(config)
	go uchiwa.Fetch(config.Uchiwa.Refresh, func() {})

	authentication := auth.New()

	if config.Uchiwa.Auth == "" {
		authentication.None()
	} else {
		authentication.Simple(config.Uchiwa.User, config.Uchiwa.Pass)
	}

	uchiwa.WebServer(config, publicPath, authentication)
}
