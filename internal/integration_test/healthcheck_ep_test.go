package integration_test

import (
	"bookmark-management/internal/api"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHealthCheckEndPoint(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name string

		setupTestHTTP func(api api.Engine) *httptest.ResponseRecorder

		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name: "normal case",
			setupTestHTTP: func(api api.Engine) *httptest.ResponseRecorder {
				req := httptest.NewRequest(http.MethodGet, "/health-check", nil)
				respRecorder := httptest.NewRecorder()

				api.ServeHTTP(respRecorder, req)
				return respRecorder
			},
			expectedStatusCode:   http.StatusOK,
			expectedResponseBody: `{"message":"OK","service_name":"Bookmark-Management","instance_id":"test-instance-xxx"}`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			testAPI := api.NewEngine(&api.Config{
				ServiceName: "Bookmark-Management",
				InstanceID:  "test-instance-xxx",
			})
			recorder := tc.setupTestHTTP(testAPI)

			assert.Equal(t, tc.expectedStatusCode, recorder.Code)
			assert.Containsf(t, recorder.Body.String(), tc.expectedResponseBody, `{"message":"OK","service_name":"Bookmark-Management","instance_id":"test-instance-xxx"}`)
		})
	}
}
