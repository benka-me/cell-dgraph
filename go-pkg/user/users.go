package user

import (
	conn2 "github.com/benka-me/PaasEnger/services/db/pkg/conn"
)

/**
****************************************************************************************************
BY Users
****************************************************************************************************
*/
func (users Users) Get(d conn2.Dgraph) (Users, error) {
	panic("implement me")
}

func (users Users) Set(d conn2.Dgraph) (Uids, error) {
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

func (users Users) Delete(d conn2.Dgraph) (int, error) {
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

func (users Users) SetAndGet(d conn2.Dgraph) (Users, error) {
	panic("implement me")
}

func (users Users) DelAndGet(d conn2.Dgraph) (Users, error) {
	panic("implement me")
}
