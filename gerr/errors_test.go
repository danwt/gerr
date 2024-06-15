package gerr

import (
	"testing"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestFoo(t *testing.T) {
	errA := status.Error(codes.Unauthenticated, "")
	errB := status.FromProto()
}
