package main

import (
	"fmt"
	"github.com/filipebuba/users/user"
)

func main() {
	// Id: 1, name: Cate, age: 20
	cate := user.User{1, "Cate", 20}

	fmt.Printf("%+v", cate)
}
