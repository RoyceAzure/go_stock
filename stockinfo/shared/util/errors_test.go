package util

import (
	"errors"
	"fmt"
	"testing"

	"github.com/RoyceAzure/go-stockinfo/shared/util/constants"
	"github.com/stretchr/testify/require"
)

func TestFmtErrorf(t *testing.T) {
	err := errors.New("i am test err")
	err2 := fmt.Errorf("warp err %w", err)
	err3 := fmt.Errorf("warp err %w", err2)
	require.True(t, errors.Is(err2, err))
	require.True(t, errors.Is(err3, err))
	require.True(t, errors.Is(err3, err2))
	require.False(t, errors.Is(err2, err3))
}

func TestErrJoin(t *testing.T) {
	err := errors.New("i am test err")
	err2 := errors.New("i am test err2")
	err3 := errors.New("i am test err3")
	errTotal := errors.Join(err, err2, err3)
	require.True(t, errors.Is(errTotal, err))
	require.True(t, errors.Is(errTotal, err2))
	require.True(t, errors.Is(errTotal, err3))
}

func TestCustomNew(t *testing.T) {
	err := constants.ErrInValidatePreConditionOp
	err2 := fmt.Errorf("err2 %w", err)
	err3 := fmt.Errorf("err3 %w", err2)
	require.True(t, errors.Is(err3, err2))
	require.True(t, errors.Is(err3, constants.ErrInValidatePreConditionOp))
	require.True(t, errors.Is(err2, constants.ErrInValidatePreConditionOp))
}
