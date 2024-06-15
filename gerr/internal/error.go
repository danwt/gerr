package internal

import "google.golang.org/grpc/status"

type Error struct {
	*status.Status
}

func (e *Error) Error() string {
	return ""
}
