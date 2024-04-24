package main

import (
	"database/sql"
	"log"

	"github.com/tranhuuhuy297/simple-bank-golang/api"
	db "github.com/tranhuuhuy297/simple-bank-golang/db/sqlc"
	"github.com/tranhuuhuy297/simple-bank-golang/util"

	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config: ", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal(err)
	}

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server")
	}
}
