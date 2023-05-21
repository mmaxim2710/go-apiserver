package apiserver

import (
	"bytes"
	"encoding/json"
	"github.com/mmaxim2710/firstWebApp/internal/app/store/teststore"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServer_HandleUsersCreate(t *testing.T) {
	s := newServer(teststore.New())

	testCases := []struct {
		name         string
		payload      interface{}
		expectedCode int
	}{
		{
			name: "valid",
			payload: map[string]string{
				"email":    "user@example.org",
				"password": "password",
			},
			expectedCode: http.StatusCreated,
		},

		{
			name: "invalid email",
			payload: map[string]string{
				"email":    "invalid",
				"password": "password",
			},
			expectedCode: http.StatusUnprocessableEntity,
		},

		{
			name: "invalid password",
			payload: map[string]string{
				"email":    "invalid",
				"password": "passwordpasswordpasswordpasswordpasswordpasswordpasswordpasswordpasswordpassword",
			},
			expectedCode: http.StatusUnprocessableEntity,
		},

		{
			name: "empty email and password",
			payload: map[string]string{
				"email":    "",
				"password": "",
			},
			expectedCode: http.StatusUnprocessableEntity,
		},

		{
			name:         "invalid payload",
			payload:      "invalid",
			expectedCode: http.StatusBadRequest,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			buf := &bytes.Buffer{}
			err := json.NewEncoder(buf).Encode(tc.payload)
			if err != nil {
				t.Fatal(err)
			}
			req, _ := http.NewRequest(http.MethodPost, "/users", buf)
			s.ServeHTTP(rec, req)
			assert.Equal(t, tc.expectedCode, rec.Code)
		})
	}

	//rec := httptest.NewRecorder()
	//
	//
	//
}
