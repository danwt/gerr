package gerr

import (
	"fmt"
	"net/http"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

/*
See https://github.com/googleapis/googleapis/blob/master/google/rpc/code.proto
See also package doc
*/

type Error struct {
	*status.Status
}

// Error returns a business logic friendly error. This is NOT the same as what
// you get from calling e.Status.Err().Error() because it does not include
// a reference to 'rpc'.
func (e Error) Error() string {
	return e.Status.Message()
}

func (e Error) HttpCode() int {
	return toHttp(e.Code())
}

// GrpcCode returns the grpc code. Named differently to help differentiate with http code.
func (e Error) GrpcCode() codes.Code {
	return e.Code()
}

// GRPCStatus returns the Status represented by se.
func (e Error) GRPCStatus() *status.Status {
	return e.Status
}

func toHttp(code codes.Code) int {
	switch code {
	case codes.OK:
		return http.StatusOK
	case codes.InvalidArgument:
		return http.StatusBadRequest
	case codes.FailedPrecondition:
		return http.StatusBadRequest
	case codes.OutOfRange:
		return http.StatusBadRequest
	case codes.Unauthenticated:
		return http.StatusUnauthorized
	case codes.PermissionDenied:
		return http.StatusForbidden
	case codes.NotFound:
		return http.StatusNotFound
	case codes.AlreadyExists:
		return http.StatusConflict
	case codes.Aborted:
		return http.StatusConflict
	case codes.ResourceExhausted:
		return http.StatusTooManyRequests
	case codes.Canceled:
		return 499 // Client closed request
	case codes.Unknown:
		return http.StatusInternalServerError
	case codes.Internal:
		return http.StatusInternalServerError
	case codes.DataLoss:
		return http.StatusInternalServerError
	case codes.Unimplemented:
		return http.StatusNotImplemented
	case codes.Unavailable:
		return http.StatusServiceUnavailable
	case codes.DeadlineExceeded:
		return http.StatusGatewayTimeout
	}
	panic(fmt.Sprintf("unknown code: %d", code))
}

// ErrCancelled : The operation was cancelled, typically by the caller.
//
// HTTP Mapping: 499 Client Closed Request
var ErrCancelled = Error{Status: status.New(codes.Canceled, "cancelled")}

// ErrUnknown : Unknown error.  For example, this error may be returned when
// a `Status` value received from another address space belongs to
// an error space that is not known in this address space.  Also
// errors raised by APIs that do not return enough error information
// may be converted to this error.
//
// HTTP Mapping: 500 Internal Server Error
var ErrUnknown = Error{Status: status.New(codes.Unknown, "unknown")}

// ErrInvalidArgument : The client specified an invalid argument.  Note that this differs
// from `FAILED_PRECONDITION`.  `INVALID_ARGUMENT` indicates arguments
// that are problematic regardless of the state of the system
// (e.g., a malformed file name).
//
// HTTP Mapping: 400 Bad Request
var ErrInvalidArgument = Error{Status: status.New(codes.InvalidArgument, "invalid argument")}

// ErrDeadlineExceeded : The deadline expired before the operation could complete. For operations
// that change the state of the system, this error may be returned
// even if the operation has completed successfully.  For example, a
// successful response from a server could have been delayed long
// enough for the deadline to expire.
//
// HTTP Mapping: 504 Gateway Timeout
var ErrDeadlineExceeded = Error{Status: status.New(codes.DeadlineExceeded, "deadline exceeded")}

// ErrNotFound : Some requested entity (e.g., file or directory) was not found.
//
// Note to server developers: if a request is denied for an entire class
// of users, such as gradual feature rollout or undocumented allowlist,
// `NOT_FOUND` may be used. If a request is denied for some users within
// a class of users, such as user-based access control, `PERMISSION_DENIED`
// must be used.
//
// HTTP Mapping: 404 Not Found
var ErrNotFound = Error{Status: status.New(codes.NotFound, "not found")}

// ErrAlreadyExists : The entity that a client attempted to create (e.g., file or directory)
// already exists.
//
// HTTP Mapping: 409 Conflict
var ErrAlreadyExists = Error{Status: status.New(codes.AlreadyExists, "already exists")}

// ErrPermissionDenied : The caller does not have permission to execute the specified
// operation. `PERMISSION_DENIED` must not be used for rejections
// caused by exhausting some resource (use `RESOURCE_EXHAUSTED`
// instead for those errors). `PERMISSION_DENIED` must not be
// used if the caller can not be identified (use `UNAUTHENTICATED`
// instead for those errors). This error code does not imply the
// request is valid or the requested entity exists or satisfies
// other pre-conditions.
//
// HTTP Mapping: 403 Forbidden
var ErrPermissionDenied = Error{Status: status.New(codes.PermissionDenied, "permission denied")}

// ErrUnauthenticated : The request does not have valid authentication credentials for the
// operation.
//
// HTTP Mapping: 401 Unauthorized
var ErrUnauthenticated = Error{Status: status.New(codes.Unauthenticated, "unauthenticated")}

// ErrResourceExhausted : Some resource has been exhausted, perhaps a per-user quota, or
// perhaps the entire file system is out of space.
//
// HTTP Mapping: 429 Too Many Requests
var ErrResourceExhausted = Error{Status: status.New(codes.ResourceExhausted, "resource exhausted")}

// ErrFailedPrecondition : The operation was rejected because the system is not in a state
// required for the operation's execution.  For example, the directory
// to be deleted is non-empty, an rmdir operation is applied to
// a non-directory, etc.
//
// Service implementors can use the following guidelines to decide
// between `FAILED_PRECONDITION`, `ABORTED`, and `UNAVAILABLE`:
//
//	(a) Use `UNAVAILABLE` if the client can retry just the failing call.
//	(b) Use `ABORTED` if the client should retry at a higher level. For
//	    example, when a client-specified test-and-set fails, indicating the
//	    client should restart a read-modify-write sequence.
//	(c) Use `FAILED_PRECONDITION` if the client should not retry until
//	    the system state has been explicitly fixed. For example, if an "rmdir"
//	    fails because the directory is non-empty, `FAILED_PRECONDITION`
//	    should be returned since the client should not retry unless
//	    the files are deleted from the directory.
//
// HTTP Mapping: 400 Bad Request
var ErrFailedPrecondition = Error{Status: status.New(codes.FailedPrecondition, "unmet precondition")}

// ErrAborted : The operation was aborted, typically due to a concurrency issue such as
// a sequencer check failure or transaction abort.
//
// See the guidelines above for deciding between `FAILED_PRECONDITION`,
// `ABORTED`, and `UNAVAILABLE`.
//
// HTTP Mapping: 409 Conflict
var ErrAborted = Error{Status: status.New(codes.Aborted, "aborted")}

// ErrOutOfRange : The operation was attempted past the valid range.  E.g., seeking or
// reading past end-of-file.
//
// Unlike `INVALID_ARGUMENT`, this error indicates a problem that may
// be fixed if the system state changes. For example, a 32-bit file
// system will generate `INVALID_ARGUMENT` if asked to read at an
// offset that is not in the range [0,2^32-1], but it will generate
// `OUT_OF_RANGE` if asked to read from an offset past the current
// file size.
//
// There is a fair bit of overlap between `FAILED_PRECONDITION` and
// `OUT_OF_RANGE`.  We recommend using `OUT_OF_RANGE` (the more specific
// error) when it applies so that callers who are iterating through
// a space can easily look for an `OUT_OF_RANGE` error to detect when
// they are done.
//
// HTTP Mapping: 400 Bad Request
var ErrOutOfRange = Error{Status: status.New(codes.OutOfRange, "out of range")}

// ErrUnimplemented : The operation is not implemented or is not supported/enabled in this
// service.
//
// HTTP Mapping: 501 Not Implemented
var ErrUnimplemented = Error{Status: status.New(codes.Unimplemented, "unimplemented")}

// ErrInternal : Internal errors.  This means that some invariants expected by the
// underlying system have been broken.  This error code is reserved
// for serious errors.
//
// HTTP Mapping: 500 Internal Server Error
var ErrInternal = Error{Status: status.New(codes.Internal, "internal")}

// ErrUnavailable : The service is currently unavailable.  This is most likely a
// transient condition, which can be corrected by retrying with
// a backoff. Note that it is not always safe to retry
// non-idempotent operations.
//
// See the guidelines above for deciding between `FAILED_PRECONDITION`,
// `ABORTED`, and `UNAVAILABLE`.
//
// HTTP Mapping: 503 Service Unavailable
var ErrUnavailable = Error{Status: status.New(codes.Unavailable, "unavailable")}

// ErrDataLoss : Unrecoverable data loss or corruption.
//
// HTTP Mapping: 500 Internal Server Error
var ErrDataLoss = Error{Status: status.New(codes.DataLoss, "data loss")}
