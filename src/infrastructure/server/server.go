package server

import (
	"database/sql"

	"github.com/Nagoya-Caravan-Hackathon-PDD/Github-api/src/infrastructure/http/router"
	"github.com/diegosz/go-graphql-client"
)

type httpServer struct {
	db            *sql.DB
	graphqlClient *graphql.Client
}

func NewHTTPserver(db *sql.DB, graphqlClient *graphql.Client) *httpServer {
	return &httpServer{db, graphqlClient}
}

func (s *httpServer) Run() {
	serv := router.NewRouter(s.db, s.graphqlClient)
	runWithGracefulShutdown(serv)
}
