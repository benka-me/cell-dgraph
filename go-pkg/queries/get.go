package queries

import (
	"context"
	"encoding/json"
	"fmt"
	conn2 "github.com/benka-me/PaasEnger/services/db/pkg/conn"
)

func CheckPassword(uid, pwd string, dgraph conn2.Dgraph) bool {
	txn := dgraph.NewTxn()
	defer txn.Discard(context.TODO())

	q := fmt.Sprintf(`
		{
		  check(func: uid(%s)) {
		    name
		    checkpwd(Password, "%s")
		  }
		}
	`, uid, pwd)

	res, err := txn.Query(context.Background(), q)
	if err != nil {
		return false
	}
	var tmp map[string][]map[string]interface{}
	err = json.Unmarshal(res.Json, &tmp)
	if !tmp["check"][0]["checkpwd(Password)"].(bool) {
		return false
	}
	return true
}

func Get(q string, dgraph conn2.Dgraph) ([]byte, error) {
	txn := dgraph.NewTxn()
	defer txn.Discard(context.TODO())

	res, err := txn.Query(context.Background(), q)
	if err != nil {
		return nil, err
	}
	return res.Json, err
}

