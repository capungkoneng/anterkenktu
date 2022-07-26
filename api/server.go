package api

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Server struct {
	// config util.Config
	router *gin.Engine
}

func NewServer() *Server {
	server := &Server{
		// config: config,
	}
	router := gin.Default()
	router.Use(cors.Default())

	router.GET("/mobil", server.GetListMobil)
	router.POST("/mobil/tambahmobil", server.CreateMobil)

	server.router = router
	return server
}

func (server *Server) Start() error {
	return server.router.Run()
}
