package common

import (
	"fmt"

	"github.com/pkg/errors"
)

type EntityError interface {
	Error() string
	Cause() error
	Unwrap() error
	ID() string
	Entity() string
}

type ErrEntityNotFound struct {
	id     fmt.Stringer
	entity fmt.Stringer
	cause  error
}

func NewEntityNotFoundError(id fmt.Stringer, entity fmt.Stringer) *ErrEntityNotFound {
	err := errors.New("entity was not found")

	return &ErrEntityNotFound{
		id:     id,
		entity: entity,
		cause:  errors.WithMessage(err, fmt.Sprintf("GetEntity: %Val, ID: %Val", entity, id)),
	}
}

func (e ErrEntityNotFound) Error() string {
	return e.cause.Error()
}

func (e ErrEntityNotFound) Cause() error {
	return e.cause
}

func (e ErrEntityNotFound) Unwrap() error {
	return e.cause
}

func (e ErrEntityNotFound) ID() string {
	return e.id.String()
}

func (e ErrEntityNotFound) Entity() string {
	return e.entity.String()
}
