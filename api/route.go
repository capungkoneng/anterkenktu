package api

import (
	"time"

	db "github.com/capungkoneng/anterkenktu/db/sqlc"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Server struct {
	store  db.Store
	router *gin.Engine
}

func NewServer(store *db.Store) *Server {
	server := &Server{store: *store}
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"https://anterkenktu07.herokuapp.com"},
		AllowMethods:     []string{"PUT", "PATCH", "DELETE", "OPTIONS", "POST", "GET"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "https://anterkenktu07.herokuapp.com"
		},
		MaxAge: 12 * time.Hour,
	}))
	router.GET("/kategori/", server.GetListKategori)
	router.POST("/kategori", server.CreateKategori)

	server.router = router
	return server
}

func (server *Server) Start() error {
	return server.router.Run()
}
