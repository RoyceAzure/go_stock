package util

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRandomDecimal(t *testing.T) {
	for i := 0; i < 100; i++ {
		_, err := RandomDecimal(6, 2)
		require.NoError(t, err)
	}
}

func TestRandomNumeric(t *testing.T) {
	for i := 0; i < 100; i++ {
		_, err := RandomNumeric(6, 2)
		require.NoError(t, err)
	}
}
