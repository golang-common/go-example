// @Author: Perry
// @Date  : 2020/1/21
// @Desc  : 

package token_auth

import (
	"context"
	pb "dpy/exp/grpc_exp/proto"
	"golang.org/x/oauth2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
	"log"
	"testing"
	"time"
)

const (
	Address   = "127.0.0.1:51002"
	AuthToken = "dpy_token"
)

// TokenSource supplies PerRPCCredentials from an oauth2.TokenSource.
type TokenSource struct {
	oauth2.TokenSource
}

// GetRequestMetadata gets the request metadata as a map from a TokenSource.
func (ts TokenSource) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	token, err := ts.Token()
	if err != nil {
		return nil, err
	}
	return map[string]string{
		"authorization": token.Type() + " " + token.AccessToken,
	}, nil
}

// RequireTransportSecurity indicates whether the credentials requires transport security.
func (ts TokenSource) RequireTransportSecurity() bool {
	return false
}

type OAuth2TokenSource struct {
	accessToken string
}

func (ts *OAuth2TokenSource) Token() (*oauth2.Token, error) {
	t := &oauth2.Token{
		AccessToken: ts.accessToken,
		Expiry:      time.Now().Add(1 * time.Hour),
		TokenType:   "dpy",
	}
	return t, nil
}

func TestClient(t *testing.T) {
	tokenCreds := TokenSource{TokenSource: &OAuth2TokenSource{accessToken: AuthToken}}
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	opts = append(opts, grpc.WithPerRPCCredentials(tokenCreds))

	conn, err := grpc.Dial(Address, opts...)
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	c := pb.NewHelloClient(conn)

	reqBody := new(pb.HelloRequest)
	reqBody.Name = "dpy_name"
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
