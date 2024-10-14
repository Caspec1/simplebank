package gapi

import (
	"fmt"

	db "github.com/Caspec1/simplebank/db/sqlc"
	"github.com/Caspec1/simplebank/pb"
	"github.com/Caspec1/simplebank/token"
	"github.com/Caspec1/simplebank/util"
	"github.com/Caspec1/simplebank/worker"
)

type Server struct {
	pb.UnimplementedSimpleBankServer
	config          util.Config
	store           db.Store
	tokenMaker      token.Maker
	taskDistributor worker.TaskDistributor
}

// NewServer creates a new gRPC server
func NewServer(config util.Config, store db.Store, taskDistributor worker.TaskDistributor) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		config:          config,
		store:           store,
		tokenMaker:      tokenMaker,
		taskDistributor: taskDistributor,
	}
	return server, nil
}
