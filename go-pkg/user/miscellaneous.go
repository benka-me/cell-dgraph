package user

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	conn2 "github.com/benka-me/PaasEnger/services/db/pkg/conn"
	queries2 "github.com/benka-me/PaasEnger/services/db/pkg/queries"
	"github.com/dgraph-io/dgo/protos/api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (lr LoginRequest) Login(dgraph conn2.Dgraph) (User, error) {
	type Check struct {
		Check map[string]bool `json:"check"`
	}
	var user User

	users, err := Username(lr.Username).Get(dgraph)
	if err != nil {
		return user, err
	}

	if len(users) > 0 {
		user = users[0]
	} else {
		return user, errors.New("can't find username")
	}

	if !queries2.CheckPassword(user.Id, lr.Password, dgraph) {
		return User{}, errors.New("password don't match")
	}

	return user, nil
}

func QueryUsers(q string, dgraph conn2.Dgraph) (Users, error) {
	var tmp RetUsers
	jsoned, err := queries2.Get(q, dgraph)
	if err != nil {
		return tmp.Users, err
	}

	err = json.Unmarshal(jsoned, &tmp)
	return tmp.Users, nil
}

func (nu NewUser) NewUser(dgraph conn2.Dgraph) (User, error) {
	exist, _ := Username(nu.Username).Get(dgraph)
	nu.Owner = true

	if len(exist) > 0 {
		err := status.Error(codes.AlreadyExists, "username already exist")
		return User{}, err
	}

	txn := dgraph.NewTxn()
	defer txn.Discard(context.TODO())

	mu := &api.Mutation{
		CommitNow: true,
	}
	pb, err := json.Marshal(nu)
	if err != nil {
		return User{}, err
	}

	mu.SetJson = pb
	response, err := txn.Mutate(context.Background(), mu)
	var ret User
	for _, uid := range response.GetUids() {
		ret, err := Uid(uid).Get(dgraph)
		return ret[0], err
	}
	return ret, err
}

func GetHas(dgraph conn2.Dgraph, has string) (ret Uids, err error) {
	q := fmt.Sprintf(`
		{
			res(func: has(%s)) {
				uid
			}
		}
	`, has)
	var tmp res
	jsoned, err := queries2.Get(q, dgraph)
	if err != nil {
		return
	}

	err = json.Unmarshal(jsoned, &tmp)
	if err != nil {
		return
	}

	ret = make(Uids, len(tmp.Res))
	for i, t := range tmp.Res {
		ret[i] = t[Uid_]
	}
	return
}

func CountHas(dgraph conn2.Dgraph, has string) int {
	count, _ := GetHas(dgraph, has)
	return len(count)
}

func DeleteHas(dgraph conn2.Dgraph, has string) error {
	uids, err := GetHas(dgraph, has)
	if err != nil {
		return err
	}

	_, err = uids.Delete(dgraph)
	return err
}

