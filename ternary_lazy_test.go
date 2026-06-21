package hxhelpers

import "testing"

func TestTernaryLazy(t *testing.T) {
	// A helper function that panics if called, used to prove lazy evaluation.
	panicFunc := func() string {
		panic("this should not be called!")
	}

	t.Run(
		"1. True Condition - Evaluates Only Value1",
		func(t *testing.T) {
			value1Called := false
			v1 := func() string {
				value1Called = true

				return "success"
			}

			// Since condition is true, panicFunc (value2) should never run
			result := TernaryLazy(true, v1, panicFunc)

			if result != "success" {
				t.Errorf(
					"expected 'success', got %q",
					result,
				)
			}

			if !value1Called {
				t.Error(
					"expected value1 to be executed, but it was not",
				)
			}
		},
	)

	t.Run(
		"2. False Condition - Evaluates Only Value2",
		func(t *testing.T) {
			value2Called := false
			v2 := func() string {
				value2Called = true

				return "fallback"
			}

			// Since condition is false, panicFunc (value1) should never run
			result := TernaryLazy(false, panicFunc, v2)

			if result != "fallback" {
				t.Errorf("expected 'fallback', got %q", result)
			}

			if !value2Called {
				t.Error("expected value2 to be executed, but it wasn't")
			}
		},
	)
}
