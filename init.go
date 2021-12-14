package authzclient

import (
	pb "github.com/es-hs/erpc/authz"
	"google.golang.org/grpc"
)

var (
	Conn *grpc.ClientConn
	C    pb.AuthzRPCClient
)

func InitAuthClient(target string, opts ...grpc.DialOption) error {
	var err error
	Conn, err = grpc.Dial(target, opts...)
	if err != nil {
		return err
	}
	C = pb.NewAuthzRPCClient(Conn)
	return nil
}
