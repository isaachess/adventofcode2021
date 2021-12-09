package one

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestA(t *testing.T) {
	numIncreases, err := A("../../inputs/1/input.txt")
	require.NoError(t, err)
	require.Equal(t, 1832, numIncreases)
}

func TestB(t *testing.T) {
	numIncreases, err := B("../../inputs/1/input.txt")
	require.NoError(t, err)
	require.Equal(t, 1858, numIncreases)
}
