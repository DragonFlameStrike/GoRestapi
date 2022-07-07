package apiserver

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAPIServer_FirstBanner(t *testing.T) {
	s := New(NewConfig())
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/root/api/1", nil)
	s.firstBanner().ServeHTTP(rec, req)
	assert.Equal(t, rec.Body.String(), "{\"name\":\"Banner1\",\"price\":100,\"text\":\"Test\",\"deleted\":false,\"idBanner\":1}")
}
