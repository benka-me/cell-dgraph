package tests

import (
	"fmt"
	"github.com/benka-me/cell-dgraph/go-pkg/conn"
	"github.com/benka-me/cell-dgraph/go-pkg/fake"
	"github.com/benka-me/cell-dgraph/go-pkg/user"
	"github.com/google/go-cmp/cmp"
	"testing"
)

var dgh conn.Dgraph
var err error

func init() {
	dgh.Dgraph, err = conn.NewClient()
}

var usersGood = user.Users{
	user.User{Username: "Cjaou", Firstname: "Claire", Lastname: "Jaouen", Age: 27, Owner: false},
	user.User{Username: "Cjaou28", Firstname: "Claire", Lastname: "Jaouen", Age: 43, Owner: true},
	user.User{Username: "C", Firstname: "", Lastname: "", Age: 43, Owner: false},
	user.User{Username: "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"},
	user.User{Username: "❤️"},
	user.User{Owner: true},
}

var usersBad = user.Users{
	user.User{},
	user.User{Username: ""},
}

/**
Test newUser/login.
*/

func TestLogin(t *testing.T) {
	var tests = []struct {
		nu      user.NewUser
		wantErr bool
	}{
		{user.NewUser{Username: "Cjaou", Password: "Apassword", Email: "fasdfasdf@fasdfa.fd"}, false},
		{user.NewUser{Username: "benka", Password: "Apaaaa", Email: "fasdfasdf@fasdfa.fd"}, false},
		{user.NewUser{Username: "benka22", Password: "A", Email: "fasdfasdf@fasdfa.fd"}, true},
	}
	for _, tt := range tests {
		nu, err := tt.nu.NewUser(dgh)
		if (err != nil) != tt.wantErr {
			t.Error(err)
		} else {
			fmt.Println("New User: ", nu)

			lg := user.LoginRequest{Username: tt.nu.Username, Password: tt.nu.Password}
			ul, err := lg.Login(dgh)
			if (err != nil) != tt.wantErr {
				t.Error(err)
			}
			fmt.Println("User logged: ", ul)

			lg = user.LoginRequest{Username: tt.nu.Username, Password: tt.nu.Password + "4"}
			ul, err = lg.Login(dgh)
			if err == nil {
				t.Error(err)
			}
			fmt.Println("User not logged: ", err, " OK")

			_, err = user.Username(nu.Username).Delete(dgh)
			if err != nil {
				t.Error(err)
			}
		}
	}
}

func TestNewUser(t *testing.T) {
	var tests = []user.NewUser{
		{Username: "Cjaou", Password: "Apassword", Email: "fasdfasdf@fasdfa.fd"},
	}
	for _, tt := range tests {
		u, err := tt.NewUser(dgh)
		if err != nil {
			t.Error(err)
		} else {
			_, err := u.Delete(dgh)
			if err != nil {
				t.Error(err)
			}
		}
	}
}

/**
Test adding n fake users, count them, delete them and check the difference.
*/
func Test1(t *testing.T) {
	nStart := user.CountHas(dgh, "Username")
	nFake := 20

	fakes := fake.Users(nFake)
	if len(fakes) != nFake {
		t.Errorf("fakeUsers returned bad number of fakes")
	}

	fakeAdded, err := fakes.Set(dgh)
	if err != nil {
		t.Error(err)
	}

	nWithFakes := user.CountHas(dgh, "Username")
	if len(fakeAdded)+nStart != nWithFakes {
		t.Error("has Username != nStart + lent(fakeAdded)", len(fakeAdded), nStart, nWithFakes)
	}

	nDeleted, err := fakeAdded.Delete(dgh)
	if err != nil {
		t.Error(err)
	} else if nDeleted != nWithFakes-nStart {
		t.Error("nDeleted != nWithFakes - nStart", nDeleted, nWithFakes, nStart)
	}
}

func TestUsername_DelAndGet(t *testing.T) {
	tests := []*struct {
		name string
		u    user.Username
	}{
		{"Del Cjaou", "Cjaou"},
		{"Del Benka", "Benka"},
	}
	for _, tt := range tests {
		t.Run("DEL/"+tt.name, func(t *testing.T) {
			got, err := tt.u.DelAndGet(dgh)
			if err != nil {
				t.Error(err)
			} else {
				fmt.Println(got)
			}
		})
	}
}

func TestUserGetDel(t *testing.T) {
	tests := []*struct {
		name       string
		u          user.User
		wantSet    user.Users
		wantDel    user.Users
		wantErrSet bool
		wantErrDel bool
	}{
		{"Full fields valid", usersGood[0], user.Users{usersGood[0]}, user.Users{user.User{}}, false, false},
		{"Full fields valid same name", usersGood[1], user.Users{usersGood[1]}, user.Users{user.User{}}, false, false},
		{"Username, Firstname, Lastname, Age", usersGood[2], user.Users{usersGood[2]}, user.Users{user.User{}}, false, false},
		{"Username only, too long char", usersGood[3], user.Users{usersGood[3]}, user.Users{user.User{}}, false, false},
		{"Username only, ❤️", usersGood[4], user.Users{usersGood[4]}, user.Users{user.User{}}, false, false},
		{"Owner only️", usersGood[5], user.Users{usersGood[5]}, user.Users{user.User{}}, false, false},
		{"EMPTY Object", usersBad[0], user.Users{}, user.Users{}, true, true},
		{"Username EMPTY", usersBad[1], user.Users{}, user.Users{}, true, true},
	}
	for _, tt := range tests {
		t.Run("SET/"+tt.name, func(t *testing.T) {
			got, err := tt.u.SetAndGet(dgh)
			if (err != nil) != tt.wantErrSet {
				t.Errorf("Set() error = %v, wantErr %v", err, tt.wantErrSet)
				return
			}
			// add Id to the wantSet var for the comparison && save the id on tt.u for the DelAndGet test
			for i, _ := range tt.wantSet {
				if len(got) > i {
					tt.wantSet[i].Id = got[i].Id
					tt.u.Id = got[i].Id
				}
			}
			if len(got) != len(tt.wantSet) {
				t.Errorf("Set() got (%v %T %d), want (%v %T %d)", got, got, len(got), tt.wantSet, tt.wantSet, len(tt.wantSet))
			} else {
				for i, _ := range got {
					if !cmp.Equal(got[i], tt.wantSet[i]) {
						t.Errorf("Set() got (%v %T), want (%v %T)", got[i], got[i], tt.wantSet[i], tt.wantSet[i])
					}
				}
			}
		})
	}
	for _, tt := range tests {
		t.Run("DEL/"+tt.name, func(t *testing.T) {
			got, err := tt.u.DelAndGet(dgh)
			if (err != nil) != tt.wantErrDel {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErrDel)
				t.Error(tt.u)
				return
			}
			if len(got) != len(tt.wantDel) {
				t.Errorf("Delete() got (%v %T %d), want (%v %T %d)", got, got, len(got), tt.wantDel, tt.wantDel, len(tt.wantDel))
			} else {
				for i, _ := range got {
					tt.wantDel[i].Id = got[i].Id
					if !cmp.Equal(got[i], tt.wantDel[i]) {
						t.Errorf("Delete() got (%v %T), want (%v %T)", got[i], got[i], tt.wantDel[i], tt.wantDel[i])
					}
				}
			}
		})
	}
}
