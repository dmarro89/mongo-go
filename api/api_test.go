package api

import (
	"mongo-go/service"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestNew(t *testing.T) {
	api := New(service.New(nil))
	assert.Equal(t, api != nil, true)
}
