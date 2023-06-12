package utils

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBinary(t *testing.T) {
	tcs := map[string]struct {
		input    []int
		target   int
		expIndx  int
		expItera int
		expErr   error
	}{
		"5 index": {
			input:    []int{1, 2, 3, 4, 5},
			target:   5,
			expIndx:  4,
			expItera: 2,
			expErr:   nil,
		},
		"not found": {
			input:    []int{1, 2, 3, 4, 5},
			target:   9,
			expIndx:  0,
			expItera: 3,
			expErr:   errors.New("not found"),
		},
	}
	for name, tc := range tcs {
		t.Run(name, func(t *testing.T) {
			index, iteration, err := BinarySearch(tc.input, 0, len(tc.input), tc.target, 0)

			if tc.expErr != nil {
				require.Equal(t, tc.expErr.Error(), err.Error())
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.expIndx, index)
				require.Equal(t, tc.expItera, iteration)
			}
		})
	}
}
