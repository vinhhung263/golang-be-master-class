package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/vinhhung263/simplebank/api"
	db "github.com/vinhhung263/simplebank/db/sqlc"
	"github.com/vinhhung263/simplebank/util"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config: ", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	store := db.NewStore(conn)
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create server")
	}

	err = server.Start(config.ServerAddr)
	if err != nil {
		log.Fatal("can not start server:", err)
	}
}
