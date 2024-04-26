package decorators

import (
	"context"
	"log/slog"

	"github.com/Drafteame/engine"
)

type LogEventConfig[T, R any] struct {
	LogFunc func(context.Context, T, R, error)
}

func DefaultLogEventLogFunc[T, R any](ctx context.Context, evt T, res R, err error) {
	log := slog.Default()

	if err != nil {
		log.ErrorContext(ctx, "error occurred", "error", err, "event", evt, "response", res)
		return
	}

	log.InfoContext(ctx, "event processed", "event", evt, "response", res)
}

func DefaultLogEventConfig[T, R any]() LogEventConfig[T, R] {
	return LogEventConfig[T, R]{LogFunc: DefaultLogEventLogFunc[T, R]}
}

func LogEventWithConfig[T, R any](config LogEventConfig[T, R]) engine.Decorator[T, R] {
	return func(handler engine.Handler[T, R]) engine.Handler[T, R] {
		return func(ctx context.Context, evt T) (R, error) {
			res, err := handler(ctx, evt)

			logFunc := DefaultLogEventLogFunc[T, R]
			if config.LogFunc != nil {
				logFunc = config.LogFunc
			}

			logFunc(ctx, evt, res, err)

			return res, err
		}
	}
}

func LogEvent[T, R any]() engine.Decorator[T, R] {
	return LogEventWithConfig(DefaultLogEventConfig[T, R]())
}
