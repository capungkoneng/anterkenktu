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

func main() {
	conn, err := sql.Open("postgres", fmt.Sprintf("host=%s user=%s password=%s port=%s dbname=%s sslmode=require"))
	if err != nil {
		log.Fatalf("could not connect to postgres database: %v", err)
	}

	dbCon = db.New(conn)

	fmt.Println("PostgreSQL connected successfully...", dbCon, conn)
	server = gin.Default()

	// repo := db.NewStore(conn)
	// router := gin.Default()
	// fmt.Println("koneksi:", conn)
	v1 := server.Group("/api")

	v1.POST("/kategori", api.CreateKategori)

	// Hanlers := NewHandlers(repo)
	// fmt.Println("repo:", Hanlers)
	server.Run()

}
