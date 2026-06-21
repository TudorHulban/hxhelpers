package piers

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tudorhulban/hxerrors"
)

func TestErrorsValidatePiers(t *testing.T) {
	type Object struct {
		Field any
	}

	object := &Object{}

	require.Error(t, ValidatePiers(object))
}

type structValid struct {
	A *int
}

type structInvalid struct {
	A *int
}

type unexportedStruct struct {
	a *int // unexported, must be ignored
	B *int
}

func TestValidatePiers(t *testing.T) {
	someNumber := 10

	tests := []struct {
		in          any
		want        error
		description string
	}{
		{
			description: "1. Completely untyped nil",
			in:          nil,
			want:        hxerrors.ErrValidation{},
		},
		{
			description: "2. Typed nil pointer",
			in:          (*structValid)(nil),
			want:        hxerrors.ErrValidation{},
		},
		{
			description: "3. Non-struct input - invalid type (int)",
			in:          123,
			want:        hxerrors.ErrValidation{},
		},
		{
			description: "4. Struct with exported nil pointer field",
			in: structInvalid{
				A: nil,
			},
			want: hxerrors.ErrValidation{},
		},
		{
			description: "5. Struct with exported non-nil pointer field",
			in: structValid{
				A: &someNumber,
			},
			want: nil,
		},
		{
			description: "6. Struct with unexported nil pointer field (ignored)",
			in: unexportedStruct{
				a: nil,
				B: &someNumber,
			},
			want: nil,
		},
		{
			description: " 7. Nested pointer: **T - double pointer to valid struct",
			in: func() **structValid {
				s := structValid{
					A: &someNumber,
				}

				p := &s

				return &p
			}(),
			want: nil,
		},
		{
			description: " 8. Nested pointer: **T but inner is nil - double pointer to nil struct",
			in: func() **structValid {
				var p *structValid = nil //nolint:revive

				return &p
			}(),
			want: hxerrors.ErrValidation{},
		},
	}

	for _, tc := range tests {
		t.Run(
			tc.description,
			func(t *testing.T) {
				errValidate := ValidatePiers(tc.in)

				switch {
				case tc.want == nil && errValidate != nil:
					t.Fatalf(
						"unexpected error: %v",
						errValidate,
					)

				case tc.want != nil && errValidate == nil:
					t.Fatal("expected error but got nil")

				case tc.want != nil && errValidate != nil:
					// Compare only the type of the error, not full struct
					if reflect.TypeOf(errValidate) != reflect.TypeOf(tc.want) {
						t.Fatalf(
							"wrong error type: got %T want %T",
							errValidate,
							tc.want,
						)
					}
				}
			},
		)
	}
}
