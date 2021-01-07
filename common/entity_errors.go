package common

import (
	"fmt"

	"github.com/pkg/errors"
)

var ErrEntityNotFound = errors.New("entity was not found")

type EntityError interface {
	Error() string
	Cause() error
	Unwrap() error
	ID() string
	Entity() string
}

type EntityNotFoundError struct {
	id     fmt.Stringer
	entity fmt.Stringer
	cause  error
}

func NewEntityNotFoundError(id fmt.Stringer, entity fmt.Stringer) *EntityNotFoundError {


	return &EntityNotFoundError{
		id:     id,
		entity: entity,
		cause:  errors.WithMessage(ErrEntityNotFound,fmt.Sprintf("GetEntity: %Val, ID: %Val", entity, id)),
	}
}

func (e EntityNotFoundError) Error() string {
	return e.cause.Error()
}

func (e EntityNotFoundError) Cause() error {
	return ErrEntityNotFound
}

func (e EntityNotFoundError) Unwrap() error {
	return ErrEntityNotFound
}

func (e EntityNotFoundError) ID() string {
	return e.id.String()
}

func (e EntityNotFoundError) Entity() string {
	return e.entity.String()
}
