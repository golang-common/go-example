// @Author: Perry
// @Date  : 2020/1/1
// @Desc  : 

package server

import (
	"context"
	pb "golang_exp/grpc_exp/simple_exp/helloworld"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"testing"
)

const port = ":50051"

type greetee struct {
}

func (s *greetee) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("message received: %s", in.Name)
	return &pb.HelloReply{Message: "Hello" + in.Name}, nil
}

type greeter struct {
}

func (s *greeter) SayHello(ctx context.Context, in *pb.HelloGoRequest) (*pb.HelloGoReply, error) {
	log.Printf("message received: %s", in.Name)
	return &pb.HelloGoReply{Message: "Hello" + in.Name}, nil
}

func Test1(t *testing.T) {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen:%v", err)
	} else {
		log.Print("port listened")
	}
	s := grpc.NewServer()
	pb.RegisterGreeteeServer(s, &greetee{})
	pb.RegisterGreeterServer(s, &greeter{})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve:%v", err)
	} else {
		log.Print("server started")
	}
}
