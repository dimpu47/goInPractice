package main

import (
	"errors"
	"fmt"
	"math/rand"
)

var ErrTimeout = errors.New("The request is timed out.")
var ErrRejected = errors.New("The request was rejected.")

var random = rand.New(rand.NewSource(35))

func main() {
	resp, err := SendReq("Hello.")

	for err == ErrTimeout {
		fmt.Println("Timeout. Retrying.")
		resp, err = SendReq("Hello.")
	}
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(resp)
	}

}
func SendReq(req string) (string, error) {
	switch random.Int() % 3 {
	case 0:
		return "Success", nil
	case 1:
		return "", ErrRejected
	default:
		return "", ErrTimeout
	}
}
