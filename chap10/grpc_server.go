package main

import (
	"fmt"
	"net"
	"os"

	pb "github.com/dimpu47/goInPractice/chap10/hellopb"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type server struct{}

func (s *server) Say(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	msg := "Hola " + in.Name + "!"
	return &pb.HelloResponse{Message: msg}, nil
}

func main() {
	l, err := net.Listen("tcp", ":55555")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	s := grpc.NewServer()
	pb.RegisterHelloServer(s, &server{})
	s.Serve(l)
}
