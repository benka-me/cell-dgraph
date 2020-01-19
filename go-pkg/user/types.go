package user

import (
	"github.com/benka-me/cell-dgraph/go-pkg/conn"
	"github.com/benka-me/hive/go-pkg/hive"
)

const Uid_ = "uid"

type (
	UserQuery interface {
		Get(dgraph conn.Dgraph) (Users, error)
		Set(dgraph conn.Dgraph) (Uids, error)
		Delete(dgraph conn.Dgraph) (int, error)
		SetAndGet(dgraph conn.Dgraph) (Users, error)
		DelAndGet(dgraph conn.Dgraph) (Users, error)
	}
	RetUsers struct {
		Users Users `json:"users"`
	}
	res struct {
		Res []map[string]Uid `json:"res"`
	}

	Uid      string
	User     hive.User
	Username string

	Users []User
	Uids  []Uid

	NewUser      hive.NewUser
	LoginRequest hive.LoginRequest
)

