package user

import (
	"github.com/benka-me/PaasEnger/service-hive/go/go-proto"
	conn2 "github.com/benka-me/PaasEnger/services/db/pkg/conn"
)

const Uid_ = "uid"

type (
	UserQuery interface {
		Get(dgraph conn2.Dgraph) (Users, error)
		Set(dgraph conn2.Dgraph) (Uids, error)
		Delete(dgraph conn2.Dgraph) (int, error)
		SetAndGet(dgraph conn2.Dgraph) (Users, error)
		DelAndGet(dgraph conn2.Dgraph) (Users, error)
	}
	RetUsers struct {
		Users Users `json:"users"`
	}
	res struct {
		Res []map[string]Uid `json:"res"`
	}

	Uid      string
	User     proto.User
	Username string

	Users []User
	Uids  []Uid

	NewUser      proto.NewUser
	LoginRequest proto.LoginRequest
)

