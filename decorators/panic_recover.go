package decorators

import (
	"context"
	"fmt"

	"github.com/Drafteame/engine"
)

// PanicRecoverConfig is the configuration for the PanicRecover decorator.
type PanicRecoverConfig[T any] struct {
	LogFunc func(context.Context, T, error)
}

// DefaultPanicRecoverConfig returns the default configuration for the PanicRecover decorator.
func DefaultPanicRecoverConfig[T any]() PanicRecoverConfig[T] {
	return PanicRecoverConfig[T]{
		LogFunc: nil,
	}
}

// PanicRecover is a decorator that recovers from panic and returns an error.
func PanicRecover[T, R any]() engine.Decorator[T, R] {
	return PanicRecoverWithConfig[T, R](DefaultPanicRecoverConfig[T]())
}

// PanicRecoverWithConfig is a decorator that recovers from panic and returns an error with a custom configuration.
func PanicRecoverWithConfig[T, R any](config PanicRecoverConfig[T]) engine.Decorator[T, R] {
	return func(handler engine.Handler[T, R]) engine.Handler[T, R] {
		return func(ctx context.Context, request T) (response R, err error) {
			defer func() {
				if r := recover(); r != nil {
					err = fmt.Errorf("panic: %v", r)

					if config.LogFunc != nil {
						config.LogFunc(ctx, request, err)
					}
				}
			}()

			return handler(ctx, request)
		}
	}
}
