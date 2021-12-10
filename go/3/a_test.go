package three

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestA(t *testing.T) {
	power, err := A("../../inputs/3/input.txt")
	require.NoError(t, err)
	require.Equal(t, int64(4001724), power)
}
