package daemon

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStop(t *testing.T) {
	startForTest()
	err := Stop()
	require.NoError(t, err)
}
