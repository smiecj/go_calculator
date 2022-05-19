package calculator

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCompare(t *testing.T) {
	ret, err := Compare("2 >= 1")
	require.Nil(t, err)
	require.True(t, ret)

	ret, err = Compare("1 < 100")
	require.Nil(t, err)
	require.True(t, ret)

	ret, err = Compare("-10 > -20")
	require.Nil(t, err)
	require.True(t, ret)
}
