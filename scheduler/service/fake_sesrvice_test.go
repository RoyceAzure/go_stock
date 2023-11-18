package service

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGenerateFakeData(t *testing.T) {
	fakedataService := FakeSPRDataService{
		dao: testDao,
	}
	prototype, err := fakedataService.GeneratePrototype()

	require.NoError(t, err)
	require.NotEmpty(t, prototype)

	fakedataService.SetPrototype(prototype)
	fakesprs, err := fakedataService.GenerateFakeData()
	require.NoError(t, err)
	require.NotEmpty(t, fakesprs)
}
