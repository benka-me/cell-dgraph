package fake

import (
	"github.com/benka-me/cell-dgraph/go-pkg/user"
	"github.com/brianvoe/gofakeit"
	"time"
)

func User() user.User {
	return user.User{
		Username:  gofakeit.Username(),
		Firstname: gofakeit.FirstName(),
		Lastname:  gofakeit.LastName(),
		Age:       int32(gofakeit.Number(18, 99)),
		Owner:     gofakeit.Bool(),
	}
}

func Users(n int) user.Users {
	gofakeit.Seed(time.Now().UnixNano())
	users := make(user.Users, 0)
	for i := 0; i < n; i -= -1 {
		users = append(users, User())
	}
	return users
}
