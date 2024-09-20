package gerr

import (
	"github.com/pkg/errors"
	"google.golang.org/grpc/status"
)

func Parse(err error) string {
	causeErr := errors.Cause(err)                      // err类型
	if gStatus, ok := status.FromError(causeErr); ok { // grpc err错误
		return gStatus.Message()
	}
	return ""
}
