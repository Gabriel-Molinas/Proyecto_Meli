package mediator

import (
	"context"
	"fmt"
	"reflect"
)

// Mediator defines the interface for the mediator pattern
type Mediator interface {
	// Send sends a request to the appropriate handler
	Send(ctx context.Context, request interface{}) (interface{}, error)

	// Register registers a handler for a specific request type
	Register(requestType interface{}, handler Handler)
}

// Handler defines the interface for request handlers
type Handler interface {
	Handle(ctx context.Context, request interface{}) (interface{}, error)
}

// mediator is the concrete implementation of the Mediator interface
type mediator struct {
	handlers map[string]Handler
}

// NewMediator creates a new mediator instance
func NewMediator() Mediator {
	return &mediator{
		handlers: make(map[string]Handler),
	}
}

// Send sends a request to the appropriate handler
func (m *mediator) Send(ctx context.Context, request interface{}) (interface{}, error) {
	requestType := reflect.TypeOf(request).String()

	handler, exists := m.handlers[requestType]
	if !exists {
		return nil, fmt.Errorf("no handler registered for request type: %s", requestType)
	}

	return handler.Handle(ctx, request)
}

// Register registers a handler for a specific request type
func (m *mediator) Register(requestType interface{}, handler Handler) {
	typeName := reflect.TypeOf(requestType).String()
	m.handlers[typeName] = handler
}

// HandlerFunc is a function type that implements the Handler interface
type HandlerFunc func(ctx context.Context, request interface{}) (interface{}, error)

// Handle implements the Handler interface for HandlerFunc
func (hf HandlerFunc) Handle(ctx context.Context, request interface{}) (interface{}, error) {
	return hf(ctx, request)
}
