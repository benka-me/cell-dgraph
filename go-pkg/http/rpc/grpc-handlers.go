package rpc

import (
	"context"
	"fmt"
	"github.com/benka-me/PaasEnger/service-hive/go/go-proto"
	conn2 "github.com/benka-me/PaasEnger/services/db/pkg/conn"
	customer2 "github.com/benka-me/PaasEnger/services/db/pkg/customer"
	user2 "github.com/benka-me/PaasEnger/services/db/pkg/user"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) InitDb(ctx context.Context, in *proto.Empty) (*proto.Empty, error) {
	fmt.Println("============ Init DB ============")
	err := conn2.InitDgraph(dgraph.Dgraph, true)
	if err != nil {
		return in, status.Error(codes.Internal, err.Error())
	}
	err = conn2.InitDgraph(dgraph.Dgraph, false)
	if err != nil {
		return in, status.Error(codes.Internal, err.Error())
	}
	return in, status.Error(codes.OK, "successfully init")
}

func (s *Server) RegisterCustomer(ctx context.Context, handshake *proto.HandshakeReq) (*proto.Welcome, error) {
	wel, err := customer2.Handshake(handshake, dgraph)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return wel, status.Error(codes.OK, "")
}

func (s *Server) RegisterUser(ctx context.Context, newU *proto.NewUser) (*proto.User, error) {
	u, err := user2.NewUser(*newU).NewUser(dgraph)
	ret := proto.User(u)
	return &ret, err
}

func (s *Server) Login(ctx context.Context, req *proto.LoginRequest) (*proto.User, error) {
	u, err := user2.LoginRequest(*req).Login(dgraph)
	ret := proto.User(u)
	return &ret, err
}

func (s *Server) UidGet(ctx context.Context, str *proto.Str) (*proto.User, error) {
	get, err := user2.Uid(str.Ing).Get(dgraph)
	if err != nil {
		return nil, err
	}
	ret := proto.User(get[0])
	return &ret, nil
}
func (s *Server) UidDelete(ctx context.Context, str *proto.Str) (*proto.N, error) {
	got, err := user2.Uid(str.Ing).Delete(dgraph)
	if err != nil {
		return nil, err
	}
	return &proto.N{Deleted: got}, nil
}

func (s *Server) UsernameGet(ctx context.Context, str *proto.Str) (*proto.User, error) {
	get, err := user2.Username(str.Ing).Get(dgraph)
	if err != nil {
		return nil, err
	}
	ret := proto.User(get[0])
	return &ret, nil
}
func (s *Server) UsernameDelete(ctx context.Context, str *proto.Str) (*proto.N, error) {
	got, err := user2.Username(str.Ing).Delete(dgraph)
	if err != nil {
		return nil, err
	}
	return &proto.N{Deleted: got}, nil
}

func (s *Server) UserSet(ctx context.Context, req *proto.User) (*proto.Str, error) {
	got, err := user2.User(*req).Set(dgraph)
	if err != nil {
		return nil, err
	}
	ret := &proto.Str{Ing: string(got)}
	return ret, nil
}
func (s *Server) UserDelete(ctx context.Context, req *proto.User) (*proto.N, error) {
	got, err := user2.User(*req).Delete(dgraph)
	if err != nil {
		return nil, err
	}
	return &proto.N{Deleted: got}, nil
}
