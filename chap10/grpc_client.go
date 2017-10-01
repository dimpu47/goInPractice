package main

import (
	"fmt"
	"os"

	pb "github.com/dimpu47/goInPractice/chap10/hellopb"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	address := "localhost:55555"
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		fmt.Println("Unable to connect: ", err)
		os.Exit(1)
	}

	defer conn.Close()
	c := pb.NewHelloClient(conn)

	name := "Mundo"

	hr := &pb.HelloRequest{Name: name}
	r, err := c.Say(context.Background(), hr)
	if err != nil {
		fmt.Println("Could not say: ", err)
		os.Exit(1)
	}

	fmt.Println(r.Message)
}
