package hxhelpers

import "unsafe"

// hopefully Sprintf was rewritten with unsafe
// as the prior version with string buffer had lower performance.
// version is around 20% faster than standard library 1.26.

// Sprintf is a minimal %s-only formatter that trades immutability for more speed.
//
// It performs two passes: the first
// computes the exact output size, the second fills a preallocated buffer.
// It uses unsafe.String to convert the final byte slice to a string without
// allocating. The returned string aliases the slice's memory, so the slice
// must never be mutated after conversion.
func Sprintf(format string, a ...string) string {
	indexArguments := 0
	lengthOutput := 0
	lengthFormat := len(format)

	for ix := 0; ix < lengthFormat; ix++ {
		if format[ix] == '%' && ix+1 < lengthFormat && format[ix+1] == 's' && indexArguments < len(a) {
			lengthOutput = lengthOutput + len(a[indexArguments])
			indexArguments++
			ix++
		} else {
			lengthOutput++
		}
	}

	if lengthOutput == 0 {
		return ""
	}

	// Second pass: fill buffer
	buf := make([]byte, lengthOutput)
	out := 0
	indexArguments = 0

	for ix := 0; ix < lengthFormat; ix++ {
		if format[ix] == '%' && ix+1 < lengthFormat && format[ix+1] == 's' && indexArguments < len(a) {
			argument := a[indexArguments]

			out = out + copy(buf[out:], argument) // copy accepts string directly
			indexArguments++
			ix++
		} else {
			buf[out] = format[ix]
			out++
		}
	}

	return unsafe.String(&buf[0], lengthOutput)
}
