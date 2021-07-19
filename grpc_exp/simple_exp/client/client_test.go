// @Author: Perry
// @Date  : 2020/1/1
// @Desc  : 

package client

import (
	"context"
	pb "golang_exp/grpc_exp/simple_exp/helloworld"
	"google.golang.org/grpc"
	"log"
	"os"
	"testing"
	"time"
)

const (
	address     = "localhost:50051"
	defaultName = "Daipengyuan"
)

func Test1(t *testing.T) {
	//建立链接
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)
	v := pb.NewGreeteeClient(conn)
	// Contact the server and print out its response.
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	// 1秒的上下文
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SayHello(ctx, &pb.HelloGoRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.Message)
	r2, err2 := v.SayHello(ctx, &pb.HelloRequest{Name: name})
	if err2 != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r2.Message)
}
