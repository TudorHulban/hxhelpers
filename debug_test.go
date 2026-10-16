package hxhelpers

import (
	"testing"
	"time"
)

func TestTraceExit(t *testing.T) {
	defer TraceExit()
}

func TestTraceExitTime(t *testing.T) {
	defer TraceExitTime(time.Now())
}
