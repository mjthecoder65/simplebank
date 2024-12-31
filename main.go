package main

import (
	"database/sql"
	"log"
	"net"

	_ "github.com/lib/pq"
	"github.com/mjthecoder65/simplebank/api"
	db "github.com/mjthecoder65/simplebank/db/sqlc"
	"github.com/mjthecoder65/simplebank/gapi"
	"github.com/mjthecoder65/simplebank/pb"
	"github.com/mjthecoder65/simplebank/util"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	config, err := util.LoadConfig(".")

	if err != nil {
		log.Fatal("failed to load config", err)
		return
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)

	if err != nil {
		log.Fatal("failed to connect to the database:", err)
	}

	store := db.NewStore(conn)
	runGRPCServer(config, store)
}

func runGRPCServer(config util.Config, store *db.Store) {
	server, err := gapi.NewServer(config, store)

	if err != nil {
		log.Fatal("failed to create new server")
	}

	grpcServer := grpc.NewServer()
	pb.RegisterSimpleBankServer(grpcServer, server)
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", config.GRPCServerAddress)
	if err != nil {
		log.Fatal("cannot create listener", err)
	}

	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("cannot start gRPC server")
	}
}

func runGinHTTPServer(config util.Config, store *db.Store) {
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("failed to create new server instance", err)
	}
	err = server.Start(config.ServerAddress)

	if err != nil {
		log.Fatal("cannot start server")
	}
}
