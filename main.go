package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	db "github.com/capungkoneng/anterkenktu/db/sqlc"
	_ "github.com/lib/pq"

	// "github.com/capungkoneng/anterkenktu/api"
	// db "github.com/capungkoneng/anterkenktu/db/sqlc"
	"github.com/gin-gonic/gin"
)

type Hanlers struct {
	Repo *db.Repo
}

func NewHandlers(repo *db.Repo) *Hanlers {
	return &Hanlers{Repo: repo}
}
func main() {
	conn, err := sql.Open("postgres", fmt.Sprintf("host=%s user=%s password=%s port=%s dbname=%s sslmode=require"))
	if err != nil {
		log.Fatal("cannot connect to database", err)
	}

	repo := db.NewStore(conn)
	router := gin.Default()
	fmt.Println(conn)
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	Hanlers := NewHandlers(repo)
	fmt.Println(Hanlers)
	router.Run() //
	// store := db.NewStore(conn)
	// server := api.NewServer(&store)

	// err = server.Start("")
	// if err != nil {
	// 	log.Fatal("cannot connect to server", err)
	// }

	// server.SetTrustedProxies(nil)
	// listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
