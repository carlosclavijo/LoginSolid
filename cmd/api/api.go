package api

import (
	"database/sql"
	"log"
	"net/http"
)

type ApiServer struct {
	addr string
	db   *sql.DB
}

func NewApiServer(addr string, db *sql.DB) *ApiServer {
	return &ApiServer{
		addr: addr,
		db:   db,
	}
}

func (s *ApiServer) Run() error {
	server := &http.Server{
		Addr:    s.addr,
		Handler: router(),
	}
	log.Printf("Server has started %s", s.addr)
	return server.ListenAndServe()
}
