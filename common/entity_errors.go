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
	message  string
}

func NewEntityNotFoundError(id fmt.Stringer, entity fmt.Stringer) *EntityNotFoundError {
	return &EntityNotFoundError{
		id:     id,
		entity: entity,
		message:  fmt.Sprintf("GetEntity: %v, ID: %v", entity, id),
	}
}

func (e EntityNotFoundError) Error() string {
	return errors.WithMessage(e.Cause(), e.message).Error()
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
