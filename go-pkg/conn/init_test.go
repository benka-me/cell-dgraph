package conn

import (
	"fmt"
	user2 "github.com/benka-me/PaasEnger/services/db/pkg/user"
	"testing"
)

var dgh Dgraph
var err error

func init() {
	Dgraph, err = NewClient()
}

func TestReset(t *testing.T) {
	err := InitDgraph(Dgraph, true)
	if err != nil {
		t.Error(err)
	}
}

func TestInitSchema(t *testing.T) {
	err := InitDgraph(Dgraph, false)
	if err != nil {
		t.Error(err)
	}
}

func TestGetHas(t *testing.T) {
	ret, err := user2.GetHas(dgh, "Username")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(ret)
}

func _TestDeleteHasUsername(t *testing.T) {
	err := user2.DeleteHas(dgh, "Username")
	if err != nil {
		t.Error(err)
	}
}
