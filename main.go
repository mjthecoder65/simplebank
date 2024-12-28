package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/mjthecoder65/simplebank/api"
	db "github.com/mjthecoder65/simplebank/db/sqlc"
	"github.com/mjthecoder65/simplebank/util"
)

func main() {
	config, err := util.LoadConfig(".")

	if err != nil {
		log.Fatal("failed to load config", err)
		return
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)

	if err != nil {
		log.Fatal("failed to connect to the database:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(config, store)

	err = server.Start(config.ServerAddress)

	if err != nil {
		log.Fatal("cannot start server")
	}
}
