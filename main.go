package main

import (
	"encoding/json"
	"fmt"
	"sera/ops/domain"
)

func main() {
	user := domain.NewUser()
	Output(user)

	user.Update(domain.WithFirstName("Jamie"))
	user.Update(domain.WithMiddleName("Michael"))
	Output(user)

	user.Delete()
	Output(user)
}

func Output(user *domain.User) {
	b, err := json.MarshalIndent(user, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))
}
