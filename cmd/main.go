package main

import (
	"log"

	"github.com/Ridju/ticktr/config"
	"github.com/Ridju/ticktr/internal/db"
)

func main() {
	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot open config file", err)
		return
	}
	_, err = db.Init(config.DBSource)
	if err != nil {
		log.Fatal("Could not connect to DB!", err)
		return
	}

}
