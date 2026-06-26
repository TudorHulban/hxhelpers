package hxhelpers

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNotInSliceSource(t *testing.T) {
	tests := []struct {
		description string
		source      []int
		elements    []int
		want        []int
	}{
		{
			description: "1. Elements exist in source - should return empty slice",
			source:      []int{1, 2, 3, 4, 5},
			elements:    []int{2, 4},
			want:        []int{},
		},
		{
			description: "2. Elements do not exist in source - should return all elements",
			source:      []int{1, 2, 3},
			elements:    []int{4, 5, 6},
			want:        []int{4, 5, 6},
		},
		{
			description: "3. Mixed elements - should return only the missing ones",
			source:      []int{1, 2, 3},
			elements:    []int{2, 3, 4, 5},
			want:        []int{4, 5},
		},
		{
			description: "4. Empty source - should return all elements unmodified",
			source:      []int{},
			elements:    []int{1, 2, 3},
			want:        []int{1, 2, 3},
		},
		{
			description: "5. Empty elements - should return nil or empty slice",
			source:      []int{1, 2, 3},
			elements:    []int{},
			want:        nil,
		},
		{
			description: "6. Duplicate elements in input - should preserve duplicates if missing from source",
			source:      []int{1, 2},
			elements:    []int{3, 3, 4},
			want:        []int{3, 3, 4},
		},
	}

	for _, tc := range tests {
		t.Run(
			tc.description,
			func(t *testing.T) {
				got := NotInSliceSource(tc.source, tc.elements...)

				// Using reflect.DeepEqual because order and contents matter for slices.
				// Note: If both are empty/nil, DeepEqual might care about nil vs []int{},
				// so we handle the length 0 edge case gracefully.
				if len(got) == 0 && len(tc.want) == 0 {
					return
				}

				require.True(t,
					reflect.DeepEqual(got, tc.want),

					"NotInSliceSource() = %v, want %v",
					got,
					tc.want,
				)
			},
		)
	}
}
