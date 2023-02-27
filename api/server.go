package server

import (
	db "github.com/CRAZYKAYZY/web/db/sqlc"
	"github.com/gin-gonic/gin"
)

type Server struct {
	store  *db.Store
	router *gin.Engine
}

func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
