package apiserver

import (
	"github.com/mmaxim2710/firstWebApp/internal/app/config"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAPIServer_HandleHello(t *testing.T) {
	newCfg, err := config.NewConfig("../../../configs/apiserver_test.yaml")
	if err != nil {
		t.Fatal(err)
	}

	s := New(newCfg)
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/hello", nil)
	s.handleHello().ServeHTTP(rec, req)
	assert.Equal(t, rec.Body.String(), "hello")
}
