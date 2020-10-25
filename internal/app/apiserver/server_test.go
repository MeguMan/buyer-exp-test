package apiserver_test

import (
	"bytes"
	"encoding/json"
	"github.com/MeguMan/buyer-exp-test/internal/app/apiserver"
	"github.com/MeguMan/buyer-exp-test/internal/app/emailsender"
	"github.com/MeguMan/buyer-exp-test/internal/app/store/teststore"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServer_Follow(t *testing.T) {
	s := apiserver.NewServer(teststore.New(), emailsender.Sender{
		Email:    "",
		Password: "",
		TLSPort:  "",
	})

	testCases := []struct {
		name         string
		d            apiserver.Data
		expectedCode int
	}{
		{
			name: "valid",
			d: apiserver.Data{
				Email: "sdpihaoup@gmail.com",
				Link:  "avito.ru",
			},
			expectedCode: http.StatusCreated,
		},
		{
			name: "invalid email",
			d: apiserver.Data{
				Email: "sdpihaoup@gmail",
				Link:  "avito.ru",
			},
			expectedCode: http.StatusBadRequest,
		},
		{
			name: "invalid link",
			d: apiserver.Data{
				Email: "sdpihaoup@gmail.com",
				Link:  "avito",
			},
			expectedCode: http.StatusBadRequest,
		},
		{
			name: "empty email",
			d: apiserver.Data{
				Email: "",
				Link:  "avito.ru",
			},
			expectedCode: http.StatusBadRequest,
		},
		{
			name: "empty link",
			d: apiserver.Data{
				Email: "sdpihaoup@gmail.com",
				Link:  "",
			},
			expectedCode: http.StatusBadRequest,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			requestBody, err := json.Marshal(tc.d)
			if err != nil {
				panic(err)
			}
			rec := httptest.NewRecorder()
			req, _ := http.NewRequest("POST", "/follow", bytes.NewBuffer(requestBody))
			defer req.Body.Close()
			s.ServeHTTP(rec, req)
			assert.Equal(t, tc.expectedCode, rec.Code)
		})
	}
}
