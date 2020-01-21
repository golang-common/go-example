// @Author: Perry
// @Date  : 2020/1/20
// @Desc  : 证书认证客户端

package tls_test

import (
	"context"
	pb "dpy/exp/grpc_exp/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"testing"
)

const Address = "127.0.0.1:50052"

func TestClient(t *testing.T) {
	creds, err := credentials.NewClientTLSFromFile("../../keys/server.pem", "dpytest")
	if err != nil {
		log.Fatalf("failed to create TLS credentials %v", err)
	}
	conn, err := grpc.Dial(Address, grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	c := pb.NewHelloClient(conn)

	reqBody := new(pb.HelloRequest)
	reqBody.Name = "gRpc"
	r, err := c.SayHello(context.Background(), reqBody)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(r.Message)
}
