package service

import "errors"

// Generic errors  that can be returned from services
var (
	ErrBadRequest      = errors.New("bad request")
	ErrInternalFailure = errors.New("internal failure")
	ErrNotFound        = errors.New("not found")
)

// Error is a custom error type that combines application and service errors
type Error struct {
	appError error
	svcError error
}

func (e Error) AppError() error {
	return e.appError
}

func (e Error) SvcError() error {
	return e.svcError
}

// NewError creates a new Error
func NewError(svcErr, appErr error) Error {
	return Error{
		appError: appErr,
		svcError: svcErr,
	}
}

// Error returns the combined error message
func (e Error) Error() string {
	return errors.Join(e.svcError, e.appError).Error()
}
