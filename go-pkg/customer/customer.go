package customer

import (
	"context"
	"encoding/json"
	"github.com/benka-me/cell-dgraph/go-pkg/conn"
	"github.com/benka-me/cell-dgraph/go-pkg/room"
	"github.com/benka-me/hive/go-pkg/hive"
	"github.com/dgraph-io/dgo/protos/api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Handshake(handshake *hive.HandshakeReq, dgraph conn.Dgraph) (*hive.Welcome, error) {
	wel := &hive.Welcome{
		Customer: &hive.Customer{CustomerName: handshake.GetName()},
		Room:     nil,
		Message:  "",
	}

	txn := dgraph.NewTxn()
	defer txn.Discard(context.TODO())
	mu := &api.Mutation{
		CommitNow: true,
	}

	pb, err := json.Marshal(wel.Customer)
	if err != nil {
		return nil, status.Error(codes.Internal, "new customer: error marshall")
	}

	mu.SetJson = pb
	response, err := txn.Mutate(context.Background(), mu)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	for _, uid := range response.GetUids() {
		wel.Customer.Id = uid
	}
	wel.Room, err = room.InitEdges(handshake, wel.Customer.Id, dgraph)

	return wel, err
}
