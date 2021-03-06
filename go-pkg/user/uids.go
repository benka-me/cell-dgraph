package user

import (
	"errors"
	"github.com/benka-me/cell-dgraph/go-pkg/conn"
)

/**
****************************************************************************************************
BY Uids
****************************************************************************************************
*/
func (uids Uids) Set(d conn.Dgraph) (string, error) {
	return "", errors.New("you cannot insert users by uid")
}
func (uids Uids) Delete(d conn.Dgraph) (int, error) {
	var ret int = 0

	for i, u := range uids {
		_, err := u.Delete(d)
		if err != nil {
			return ret, err
		}
		ret = i + 1
	}

	return ret, nil
}

func (uids Uids) Get(d conn.Dgraph) (Users, error) {
	panic("implement me")
}

func (uids Uids) SetAndGet(d conn.Dgraph) (Users, error) {
	panic("implement me")
}

func (uids Uids) DelAndGet(d conn.Dgraph) (Users, error) {
	panic("implement me")
}

