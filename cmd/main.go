package main

import (
	"log"

	"github.com/Ridju/ticktr/config"
	"github.com/Ridju/ticktr/internal/db"
	"github.com/Ridju/ticktr/internal/token"
	"github.com/Ridju/ticktr/internal/user"
	"github.com/gin-gonic/gin"
)

func main() {
	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot open config file", err)
		return
	}

	gormDB, err := db.Init(config.DBSource)
	if err != nil {
		log.Fatal("Could not connect to DB!", err)
		return
	}

	tokenMaker, err := token.NewJWTMaker(config.AccessTokenKey)
	if err != nil {
		log.Fatal("Error while creating tokenmaker", err)
		return
	}

	r := gin.Default()
	user.InitUserRouter(r.Group("/user"), gormDB, config, tokenMaker)
}
