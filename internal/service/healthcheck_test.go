package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHealthCheck_Check(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name        string
		serviceName string
		instanceID  string
		expected    HealthCheckResponse
	}{
		{
			name:        "should return health check response",
			serviceName: "bookmark-service",
			instanceID:  "instance-001",
			expected: HealthCheckResponse{
				Message:     "OK",
				ServiceName: "bookmark-service",
				InstanceID:  "instance-001",
			},
		},
		{
			name:        "should return response with empty instance ID",
			serviceName: "bookmark-service",
			instanceID:  "",
			expected: HealthCheckResponse{
				Message:     "OK",
				ServiceName: "bookmark-service",
				InstanceID:  "",
			},
		},
		{
			name:        "should return response with empty service name",
			serviceName: "",
			instanceID:  "instance-001",
			expected: HealthCheckResponse{
				Message:     "OK",
				ServiceName: "",
				InstanceID:  "instance-001",
			},
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			svc := NewHealthCheck(
				tc.serviceName,
				tc.instanceID,
			)
			result, _ := svc.Check()

			assert.Equal(t, tc.expected, result)

		})
	}
}
