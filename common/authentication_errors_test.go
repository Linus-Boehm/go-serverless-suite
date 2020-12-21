package common

import (
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestErrorBehavior(t *testing.T) {
	otherError := errors.New("SomeError")
	insufficient := NewInsufficientPermissionsError("test1")
	tests := []struct {
		name   string
		err    error
		target error
		useIs  bool
		want   bool
	}{
		{
			name:   "happy - Is",
			err:    NewInsufficientPermissionsError("test1"),
			target: &ErrInsufficientPermissions{},
			want:   true,
		},
		{
			name:   "happy - As",
			err:    NewInsufficientPermissionsError("test1"),
			target: &ErrInsufficientPermissions{},
			useIs:  false,
			want:   true,
		},
		{
			name:   "error - Is other ",
			err:    otherError,
			target: &insufficient,
			useIs:  true,
			want:   false,
		},
		{
			name:   "error - As other",
			err:    otherError,
			target: &insufficient,
			useIs:  false,
			want:   false,
		},
		{
			name:   "happy - compare Is other",
			err:    NewNotAuthorizedError(),
			target: &insufficient,
			useIs:  true,
			want:   false,
		},
		{
			name:   "happy - compare As other",
			err:    NewNotAuthorizedError(),
			target: &insufficient,
			useIs:  false,
			want:   false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			if test.useIs {
				assert.Equal(t, test.want, errors.Is(test.err, test.target))
			} else {
				assert.Equal(t, test.want, errors.As(test.err, test.target))
			}
		})
	}
}
