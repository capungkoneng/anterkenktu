package api

import (
	db "github.com/capungkoneng/anterkenktu/db/sqlc"
	"github.com/capungkoneng/anterkenktu/util"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Server struct {
	config util.Config
	store  db.Store
	router *gin.Engine
}

func NewServer(config util.Config, store db.Store) *Server {
	server := &Server{
		store:  store,
		config: config,
	}
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"https://localhost:3000"},
		AllowMethods:     []string{"POST", "PUT", "PATCH", "DELETE", "GET"},
		AllowHeaders:     []string{"Content-Type,access-control-allow-origin, access-control-allow-headers"},
		AllowCredentials: true,
	}))
	// router.Use(cors.New(cors.Config{
	// 	AllowOrigins:     []string{"https://localhost:3000"},
	// 	AllowMethods:     []string{"PUT", "PATCH", "DELETE", "OPTIONS", "POST", "GET"},
	// 	AllowHeaders:     []string{"Origin"},
	// 	ExposeHeaders:    []string{"Content-Length"},
	// 	AllowCredentials: true,
	// 	AllowOriginFunc: func(origin string) bool {
	// 		return origin == "https://localhost:3000"
	// 	},
	// 	MaxAge: 12 * time.Hour,
	// }))

	router.POST("/users", server.CreateUser)

	router.GET("/kategori/", server.GetListKategori)
	router.POST("/kategori", server.CreateKategori)

	server.router = router
	return server
}

func (server *Server) Start() error {
	return server.router.Run()
}
