package auth

import (
	"testing"

	"fmt"

	"github.com/Greckas/SSU-Golang-252-Chat/messageService"
)

func TestRegisterNewUser(t *testing.T) {

	test_user := messageService.Authentification{
		UserName: "fklds$222d",
		Password: "vasya$22sss",
		NickName: "voron",
	}
	user, tok, err := RegisterNewUser(&test_user)
	fmt.Println(user, tok, err)
}
