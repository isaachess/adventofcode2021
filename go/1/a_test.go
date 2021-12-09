package one

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestA(t *testing.T) {
	numIncreases, err := A("input.txt")
	require.NoError(t, err)
	fmt.Println(numIncreases)
	require.Equal(t, 1832, numIncreases)
}

func TestB(t *testing.T) {
	numIncreases, err := B("input.txt")
	require.NoError(t, err)
	fmt.Println(numIncreases)
	require.Equal(t, 1858, numIncreases)
}
