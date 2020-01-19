package room

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/benka-me/cell-dgraph/go-pkg/conn"
	"github.com/benka-me/hive/go-pkg/hive"
	"github.com/dgraph-io/dgo/protos/api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Insert(handshake *hive.HandshakeReq, dgraph conn.Dgraph) (*hive.Room, error) {
	var ret hive.Room
	mut := struct {
		Room string `json:"Room"`
	}{
		Room: handshake.Name + " request",
	}

	txn := dgraph.NewTxn()
	defer txn.Discard(context.TODO())
	mu := &api.Mutation{
		CommitNow: true,
	}

	pb, err := json.Marshal(mut)
	if err != nil {
		return &ret, status.Error(codes.Internal, "new room: error marshall")
	}

	mu.SetJson = pb
	response, err := txn.Mutate(context.Background(), mu)
	if err != nil {
		return &ret, status.Error(codes.Internal, err.Error())
	}
	for _, uid := range response.GetUids() {
		ret.Id = uid
	}

	return &ret, nil
}

func edge(targetUid string, roomUid string, dgraph conn.Dgraph) error {
	fmt.Println("Link to room: ", targetUid, roomUid)
	mut := struct {
		Uid  string `json:"uid"`
		Join struct {
			Uid string `json:"uid"`
		} `json:"join"`
	}{
		Uid: targetUid,
	}
	mut.Join.Uid = roomUid

	txn := dgraph.NewTxn()
	defer txn.Discard(context.TODO())
	mu := &api.Mutation{
		CommitNow: true,
	}

	pb, err := json.Marshal(mut)
	if err != nil {
		return status.Error(codes.Internal, "link to room: error marshall")
	}

	mu.SetJson = pb
	_, err = txn.Mutate(context.Background(), mu)
	if err != nil {
		return status.Error(codes.Internal, err.Error())
	}
	return status.Error(codes.OK, "")
}

func InitEdges(handshake *hive.HandshakeReq, customerUid string, dgraph conn.Dgraph) (*hive.Room, error) {
	var ret *hive.Room
	ret, err := Insert(handshake, dgraph)
	if err != nil {
		return ret, nil
	}

	err = edge(handshake.ToUserId, ret.GetId(), dgraph)
	if err != nil {
		return ret, nil
	}

	err = edge(customerUid, ret.GetId(), dgraph)
	if err != nil {
		return ret, nil
	}
	return ret, nil
}
