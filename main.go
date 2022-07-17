package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/capungkoneng/anterkenktu/api"
	db "github.com/capungkoneng/anterkenktu/db/sqlc"
	_ "github.com/lib/pq"
)

// const (
// 	dbsource = "postgres"
// 	dbdriver = "postgresql://postgres:@localhost:5432/coba?sslmode=disable"
// )

func main() {
	conn, err := sql.Open("postgres", fmt.Sprintf("host=%s user=%s password=%s port=%s dbname=%s sslmode=require"))
	// conn, err := sql.Open("postgres", "postgresql://postgres:@localhost:5432/coba?sslmode=disable")
	if err != nil {
		log.Fatal("cannot connect to database", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(&store)

	err = server.Start()
	if err != nil {
		log.Fatal("cannot start server", err)
	}

}
