package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWorkerUtils_SwitchTo(t *testing.T) {
	testCases := []struct {
		name string
		expectedBool bool
		flag string
		counter int64
	} {
		{
			"should resolve true on odd flag",
			true,
			"odd",
			1,
		},
		{
			"should resolve false on odd flag",
			false,
			"odd",
			0,
		},
		{
			"should resolve true on even flag",
			true,
			"even",
			0,
		},
		{
			"should resolve false on even flag",
			false,
			"even",
			1,
		},
		{
			"should resolve true on any other flag",
			true,
			"all",
			2,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			response := SwitchTo(tc.flag, tc.counter)

			assert.Equal(t, tc.expectedBool, response)
		})
	}
}
