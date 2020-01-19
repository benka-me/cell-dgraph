package user

import (
	"github.com/benka-me/cell-dgraph/go-pkg/conn"
)

/**
****************************************************************************************************
BY Users
****************************************************************************************************
*/
func (users Users) Get(d conn.Dgraph) (Users, error) {
	panic("implement me")
}

func (users Users) Set(d conn.Dgraph) (Uids, error) {
	ret := make(Uids, 0)
	for _, u := range users {
		s, err := u.Set(d)
		if err != nil {
			return ret, err
		} else {
			ret = append(ret, s)
		}
	}
	return ret, nil
}

func (users Users) Delete(d conn.Dgraph) (int, error) {
	n := 0
	for i, u := range users {
		_, err := u.Delete(d)
		if err != nil {
			return n, err
		}
		n = i
	}
	return n, nil
}

func (users Users) SetAndGet(d conn.Dgraph) (Users, error) {
	panic("implement me")
}

func (users Users) DelAndGet(d conn.Dgraph) (Users, error) {
	panic("implement me")
}
