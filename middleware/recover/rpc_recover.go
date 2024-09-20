package recover

import (
	"alger/common/global"
	"context"
	"google.golang.org/grpc"
)

func RpcRecover(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	defer func() {
		r := recover()
		if r != nil {
			global.Logger.Error("recover err", r)
		}
	}()
	return handler(ctx, req)
}
