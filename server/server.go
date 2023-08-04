package server

import (
	"log"

	"bitbucket.org/adson14/api_golang_ii.git/server/routes"
	"github.com/gin-gonic/gin"
)

type Server struct {
	port   string
	server *gin.Engine
}

func NewServer() Server {
	return Server{
		port:   "5000",
		server: gin.Default(),
	}
}

func (s *Server) Run() {
	router := routes.Routes(s.server)

	log.Print("Servidor iniciado em: " + s.port)
	log.Fatal(router.Run(":" + s.port))
}
