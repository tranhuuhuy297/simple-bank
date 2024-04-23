package main

import (
	"database/sql"
	"log"

	"github.com/tranhuuhuy297/simple-bank-golang/api"
	db "github.com/tranhuuhuy297/simple-bank-golang/db/sqlc"

	_ "github.com/lib/pq"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgresql://admin:admin@localhost:5432/postgres?sslmode=disable"
	serverAddress = "0.0.0.0:3001"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("cannot start server")
	}
}
