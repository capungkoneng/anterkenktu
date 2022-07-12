package main

import (
	"database/sql"
	"log"

	"github.com/capungkoneng/anterkenktu/api"
	db "github.com/capungkoneng/anterkenktu/db/sqlc"
)

func main() {
	conn, err := sql.Open("postgres", "dbuser= dbpassword= dbhost= dbport= dbdatabse= sslmode=require")
	if err != nil {
		log.Fatal("cannot connect to database", err)
	}
	store := db.NewStore(conn)
	server := api.NewServer(&store)

	err = server.Start("")
	if err != nil {
		log.Fatal("cannot connect to server", err)
	}

	// server.SetTrustedProxies(nil)
	// listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
