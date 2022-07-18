package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/capungkoneng/anterkenktu/api"
	db "github.com/capungkoneng/anterkenktu/db/sqlc"
	"github.com/capungkoneng/anterkenktu/util"
	_ "github.com/lib/pq"
)

// const (
// 	dbsource = "postgres"
// 	dbdriver = "postgresql://postgres:@localhost:5432/coba?sslmode=disable"
// )

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config", err)
	}
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")
	conn, err := sql.Open("postgres", fmt.Sprintf("host=%s user=%s password=%s port=%s dbname=%s sslmode=require TimeZone=Asia/Shanghai", dbHost, dbUser, dbPass, dbPort, dbName))
	// conn, err := sql.Open("postgres", "postgresql://postgres:@localhost:5432/coba?sslmode=disable")
	if err != nil {
		log.Fatal("cannot connect to database", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(config, store)

	err = server.Start()
	if err != nil {
		log.Fatal("cannot start server", err)
	}

}
