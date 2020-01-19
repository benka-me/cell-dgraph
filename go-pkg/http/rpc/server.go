package rpc

import (
	"github.com/benka-me/cell-dgraph/go-pkg/conn"
	"github.com/benka-me/cell-dgraph/go-pkg/dgraph"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

type (
	Server struct{}
)

const (
	port = ":5100"
)

var (
	grpcServer *grpc.Server
	guide      *Server
	db         conn.Dgraph
)

func StartCellDgraph() {
	var err error
	db.Dgraph, err = conn.NewClient()
	if err != nil {
		panic(err)
	}

	grpcServer = grpc.NewServer()
	lis, err := net.Listen("tcp", port)

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	{
		dgraph.RegisterDgraphUserServer(grpcServer, guide)
		reflection.Register(grpcServer)
	}
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
