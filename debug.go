package hxhelpers

import (
	"fmt"
	"runtime"
	"time"
)

// Use as defer helpers.TraceExit().
func TraceExit() {
	pc, _, line, ok := runtime.Caller(1) // Get the caller of this function
	if ok {
		fmt.Printf(
			"exiting function %s at line %d.\n",

			runtime.FuncForPC(pc).Name(),
			line,
		)
	}
}

// Use as defer helpers.TraceExitSince(...).
func TraceExitSince(traceNo int, sectionName string, now time.Time) {
	pc, _, line, couldGetCaller := runtime.Caller(1)
	if couldGetCaller {
		if len(sectionName) > 0 {
			fmt.Printf(
				"trace %d (section %s): exiting function %s at line %d after %d micro-seconds.\n",

				traceNo,
				sectionName,
				runtime.FuncForPC(pc).Name(),
				line,
				time.Since(now).Microseconds(),
			)
		} else {
			fmt.Printf(
				"trace %d: exiting function %s at line %d after %d micro-seconds.\n",

				traceNo,
				runtime.FuncForPC(pc).Name(),
				line,
				time.Since(now).Microseconds(),
			)
		}
	}
}
