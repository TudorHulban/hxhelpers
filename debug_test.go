package hxhelpers

import "testing"

func TestTraceExit(t *testing.T) {
	defer TraceExit()
}
