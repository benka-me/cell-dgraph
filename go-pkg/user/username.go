package user

import (
	"fmt"
	conn2 "github.com/benka-me/PaasEnger/services/db/pkg/conn"
)

/**
****************************************************************************************************
BY Username
****************************************************************************************************
*/
func (u Username) Get(dgraph conn2.Dgraph) (Users, error) {
	q := fmt.Sprintf(`
		{
			users(func: eq(Username, "%s")) {
				uid
				Username
				Email
				Firstname
				Lastname
				Age
				Owner
			}
		}
	`, u)
	ret, err := QueryUsers(q, dgraph)
	if err != nil {
		return ret, err
	}
	return ret, err
}

func (u Username) Set(dgraph conn2.Dgraph) (Uid, error) {
	return User{Username: string(u)}.Set(dgraph)
}

func (u Username) Delete(dgraph conn2.Dgraph) (int32, error) {
	uids, err := u.Get(dgraph)
	if err != nil {
		return 0, err
	}
	n, err := uids.Delete(dgraph)
	return int32(n), err
}

func (u Username) SetAndGet(dgraph conn2.Dgraph) (Users, error) {
	panic("implement me")
}

func (u Username) DelAndGet(dgraph conn2.Dgraph) (Users, error) {
	_, err := u.Delete(dgraph)
	if err != nil {
		return nil, err
	}
	return u.Get(dgraph)
}

