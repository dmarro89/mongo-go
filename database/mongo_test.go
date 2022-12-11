package database

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestNew(t *testing.T) {
	mongo := New()
	assert.Equal(t, mongo != nil, true)
}
