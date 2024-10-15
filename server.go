package main

import (
	"github.com/gorilla/mux"
	"log"
	"log/slog"
	"net/http"
)

type Server interface {
	Start() error
	routes()
}

type MuxServer struct {
	gorilla *mux.Router
	Client
}

func NewServer(db Client) Server {

	server := &MuxServer{
		mux.NewRouter(),
		db,
	}

	server.routes()
	return server
}

func (s *MuxServer) Start() error {
	slog.Info("serving at port 8080")
	log.Fatal(http.ListenAndServe(":8080", s.gorilla))
	return nil
}
