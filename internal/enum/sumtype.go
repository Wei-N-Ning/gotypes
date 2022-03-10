package enum

// ProdError defines the error typeset
type ProdError interface {
	NotFound | InvalidPassword | Timeout | InternalError

	Error() string
	Code() int
}

const (
	CodeNotFound        = 1001
	CodeInvalidPassword = 1002
	CodeTimeout         = 1003
	CodeInternalError   = 9999
)

// The error types

type NotFound struct{}

func (_ NotFound) Error() string { return "not found" }

func (_ NotFound) Code() int { return CodeNotFound }

type InvalidPassword struct{}

func (_ InvalidPassword) Error() string { return "invalid password" }

func (_ InvalidPassword) Code() int { return CodeInvalidPassword }

type Timeout struct{}

func (_ Timeout) Error() string { return "timeout after 90 seconds" }

func (_ Timeout) Code() int { return CodeTimeout }

type InternalError struct{}

func (_ InternalError) Error() string { return "something terrible, but don't know why" }

func (_ InternalError) Code() int { return CodeInternalError }

// Map error to an integer code

func ErrorCode[E ProdError](err E) int {
	v := err.Code()
	switch v {
	case CodeInvalidPassword:
		return 1
	case CodeNotFound:
		return 2
	case CodeTimeout:
		return 3
	case CodeInternalError:
		return 4
	default:
		panic("invalid error type: " + err.Error())
	}
}
