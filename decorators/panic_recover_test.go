package decorators

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	testengine "github.com/Drafteame/engine/test/engine"
)

func TestPanicRecover(t *testing.T) {
	t.Run("should recover from panic and return error", func(t *testing.T) {
		handler := func(context.Context, string) (string, error) {
			panic("something went wrong")
		}

		res, err := testengine.New(context.Background(), "hello", handler).
			Use(PanicRecover[string, string]()).
			Run()

		assert.Empty(t, res)
		assert.Error(t, err)
		assert.Equal(t, "panic: something went wrong", err.Error())
	})

	t.Run("should recover from panic and log error", func(t *testing.T) {
		var loggedError error
		logFunc := func(_ context.Context, _ string, err error) {
			loggedError = err
		}

		handler := func(context.Context, string) (string, error) {
			panic("something went wrong")
		}

		_, _ = testengine.New(context.Background(), "hello", handler).
			Use(PanicRecoverWithConfig[string, string](PanicRecoverConfig[string]{
				LogFunc: logFunc,
			})).
			Run()

		assert.Equal(t, "panic: something went wrong", loggedError.Error())
	})
}
