// @Author: Perry
// @Date  : 2020/1/20
// @Desc  : 

package token_test

import (
	"context"
	pb "dpy/exp/grpc_exp/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/status"
	"log"
	"testing"
)

const (
	Address = "127.0.0.1:50052"
	OpenTLS = false
)

type customCredential struct{}

func (c customCredential) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		"appid":  "101010",
		"appkey": "i am a key",
	}, nil
}
func (c customCredential) RequireTransportSecurity() bool {
	if OpenTLS {
		return true
	}
	return false
}

func TestClient(t *testing.T) {
	var opts []grpc.DialOption
	if OpenTLS {
		// TLS连接
		creds, err := credentials.NewClientTLSFromFile("../../keys/server.pem", "dpytest")
		if err != nil {
			log.Fatal(err)
		}
		opts = append(opts, grpc.WithTransportCredentials(creds))
	} else {
		opts = append(opts, grpc.WithInsecure())
	}
	// 使用自定义认证
	opts = append(opts, grpc.WithPerRPCCredentials(new(customCredential)))

	conn, err := grpc.Dial(Address, opts...)
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	c := pb.NewHelloClient(conn)

	reqBody := new(pb.HelloRequest)
	reqBody.Name = "gRPC"
	r, err := c.SayHello(context.Background(), reqBody)
	if err != nil {
		if stat, ok := status.FromError(err); ok {
			log.Fatalf("%+v\n", stat.Err())
		} else {
			log.Fatalln(err)
		}
	}
	log.Println(r.Message)
}
