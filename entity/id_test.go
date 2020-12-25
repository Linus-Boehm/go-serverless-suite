package entity

import (
	"fmt"
	"testing"

	"github.com/Linus-Boehm/go-serverless-suite/common"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestIDFromStringOrNil(t *testing.T) {
	uid := uuid.Must(uuid.NewRandom())
	tests := []struct {
		name        string
		input       string
		expectId    *uuid.UUID
		expectVal   *string
		expectedStr string
	}{
		{
			name:        "happy - uuid",
			input:       uid.String(),
			expectId:    &uid,
			expectVal:   nil,
			expectedStr: uid.String(),
		},
		{
			name:        "happy - string",
			input:       "username",
			expectId:    nil,
			expectVal:   common.StringPtr("username"),
			expectedStr: "username",
		},
		{
			name:        "happy - empty",
			input:       "",
			expectId:    nil,
			expectVal:   nil,
			expectedStr: "",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			id := IDFromStringOrNil(test.input)
			if test.expectId == nil {
				assert.Nil(t, id.id)
			} else {
				assert.Equal(t, test.expectId.String(), id.id.String())

			}
			if test.expectVal == nil {
				assert.Nil(t, id.val)
			} else {
				assert.EqualValues(t, test.expectVal, id.val)
			}
			assert.Equal(t, test.expectedStr, id.String())

		})
	}
}

func TestIDFromDBString(t *testing.T) {
	uid := uuid.Must(uuid.NewRandom())
	tests := []struct {
		name      string
		input     string
		expectErr bool
		expectId  *uuid.UUID
		expectVal *string
	}{
		{
			name:      "no valid key",
			input:     "foobar",
			expectErr: true,
		},
		{
			name:      "empty 2",
			input:     "foobar##",
			expectErr: false,
		},
		{
			name:      "empty",
			input:     "foobar#",
			expectErr: false,
		},
		{
			name:      "happy uuid",
			input:     fmt.Sprintf("FOO#%s", uid.String()),
			expectErr: false,
			expectId:  &uid,
		},
		{
			name:      "happy str",
			input:     fmt.Sprintf("FOO#%s", "username"),
			expectErr: false,
			expectVal: common.StringPtr("username"),
		},
		{
			name:      "happy subid",
			input:     fmt.Sprintf("FOO#%s#sub#subsub", "username"),
			expectErr: false,
			expectVal: common.StringPtr("username"),
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			id, err := IDFromDBString(test.input)
			assert.Equal(t, test.expectErr, err != nil)
			if test.expectErr {
				return
			}
			if test.expectId == nil {
				assert.Nil(t, id.id)
			} else {
				assert.Equal(t, test.expectId.String(), id.id.String())

			}
			if test.expectVal == nil {
				assert.Nil(t, id.val)
			} else {
				assert.EqualValues(t, test.expectVal, id.val)
			}

		})
	}
}

func TestID_IsNil(t *testing.T) {
	uid := uuid.Must(uuid.NewRandom())
	tests := []struct {
		name  string
		input ID
		isNil bool
	}{
		{
			name:  "empty",
			input: ID{},
			isNil: true,
		},
		{
			name: "empty string",
			input: ID{
				val: common.StringPtr(""),
			},
			isNil: true,
		},
		{
			name: "nil UUID",
			input: ID{
				id: &uuid.Nil,
			},
			isNil: true,
		},
		{
			name: "uuid set",
			input: ID{
				id: &uid,
			},
			isNil: false,
		},
		{
			name: "string set",
			input: ID{
				val: common.StringPtr("foo"),
			},
			isNil: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.isNil, test.input.IsNil())
		})
	}
}
