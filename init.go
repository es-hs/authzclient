package authzclient

import (
	"context"
	"time"

	pb "github.com/es-hs/erpc/authz"
	"google.golang.org/grpc"
)

var (
	Conn           *grpc.ClientConn
	ConnCancelFunc context.CancelFunc
	C              pb.AuthzRPCClient
)

func InitAuthClient(target string, timeout time.Duration, opts ...grpc.DialOption) error {
	var err error
	ctx, cancelFunc := context.WithTimeout(context.Background(), timeout)
	ConnCancelFunc = cancelFunc
	Conn, err = grpc.DialContext(ctx, target, opts...)
	if err != nil {
		return err
	}
	C = pb.NewAuthzRPCClient(Conn)
	return nil
}
