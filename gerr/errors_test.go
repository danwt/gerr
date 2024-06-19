package gerr

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/status"
)

var err error = ErrCancelled // arbitrary choice

func TestCompatibleWithGoogleStatusLib(t *testing.T) {
	t.Run("FromError", func(t *testing.T) {
		_, ok := status.FromError(err)
		require.True(t, ok)
	})
}

func TestBasics(t *testing.T) {
	t.Run("text is as expected", func(t *testing.T) {
		errA := ErrCancelled
		errB := fmt.Errorf("foobar: %w", errA)
		require.Equal(t, "foobar: cancelled", errB.Error())
	})
	t.Run("as", func(t *testing.T) {
		errA := fmt.Errorf("foo: %w", err)
		errB := fmt.Errorf("bar: %w", errA)

		var e Error
		require.ErrorAs(t, errB, &e)
	})
	t.Run("is", func(t *testing.T) {
		errA := fmt.Errorf("foo: %w", err)
		errB := fmt.Errorf("bar: %w", errA)

		require.ErrorIs(t, errB, err)
	})
}
