package user

import (
	"errors"
	"fmt"
	"github.com/benka-me/cell-dgraph/go-pkg/conn"
)

/**
****************************************************************************************************
BY Uid
****************************************************************************************************
*/

func (u Uid) Set(dgraph conn.Dgraph) (string, error) {
	return "", errors.New("you cannot insert user by uid")
}

func (u Uid) Delete(dgraph conn.Dgraph) (int32, error) {
	user := User{Id: string(u)}
	return user.Delete(dgraph)
}

func (u Uid) Get(dgraph conn.Dgraph) (Users, error) {
	q := fmt.Sprintf(`
		{
			users(func: uid(%s)) {
				uid
				Username
				Firstname
				Lastname
				Age
				Owner
			}
		}
	`, u)
	ret, err := QueryUsers(q, dgraph)
	ret[0].Id = string(u)
	return ret, err
}
func (u Uid) SetAndGet(d conn.Dgraph) (Users, error) {
	panic("implement me")
}

func (u Uid) DelAndGet(d conn.Dgraph) (Users, error) {
	panic("implement me")
}
