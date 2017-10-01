package main

import (
	"encoding/json"
	"fmt"
)

type Name struct {
	First string `json: "firstname"`
	Last  string `json: "lastname"`
}

func main() {
	n := &Name{"Gauro", "Cairo"}
	data, _ := json.Marshal(n)
	fmt.Printf("%s\n", data)
}
