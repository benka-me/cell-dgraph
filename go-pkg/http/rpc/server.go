package rpc

import (
	"github.com/benka-me/PaasEnger/service-hive/go/go-proto"
	"github.com/benka-me/PaasEnger/services/db/pkg/conn"
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
	dgraph     conn.Dgraph
)

func StartCellDgraph() {
	var err error
	dgraph.Dgraph, err = conn.NewClient()
	if err != nil {
		panic(err)
	}

	grpcServer = grpc.NewServer()
	lis, err := net.Listen("tcp", port)

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	{
		proto.RegisterDgraphUserServer(grpcServer, guide)
		reflection.Register(grpcServer)
	}
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
