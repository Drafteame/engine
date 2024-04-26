package engine

import (
	"context"
	"slices"

	"github.com/Drafteame/engine"
)

// Engine is a generic engine that can be used to run a chain of decorators and a handler.
type Engine[T, R any] struct {
	ctx        context.Context
	evt        T
	handler    engine.Handler[T, R]
	decorators []engine.Decorator[T, R]
}

// New creates a new engine with the given context, event and handler.
func New[T, R any](ctx context.Context, evt T, handler engine.Handler[T, R]) *Engine[T, R] {
	return &Engine[T, R]{
		ctx:     ctx,
		evt:     evt,
		handler: handler,
	}
}

// Use adds a list of decorators to the engine.
func (e *Engine[T, R]) Use(decorators ...engine.Decorator[T, R]) *Engine[T, R] {
	e.decorators = append(e.decorators, decorators...)
	return e
}

// Run executes the engine and returns the result of the handler.
func (e *Engine[T, R]) Run() (R, error) {
	e.applyDecorators()
	return e.handler(e.ctx, e.evt)
}

func (e *Engine[T, R]) applyDecorators() {
	slices.Reverse(e.decorators)

	for _, decorator := range e.decorators {
		e.handler = decorator(e.handler)
	}
}
