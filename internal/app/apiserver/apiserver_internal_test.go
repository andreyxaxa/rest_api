package apiserver

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAPIServer_HandleHello(t *testing.T) {
	s := New(NewConfig())

	rec := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/hello", nil)

	s.handleHello().ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, "HELLO!", rec.Body.String())
}
