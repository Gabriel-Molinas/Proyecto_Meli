/*
Package mediator implementa el ptrón Mediator para desacoplar controladores REST
de la lógica de negocio en la aplicación.

El patrón Mediator permite que los controladores HTTP envíen queries/comandos sin
conocer directamente los handlers que los procesan, facilitando la mantenibilidad
y testabilidad del código.

Características principales:
- Registro dinámico de handlers por tipo de request
- Resolución automática de handlers basada en reflection
- Interfaz simple para envío de solicitudes
- Soporte para funciones como handlers (HandlerFunc)
*/
package mediator

import (
	"context"
	"fmt"
	"reflect"
)

// Mediator define la interfaz para el patrón mediator
type Mediator interface {
	// Send envía una solicitud al handler apropiado
	Send(ctx context.Context, request interface{}) (interface{}, error)

	// Register registra un handler para un tipo de request específico
	Register(requestType interface{}, handler Handler)
}

// Handler define la interfaz para los handlers de requests
type Handler interface {
	Handle(ctx context.Context, request interface{}) (interface{}, error)
}

// mediator es la implementación concreta de la interfaz Mediator
type mediator struct {
	handlers map[string]Handler
}

// NewMediator crea una nueva instancia de mediator
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

// HandlerFunc es un tipo de función que implementa la interfaz Handler
type HandlerFunc func(ctx context.Context, request interface{}) (interface{}, error)

// Handle implementa la interfaz Handler para HandlerFunc
func (hf HandlerFunc) Handle(ctx context.Context, request interface{}) (interface{}, error) {
	return hf(ctx, request)
}
