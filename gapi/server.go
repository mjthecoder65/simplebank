package gapi

import (
	"log"

	db "github.com/mjthecoder65/simplebank/db/sqlc"
	"github.com/mjthecoder65/simplebank/pb"
	"github.com/mjthecoder65/simplebank/token"
	"github.com/mjthecoder65/simplebank/util"
)

// Server serves gRPC requests for our banking service
type Server struct {
	pb.UnimplementedSimpleBankServer
	store  *db.Store
	maker  token.Maker
	config util.Config
}

func NewServer(config util.Config, store *db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.PasetoSecretKey)

	if err != nil {
		log.Fatal("failed to get token maker", err)
		return nil, err
	}

	server := &Server{
		store:  store,
		maker:  tokenMaker,
		config: config,
	}

	return server, nil
}
