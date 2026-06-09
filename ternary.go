package hxhelpers

// Ternary returns first value if condition true.
// Second if condition false.
//
// Go eagerly evaluates arguments therefore arguments should be valid.
// Use the lazy version if unsure.
func Ternary[T any](condition bool, value1, value2 T) T {
	if condition {
		return value1
	}

	return value2
}

// TernaryLazy evaluates a boolean condition and returns the result of value1() if true,
// or value2() if false.
//
// Unlike the standard Ternary function, it accepts functions as arguments to defer
// evaluation, making it ideal for performance-heavy operations or preventing panics
// from unchosen branches.
func TernaryLazy[T any](condition bool, value1, value2 func() T) T {
	if condition {
		return value1()
	}

	return value2()
}
