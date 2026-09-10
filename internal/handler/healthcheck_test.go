package handler

import (
	"bookmark-management/internal/service"
	"bookmark-management/internal/service/mocks"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestHealthCheckHandler_Check(t *testing.T) {
	t.Parallel()

	testsCases := []struct {
		name string

		setupRequest     func(ctx *gin.Context)
		setupMockService func(ctx *gin.Context) *mocks.HealthCheck

		expectedStatus   int
		expectedResponse string
	}{
		{
			name: "success",
			setupRequest: func(ctx *gin.Context) {
				ctx.Request = httptest.NewRequest(http.MethodGet, "/health-check", nil)
			},
			setupMockService: func(ctx *gin.Context) *mocks.HealthCheck {
				serviceMock := mocks.NewHealthCheck(t)
				serviceMock.On("Check").Return(service.HealthCheckResponse{
					Message:     "OK",
					ServiceName: "Bookmark-Management",
					InstanceID:  "test-instance-xxx",
				}, nil,
				)
				return serviceMock
			},
			expectedStatus:   http.StatusOK,
			expectedResponse: `{"message":"OK","service_name":"Bookmark-Management","instance_id":"test-instance-xxx"}`,
		},
	}

	for _, tc := range testsCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			rec := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(rec)
			tc.setupRequest(ctx)

			mockSvc := tc.setupMockService(ctx)
			testHandler := NewHealthCheck(mockSvc)

			testHandler.Check(ctx)

			assert.Equal(t, tc.expectedStatus, rec.Code)
			assert.Equal(t, tc.expectedResponse, rec.Body.String())
		})
	}

}
