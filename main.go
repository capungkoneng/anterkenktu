package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/capungkoneng/anterkenktu/controllers"
	dbConn "github.com/capungkoneng/anterkenktu/db/sqlc"
	"github.com/capungkoneng/anterkenktu/routes"
	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

var (
	server             *gin.Engine
	db                 *dbConn.Queries
	AuthController     controllers.AuthController
	KategoriController controllers.KategoriController

	KategoriRoutes routes.KategoriRoutes
	AuthRoutes     routes.AuthRoutes
)

func init() {
	// config, err := util.LoadConfigV(".")
	// if err != nil {
	// 	log.Fatal("cannot load config", err)
	// }
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")
	conn, err := sql.Open("postgres", fmt.Sprintf("host=%s user=%s password=%s port=%s dbname=%s sslmode=require TimeZone=Asia/Shanghai", dbHost, dbUser, dbPass, dbPort, dbName))
	// conn, err := sql.Open("postgres", config.DBSource)
	if err != nil {
		log.Fatalf("could not connect to postgres database: %v", err)
	}
	db = dbConn.New(conn)
	fmt.Println("PostgreSQL connected successfully...")

	AuthController = *controllers.NewAuthController(db)
	KategoriController = *controllers.NewKategoriController(db)

	AuthRoutes = routes.NewAuthRoutes(AuthController)
	KategoriRoutes = routes.NewKategoriRoutes(KategoriController)

	server = gin.Default()
}

func main() {
	server.Use(cors.Default())

	router := server.Group("/api")

	router.GET("/healthchecker", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"success": true, "massage": "Welcome "})
	})
	AuthRoutes.AuthRoute(router)
	KategoriRoutes.KategoriRoute(router)
	log.Fatal(server.Run())

	// conn, err := sql.Open("postgres", "postgresql://postgres:@localhost:5432/anter?sslmode=disable")
	// if err != nil {
	// 	log.Fatal("cannot connect to database", err)
	// }

	// store := db.NewStore(conn)
	// server := api.NewServer(config, store)

	// err = server.Start()
	// if err != nil {
	// 	log.Fatal("cannot start server", err)
	// }

}
