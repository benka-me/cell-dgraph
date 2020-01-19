package rpc

import (
	"context"
	"fmt"
	"github.com/benka-me/cell-dgraph/go-pkg/conn"
	"github.com/benka-me/cell-dgraph/go-pkg/customer"
	"github.com/benka-me/cell-dgraph/go-pkg/user"
	"github.com/benka-me/hive/go-pkg/hive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) InitDb(ctx context.Context, in *hive.Empty) (*hive.Empty, error) {
	fmt.Println("============ Init DB ============")
	err := conn.InitDgraph(db.Dgraph, true)
	if err != nil {
		return in, status.Error(codes.Internal, err.Error())
	}
	err = conn.InitDgraph(db.Dgraph, false)
	if err != nil {
		return in, status.Error(codes.Internal, err.Error())
	}
	return in, status.Error(codes.OK, "successfully init")
}

func (s *Server) RegisterCustomer(ctx context.Context, handshake *hive.HandshakeReq) (*hive.Welcome, error) {
	wel, err := customer.Handshake(handshake, db)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return wel, status.Error(codes.OK, "")
}

func (s *Server) RegisterUser(ctx context.Context, newU *hive.NewUser) (*hive.User, error) {
	u, err := user.NewUser(*newU).NewUser(db)
	ret := hive.User(u)
	return &ret, err
}

func (s *Server) Login(ctx context.Context, req *hive.LoginRequest) (*hive.User, error) {
	u, err := user.LoginRequest(*req).Login(db)
	ret := hive.User(u)
	return &ret, err
}

func (s *Server) UidGet(ctx context.Context, str *hive.Str) (*hive.User, error) {
	get, err := user.Uid(str.Ing).Get(db)
	if err != nil {
		return nil, err
	}
	ret := hive.User(get[0])
	return &ret, nil
}
func (s *Server) UidDelete(ctx context.Context, str *hive.Str) (*hive.N, error) {
	got, err := user.Uid(str.Ing).Delete(db)
	if err != nil {
		return nil, err
	}
	return &hive.N{Deleted: got}, nil
}

func (s *Server) UsernameGet(ctx context.Context, str *hive.Str) (*hive.User, error) {
	get, err := user.Username(str.Ing).Get(db)
	if err != nil {
		return nil, err
	}
	ret := hive.User(get[0])
	return &ret, nil
}
func (s *Server) UsernameDelete(ctx context.Context, str *hive.Str) (*hive.N, error) {
	got, err := user.Username(str.Ing).Delete(db)
	if err != nil {
		return nil, err
	}
	return &hive.N{Deleted: got}, nil
}

func (s *Server) UserSet(ctx context.Context, req *hive.User) (*hive.Str, error) {
	got, err := user.User(*req).Set(db)
	if err != nil {
		return nil, err
	}
	ret := &hive.Str{Ing: string(got)}
	return ret, nil
}
func (s *Server) UserDelete(ctx context.Context, req *hive.User) (*hive.N, error) {
	got, err := user.User(*req).Delete(db)
	if err != nil {
		return nil, err
	}
	return &hive.N{Deleted: got}, nil
}
