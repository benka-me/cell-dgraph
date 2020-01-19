package fake

import (
	user2 "github.com/benka-me/PaasEnger/services/db/pkg/user"
	"github.com/brianvoe/gofakeit"
	"time"
)

func User() user2.User {
	return user2.User{
		Username:  gofakeit.Username(),
		Firstname: gofakeit.FirstName(),
		Lastname:  gofakeit.LastName(),
		Age:       int32(gofakeit.Number(18, 99)),
		Owner:     gofakeit.Bool(),
	}
}

func Users(n int) user2.Users {
	gofakeit.Seed(time.Now().UnixNano())
	users := make(user2.Users, 0)
	for i := 0; i < n; i -= -1 {
		users = append(users, User())
	}
	return users
}
