// @Author: Perry
// @Date  : 2020/1/20
// @Desc  : 

package interceptor_exp

import (
	"context"
	pb "golang_exp/grpc_exp/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"testing"
)

const (

	// OpenTLS 是否开启TLS认证
	OpenTLS = false
)

// customCredential 自定义认证
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
	var err error
	var opts []grpc.DialOption

	//if OpenTLS {
	//	// TLS连接
	//	creds, err := credentials.NewClientTLSFromFile("../keys/server.pem", "server name")
	//	if err != nil {
	//		grpclog.Fatalf("Failed to create TLS credentials %v", err)
	//	}
	//	opts = append(opts, grpc.WithTransportCredentials(creds))
	//} else {
	opts = append(opts, grpc.WithInsecure())
	//}

	// 指定自定义认证
	opts = append(opts, grpc.WithPerRPCCredentials(new(customCredential)))

	conn, err := grpc.Dial(Address, opts...)

	if err != nil {
		grpclog.Fatalln(err)
	}

	defer conn.Close()

	// 初始化客户端
	c := pb.NewHelloClient(conn)

	// 调用方法
	reqBody := new(pb.HelloRequest)
	reqBody.Name = "gRPC"
	r, err := c.SayHello(context.Background(), reqBody)
	if err != nil {
		grpclog.Fatalln(err)
	}

	grpclog.Println(r.Message)
}
