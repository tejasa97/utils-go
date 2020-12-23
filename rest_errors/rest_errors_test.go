package rest_errors

import (
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewInternalServerError(t *testing.T) {
	err := NewInternalServerError("this is the message", errors.New("DB error"))
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "internal_server_error", err.Error)
	assert.EqualValues(t, "this is the message", err.Message)

	assert.NotNil(t, err.Causes)
	assert.EqualValues(t, 1, len(err.Causes))
	assert.EqualValues(t, "DB error", err.Causes[0])
}

func TestNewBadRequestError(t *testing.T) {
	err := NewBadRequestError("this is the message")
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusBadRequest, err.Status)
	assert.EqualValues(t, "bad_request", err.Error)
	assert.EqualValues(t, "this is the message", err.Message)
}

func TestNewNotFoundError(t *testing.T) {
	err := NewNotFoundError("this is the message")
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.Status)
	assert.EqualValues(t, "not_found", err.Error)
	assert.EqualValues(t, "this is the message", err.Message)
}

func TestNewNotAuthorizedError(t *testing.T) {
	err := NewUnauthorizedError("this is the message")
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusUnauthorized, err.Status)
	assert.EqualValues(t, "unauthorized", err.Error)
	assert.EqualValues(t, "this is the message", err.Message)
}
