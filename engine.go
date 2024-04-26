package engine

import (
	"context"
	"os"
	"slices"

	"github.com/aws/aws-lambda-go/lambda"
)

// Handler is a function that is complaint to with the golang aws-lambda-go sdk.
// Where "T" and "R" are types compatible with the "encoding/json" standard library.
// See https://golang.org/pkg/encoding/json/#Unmarshal for how deserialization behaves
type Handler[T, R any] func(context.Context, T) (R, error)

// Decorator is a function that receives a handler and returns a new handler wrapping the original one to add
// functionality.
type Decorator[T, R any] func(handler Handler[T, R]) Handler[T, R]

// Engine is a struct that holds the handler and the decorators to be applied to the handler.
type Engine[T, R any] struct {
	handler    Handler[T, R]
	decorators []Decorator[T, R]
}

// New creates a new Engine with the given handler.
func New[T, R any](handler Handler[T, R]) *Engine[T, R] {
	return &Engine[T, R]{
		handler: handler,
	}
}

// Use adds decorators to the engine.
func (e *Engine[T, R]) Use(decorators ...Decorator[T, R]) *Engine[T, R] {
	e.decorators = append(e.decorators, decorators...)
	return e
}

// Run starts the engine.
func (e *Engine[T, R]) Run() {
	e.applyDecorators()

	if os.Getenv("AWS_LAMBDA_RUNTIME_API") != "" {
		lambda.Start(e.handler)
		return
	}

	panic("engine: no runtime detected")
}

func (e *Engine[T, R]) applyDecorators() {
	slices.Reverse(e.decorators)

	for _, decorator := range e.decorators {
		e.handler = decorator(e.handler)
	}
}
