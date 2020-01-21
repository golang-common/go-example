// @Author: Perry
// @Date  : 2020/1/21
// @Desc  : 直接在context中插入

package token_auth

import (
	"context"
	pb "dpy/exp/grpc_exp/proto"
	"fmt"
	"github.com/grpc-ecosystem/go-grpc-middleware/util/metautils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"log"
	"testing"
)

func ctxWithToken(ctx context.Context, scheme string, token string) context.Context {
	md := metadata.Pairs("authorization", fmt.Sprintf("%s %v", scheme, token))
	nCtx := metautils.NiceMD(md).ToOutgoing(ctx)
	return nCtx
}

func TestClient2(t *testing.T) {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())

	conn, err := grpc.Dial(Address, opts...)
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	c := pb.NewHelloClient(conn)

	reqBody := new(pb.HelloRequest)
	reqBody.Name = "dpy_name"

	ctx := ctxWithToken(context.Background(), "dpy", "dpy_token")

	r, err := c.SayHello(ctx, reqBody)
	if err != nil {
		if stat, ok := status.FromError(err); ok {
			log.Fatalf("%+v\n", stat.Err())
		} else {
			log.Fatalln(err)
		}
	}
	log.Println(r.Message)
}
