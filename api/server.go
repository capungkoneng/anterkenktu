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
	router.Use(cors.Default())

	router.POST("/users", server.CreateUser)

	router.GET("/mobil/", server.GetListMobil)
	router.POST("/mobil/tambahmobil", server.CreateMobil)

	router.GET("/kategori/", server.GetListKategori)
	router.POST("/kategori", server.CreateKategori)

	server.router = router
	return server
}

func (server *Server) Start() error {
	return server.router.Run()
}
