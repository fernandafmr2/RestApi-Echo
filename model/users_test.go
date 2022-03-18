package model

import (
	"crud/config"
	"fmt"
	"testing"
)

func TestCreate(t *testing.T) {
	config.ConnectDB()
	user := new(Users)
	user.Email = "private.sardi@gmail.com"
	user.Name = "Sardi"
	user.Address = "Jalan - "
	user.PhoneNumber = "08xxxx"

	if user.CreateUser() != nil {
		t.Errorf("Failed create user")
	}
	fmt.Println(user)
}

func TestUpdate(t *testing.T) {
	config.ConnectDB()
	user, err := GetOneByName("Sardi")

	if err != nil {
		t.Error(err.Error())
	}

	fmt.Println(user)
	if user.DeleteUser() != nil {
		t.Errorf(err.Error())
	}
}

func TestGet(t *testing.T) {
	config.ConnectDB()
	// user, err := GetOneByEmail("private.sardi@gmail.com")

	// if err != nil {
	// 	t.Errorf("Error get one")
	// }

	// fmt.Println(user)
	users, err := GetAll("p")

	if err != nil {
		t.Errorf(err.Error())
	}

	fmt.Println(users)
}
