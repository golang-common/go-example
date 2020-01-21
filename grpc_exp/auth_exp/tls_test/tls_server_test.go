// @Author: Perry
// @Date  : 2020/1/20
// @Desc  : 证书认证服务器

package tls_test

import (
	"context"
	pb "dpy/exp/grpc_exp/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"net"
	"testing"
)

var HelloService = helloService{}

type helloService struct {
}

func (h helloService) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	resp := new(pb.HelloReply)
	resp.Message = "Hello" + in.Name + "."
	log.Println(resp.Message)
	return resp, nil
}

func TestServer(t *testing.T) {
	listen, err := net.Listen("tcp", Address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	/*tls认证*/
	creds, err := credentials.NewServerTLSFromFile("../../keys/server.pem", "../../keys/server.key")
	if err != nil {
		log.Fatalf("failed to generate credentials %v", err)
	}

	s := grpc.NewServer(grpc.Creds(creds))
	pb.RegisterHelloServer(s, HelloService)

	log.Println("listen on " + Address + " with tls")
	s.Serve(listen)
}
