// @Author: Perry
// @Date  : 2020/1/20
// @Desc  : 

package token_test

import (
	"context"
	pb "dpy/exp/grpc_exp/proto"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"log"
	"net"
	"testing"
)

var HelloService = helloService{}

type helloService struct{}

func (h helloService) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	// 解析metadata中的信息并验证
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "无token认证in信息")
	}
	var appid, appkey string
	if val, ok := md["appid"]; ok {
		appid = val[0]
	}
	if val, ok := md["appkey"]; ok {
		appkey = val[0]
	}
	if appid != "101010" || appkey != "i am a key" {
		return nil, status.Errorf(codes.Unauthenticated, "token认证信息无效:appid=%s,appkey=%s", appid, appkey)
	}
	resp := new(pb.HelloReply)
	resp.Message = fmt.Sprintf("Hello %s.\nToken info: appid=%s,appkey=%s", in.Name, appid, appkey)

	return resp, nil
}

func TestServer(t *testing.T) {
	listen, err := net.Listen("tcp", Address)
	if err != nil {
		log.Fatalf("failed to listen: %v\n", err)
	}

	creds, err := credentials.NewServerTLSFromFile("../../keys/server.pem", "../../keys/server.key")
	if err != nil {
		log.Fatalf("failed to generate credentials %v", err)
	}

	s := grpc.NewServer(grpc.Creds(creds))
	pb.RegisterHelloServer(s, HelloService)
	log.Println("listen on " + Address + " with TLS + Token")

	s.Serve(listen)
}
