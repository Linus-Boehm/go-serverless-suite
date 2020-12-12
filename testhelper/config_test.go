package testhelper

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "happy",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c, err := LoadConfig()
			assert.NoError(t, err)
			assert.NotNil(t, c)
		})
	}
}

