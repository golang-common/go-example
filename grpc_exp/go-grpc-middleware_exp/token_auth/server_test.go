// @Author: Perry
// @Date  : 2020/1/21
// @Desc  : 

package token_auth

import (
	"context"
	pb "dpy/exp/grpc_exp/proto"
	"fmt"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net"
	"testing"
)

type HelloService struct{}

func (h HelloService) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{
		Message: "Hello, " + in.Name,
	}, nil
}

func authFunc(ctx context.Context) (context.Context, error) {
	token, err := grpc_auth.AuthFromMD(ctx, "dpy")
	if err != nil {
		return nil, err
	}
	fmt.Printf("receive token: %s\n", token)
	if token != "dpy_token" {
		return nil, status.Errorf(codes.Unauthenticated, "invalid token")
	}
	newCtx := context.WithValue(ctx, "result", "ok")
	return newCtx, nil
}

func TestServer(t *testing.T) {
	s := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_auth.UnaryServerInterceptor(authFunc)),
	)
	pb.RegisterHelloServer(s, HelloService{})

	lis, err := net.Listen("tcp", Address)
	if err != nil {
		panic(err)
	}
	if err := s.Serve(lis); err != nil {
		panic(err)
	}
}
