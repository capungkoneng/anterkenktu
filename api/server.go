package api

import (
	"fmt"

	db "github.com/capungkoneng/anterkenktu/db/sqlc"
	"github.com/gin-gonic/gin"
)

type Server struct {
	store  db.Store
	router *gin.Engine
}

func NewServer(store *db.Store) *Server {
	server := &Server{store: *store}
	fmt.Println(server)
	router := gin.Default()

	router.GET("/", server.GetListKategori)

	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
