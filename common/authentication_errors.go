package common

import "github.com/pkg/errors"

type ErrNotAuthorized struct {
	cause error
}

func NewNotAuthorizedError() ErrNotAuthorized {
	return ErrNotAuthorized{cause: errors.New("not authorized")}
}

func (n ErrNotAuthorized) Error() string {
	return n.cause.Error()
}

func (n ErrNotAuthorized) Cause() error {
	return n.cause
}

type ErrInsufficientPermissions struct {
	cause error
}

func NewInsufficientPermissionsError(msg string) ErrInsufficientPermissions {
	err := errors.New("insufficient ermissions")
	if msg == "" {
		return ErrInsufficientPermissions{
			cause: err,
		}
	}
	return ErrInsufficientPermissions{
		cause: errors.WithMessage(err, msg),
	}
}

func (i ErrInsufficientPermissions) Error() string {
	return i.cause.Error()
}

func (i ErrInsufficientPermissions) Cause() error {
	return i.cause
}
