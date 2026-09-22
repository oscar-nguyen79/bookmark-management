package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenCode(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name string

		expectedLength int
		expectedError  error
	}{
		{
			name:           "success",
			expectedLength: 12,
			expectedError:  nil,
		},
		{
			name:           "success with custom length",
			expectedLength: 100000,
			expectedError:  nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			testSvc := NewGenCode()
			code, err := testSvc.GenCode(tc.expectedLength)

			assert.ErrorIs(t, err, tc.expectedError)
			assert.Equal(t, tc.expectedLength, len(code))
		})
	}
}
