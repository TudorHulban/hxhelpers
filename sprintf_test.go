package hxhelpers

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLocal(t *testing.T) {
	cssID := "id"
	elemLabel := "someLabel"

	local := Sprintf(
		`<label for="%s">%s:</label>`,

		cssID,
		elemLabel,
	)

	library := fmt.Sprintf(
		`<label for="%s">%s:</label>`,

		cssID,
		elemLabel,
	)

	require.Equal(t,
		local,
		library,
	)
}

// cpu: AMD Ryzen 7 5800H with Radeon Graphics
// prior to 1.26, version without reflection.
// BenchmarkLocal-16    	15257198	        75.76 ns/op	      32 B/op	       1 allocs/op
// with go 1.26 and unsafe
// BenchmarkLocal-16    	18821125	        63.71 ns/op	      32 B/op	       1 allocs/op
func BenchmarkLocal(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()

	cssID := "id"
	elemLabel := "label"

	for b.Loop() {
		_ = Sprintf(
			`<label for="%s">%s:</label>`,

			cssID,
			elemLabel,
		)
	}
}

// cpu: AMD Ryzen 7 5800H with Radeon Graphics
// prior to go 1.26.
// BenchmarkLibrary-16    	 9989842	       121.0 ns/op	      64 B/op	       3 allocs/op
// with go 1.26
// BenchmarkLibrary-16    	15374064	        78.72 ns/op	      32 B/op	       1 allocs/op
func BenchmarkLibrary(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()

	cssID := "id"
	elemLabel := "label"

	for b.Loop() {
		_ = fmt.Sprintf(
			`<label for="%s">%s:</label>`,

			cssID,
			elemLabel,
		)
	}
}
