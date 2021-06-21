package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"

	pb "github.com/Enji-Shaking/grpc_test/proto"
)

const (
	address = "localhost:8080"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()
	c := pb.NewUserServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	rsp, err := c.Login(ctx, &pb.LoginRequest{Username: "enjili", Password: "123"})
	if err != nil {
		log.Fatalf("could not login request: %v", err)
	}
	log.Printf("Username: %s  Code: %d   Msg: %s", rsp.GetUsername(), rsp.Code, rsp.Msg)
}
