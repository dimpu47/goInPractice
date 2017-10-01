package main

import (
	"fmt"
	"os"

	"github.com/dimpu47/goInPractice/chap10/user"

	"github.com/ugorji/go/codec"
)

func main() {
	jh := new(codec.JsonHandle)
	u := &user.User{
		Name:  "Gaurav Choudhary",
		Email: "Gaurav@example.com",
	}
	var out []byte
	err := codec.NewEncoderBytes(&out, jh).Encode(u)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(string(out))
}
