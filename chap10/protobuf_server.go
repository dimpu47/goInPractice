package main

import (
	"net/http"

	pb "github.com/dimpu47/goInPractice/chap10/userpb"
	"github.com/golang/protobuf/proto"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(res http.ResponseWriter, req *http.Request) {
	u := &pb.User{
		Name:  proto.String("Gaurav Choudhary"),
		Id:    proto.Int32(12346),
		Email: proto.String("gaurav@example.com"),
	}

	body, err := proto.Marshal(u)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	res.Header().Set("Content-Type", "application/x-protobuf")
	res.Write(body)
}
