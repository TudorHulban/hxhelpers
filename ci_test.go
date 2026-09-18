package hxhelpers

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsRunningInCI(t *testing.T) {
	require.True(t,
		IsRunningInCI(),
	)
}
