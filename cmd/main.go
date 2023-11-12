package main

import (
	"database/sql"
	"log"

	"github.com/Ridju/ticktr/config"
	db "github.com/Ridju/ticktr/internal/db/sqlc"
	"github.com/Ridju/ticktr/internal/token"
	"github.com/Ridju/ticktr/internal/user"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot open config file", err)
		return
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Cannot connect to DB:", err)
	}

	store := db.NewStore(conn)
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

	user.InitUserRouter(r.Group("/user"), store, config, tokenMaker)

	r.Run(config.ServerAddress)
}
