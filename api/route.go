package api

import (
	db "github.com/capungkoneng/anterkenktu/db/sqlc"
	"github.com/gin-gonic/gin"
)

type Server struct {
	store  db.Store
	router *gin.Engine
}

func NewServer(store *db.Store) *Server {
	server := &Server{store: *store}
	router := gin.Default()

	router.GET("/kategori/", server.GetListKategori)
	router.POST("/kategori", server.CreateKategori)

	server.router = router
	return server
}

func (server *Server) Start() error {
	return server.router.Run()
}
