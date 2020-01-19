package user

import (
	"reflect"
	"strings"
)

type Fields interface {
	PredicatesFields() []string
}

func (u User) PredicatesFields() []string {

	ret := make([]string, 0)

	e := reflect.ValueOf(&u).Elem()
	for i := 0; i < e.NumField(); i++ {
		s := e.Type().Field(i).Name
		if !strings.Contains(s, "XXX") && !strings.Contains(s, "Id") && len(s) > 0 {
			ret = append(ret, s)
		}
	}
	return ret
}
