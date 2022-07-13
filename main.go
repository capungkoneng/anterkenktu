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

var (
	server *gin.Engine
	dbCon  *db.Queries

	// AuthController controllers.AuthController
	// AuthRoutes     routes.AuthRoutes
)

func init() {

	conn, err := sql.Open("postgres", fmt.Sprintf("host=%s user=%s password=%s port=%s dbname=%s sslmode=require"))
	if err != nil {
		log.Fatalf("could not connect to postgres database: %v", err)
	}

	dbCon = db.New(conn)

	fmt.Println("PostgreSQL connected successfully...", dbCon)

	// AuthController = *controllers.NewAuthController(db)
	// AuthRoutes = routes.NewAuthRoutes(AuthController)

	server = gin.Default()
}

func main() {
	// conn, err := sql.Open("postgres", fmt.Sprintf("host=%s user=%s password=%s port=%s dbname=%s sslmode=require"))
	// if err != nil {
	// 	log.Fatal("cannot connect to database", err)
	// }

	// repo := db.NewStore(conn)
	// router := gin.Default()
	// fmt.Println("koneksi:", conn)
	router := server.Group("/api")

	router.GET("/", api.GetHandlerHallo)

	// Hanlers := NewHandlers(repo)
	// fmt.Println("repo:", Hanlers)
	server.Run()

}
