package user

import (
	"context"
	"encoding/json"
	conn2 "github.com/benka-me/PaasEnger/services/db/pkg/conn"
	"github.com/dgraph-io/dgo"
	"github.com/dgraph-io/dgo/protos/api"
	"github.com/pkg/errors"
)

/**
****************************************************************************************************
BY User
****************************************************************************************************
*/

func (u User) Set(dgraph conn2.Dgraph) (Uid, error) {
	var id Uid = ""
	txn := dgraph.NewTxn()
	defer txn.Discard(context.TODO())

	mu := &api.Mutation{
		CommitNow: true,
	}
	pb, err := json.Marshal(u)
	if err != nil {
		return id, err
	}

	mu.SetJson = pb
	response, err := txn.Mutate(context.Background(), mu)
	if err != nil {
		return id, err
	}

	for _, uid := range response.Uids {
		id = Uid(uid)
	}
	return id, err
}

func (u User) Delete(dgraph conn2.Dgraph) (int32, error) {
	mp := map[string]string{"uid": u.Id}

	pb, err := json.Marshal(mp)
	if err != nil {
		return 0, err
	}

	mu := &api.Mutation{
		CommitNow:  true,
		DeleteJson: pb,
	}

	dgo.DeleteEdges(mu, u.Id, User{}.PredicatesFields()...)
	_, err = dgraph.NewTxn().Mutate(context.Background(), mu)
	if err != nil {
		return 0, err
	}

	return 1, nil
}

func (u User) Get(dgraph conn2.Dgraph) (Users, error) {
	if len(u.Id) > 2 {
		return Uid(u.Id).Get(dgraph)
	} else if len(u.Username) > 3 {
		return Username(u.Username).Get(dgraph)
	} else {
		return Users{User{}}, errors.New("empty user request")
	}
}

func (u User) SetAndGet(dgraph conn2.Dgraph) (Users, error) {
	id, err := u.Set(dgraph)
	if err != nil {
		return nil, err
	}
	return id.Get(dgraph)
}

func (u User) DelAndGet(dgraph conn2.Dgraph) (Users, error) {
	_, err := u.Delete(dgraph)
	if err != nil {
		return nil, err
	}
	return u.Get(dgraph)
}
