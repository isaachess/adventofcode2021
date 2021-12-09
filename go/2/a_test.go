package two

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestA(t *testing.T) {
	num, err := A("../../inputs/2/input.txt")
	require.NoError(t, err)
	require.Equal(t, 2272262, num)
}
func TestB(t *testing.T) {
	num, err := B("../../inputs/2/input.txt")
	require.NoError(t, err)
	require.Equal(t, 2134882034, num)
}
