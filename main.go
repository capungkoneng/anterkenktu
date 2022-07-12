package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/capungkoneng/anterkenktu/api"
	db "github.com/capungkoneng/anterkenktu/db/sqlc"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type Server struct {
	store *db.Repo
}

func NewHandlers(store *db.Repo) *Server {
	server := &Server{store: store}
	return server
}
func main() {
	conn, err := sql.Open("postgres", fmt.Sprintf("host=%s user=%s password=%s port=%s dbname=%s sslmode=require"))
	if err != nil {
		log.Fatal("cannot connect to database", err)
	}

	repo := db.NewStore(conn)
	router := gin.Default()
	fmt.Println("koneksi:", conn)

	router.GET("/", api.GetHandlerHallo)

	Hanlers := NewHandlers(repo)
	fmt.Println("repo:", Hanlers)
	router.Run()

}
