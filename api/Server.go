package api

import (
	"github.com/gorilla/mux"
	"github.com/janicaleksander/GoInterviewTask1/handlers"
	"log"
	"net/http"
)

type Server struct {
	listenAddress string
}

func NewServer(port string) *Server {
	return &Server{
		listenAddress: port,
	}

}
func (s *Server) RunServer() {
	r := mux.NewRouter()
	r.HandleFunc("/api/joke", handlers.MakeHandler(handlers.HandleGetJoke))
	log.Println("Running on: ", s.listenAddress)
	http.ListenAndServe(s.listenAddress, r)
}
