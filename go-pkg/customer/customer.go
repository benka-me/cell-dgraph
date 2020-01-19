package customer

import (
	"context"
	"encoding/json"
	"github.com/benka-me/PaasEnger/service-hive/go/go-proto"
	conn2 "github.com/benka-me/PaasEnger/services/db/pkg/conn"
	room2 "github.com/benka-me/PaasEnger/services/db/pkg/room"
	"github.com/dgraph-io/dgo/protos/api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Handshake(handshake *proto.HandshakeReq, dgraph conn2.Dgraph) (*proto.Welcome, error) {
	wel := &proto.Welcome{
		Customer: &proto.Customer{CustomerName: handshake.GetName()},
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
	wel.Room, err = room2.InitEdges(handshake, wel.Customer.Id, dgraph)

	return wel, err
}
