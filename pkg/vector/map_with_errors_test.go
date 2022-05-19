package vector

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestAllSucceeds(t *testing.T) {
	vec := FromValues(1, 2, 3)
	xs, errs := MapWithErrors(vec, func(x int) (int, error) { return x, nil })
	require.NoError(t, AggregateError(errs, ""))
	require.Equal(t, 3, xs.Size())
}

func TestAllFailed(t *testing.T) {
	vec := FromValues(1, 2, 3)
	xs, errs := MapWithErrors(vec, func(int) (int, error) { return 0, fmt.Errorf("fail") })
	require.Error(t, AggregateError(errs, ""))
	require.Equal(t, 0, xs.Size())
}

func TestPartialFailure(t *testing.T) {
	vec := FromValues(1, 2, 3)
	xs, errs := MapWithErrors(vec, func(x int) (int, error) {
		if x == 1 {
			return x, nil
		}
		return 0, fmt.Errorf("fail")
	})
	require.Error(t, AggregateError(errs, ""))
	require.Equal(t, "fail, fail", AggregateError(errs, ", ").Error())
	require.Equal(t, 1, xs.Size())
}
