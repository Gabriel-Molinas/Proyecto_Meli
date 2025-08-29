package unit

import (
	"context"
	"errors"
	"testing"
	
	"meli-products-api/internal/application/mediator"
)

// Mock structures para testing
type MockRequest struct {
	Value string
}

type MockResponse struct {
	Result string
}

type MockHandler struct {
	response interface{}
	err      error
}

func (m *MockHandler) Handle(ctx context.Context, request interface{}) (interface{}, error) {
	return m.response, m.err
}

type AnotherMockRequest struct {
	ID int
}

func TestNewMediator(t *testing.T) {
	m := mediator.NewMediator()
	
	if m == nil {
		t.Error("NewMediator() returned nil")
	}
}

func TestMediatorRegisterAndSend(t *testing.T) {
	m := mediator.NewMediator()
	
	t.Run("Envío exitoso", func(t *testing.T) {
		expectedResponse := &MockResponse{Result: "success"}
		handler := &MockHandler{
			response: expectedResponse,
			err:      nil,
		}
		
		m.Register(&MockRequest{}, handler)
		
		request := &MockRequest{Value: "test"}
		result, err := m.Send(context.Background(), request)
		
		if err != nil {
			t.Errorf("Send() error = %v, wantErr nil", err)
			return
		}
		
		response, ok := result.(*MockResponse)
		if !ok {
			t.Errorf("Send() result type = %T, want *MockResponse", result)
			return
		}
		
		if response.Result != "success" {
			t.Errorf("Send() result = %v, want success", response.Result)
		}
	})
	
	t.Run("Handler devuelve error", func(t *testing.T) {
		expectedError := errors.New("handler error")
		handler := &MockHandler{
			response: nil,
			err:      expectedError,
		}
		
		m.Register(&AnotherMockRequest{}, handler)
		
		request := &AnotherMockRequest{ID: 1}
		result, err := m.Send(context.Background(), request)
		
		if err == nil {
			t.Error("Send() expected error, got nil")
			return
		}
		
		if err.Error() != "handler error" {
			t.Errorf("Send() error = %v, want handler error", err)
		}
		
		if result != nil {
			t.Errorf("Send() result = %v, want nil", result)
		}
	})
	
	t.Run("Handler no registrado", func(t *testing.T) {
		type UnregisteredRequest struct {
			Data string
		}
		
		request := &UnregisteredRequest{Data: "test"}
		result, err := m.Send(context.Background(), request)
		
		if err == nil {
			t.Error("Send() expected error for unregistered handler, got nil")
			return
		}
		
		expectedError := "no handler registered for request type: unit.UnregisteredRequest"
		if err.Error() != expectedError {
			t.Errorf("Send() error = %v, want %v", err.Error(), expectedError)
		}
		
		if result != nil {
			t.Errorf("Send() result = %v, want nil", result)
		}
	})
}

func TestHandlerFunc(t *testing.T) {
	// Test del tipo HandlerFunc
	handlerFunc := mediator.HandlerFunc(func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(*MockRequest)
		if !ok {
			return nil, errors.New("invalid request type")
		}
		
		return &MockResponse{Result: "from func: " + req.Value}, nil
	})
	
	m := mediator.NewMediator()
	m.Register(&MockRequest{}, handlerFunc)
	
	request := &MockRequest{Value: "test"}
	result, err := m.Send(context.Background(), request)
	
	if err != nil {
		t.Errorf("HandlerFunc error = %v, wantErr nil", err)
		return
	}
	
	response, ok := result.(*MockResponse)
	if !ok {
		t.Errorf("HandlerFunc result type = %T, want *MockResponse", result)
		return
	}
	
	expected := "from func: test"
	if response.Result != expected {
		t.Errorf("HandlerFunc result = %v, want %v", response.Result, expected)
	}
}

func TestMediatorWithContext(t *testing.T) {
	m := mediator.NewMediator()
	
	// Handler que verifica el context
	contextHandler := mediator.HandlerFunc(func(ctx context.Context, request interface{}) (interface{}, error) {
		if ctx == nil {
			return nil, errors.New("context is nil")
		}
		
		// Verificar si el context fue cancelado
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
			return "context ok", nil
		}
	})
	
	m.Register(&MockRequest{}, contextHandler)
	
	t.Run("Context válido", func(t *testing.T) {
		ctx := context.Background()
		request := &MockRequest{Value: "test"}
		
		result, err := m.Send(ctx, request)
		if err != nil {
			t.Errorf("Send() with context error = %v, wantErr nil", err)
			return
		}
		
		if result != "context ok" {
			t.Errorf("Send() with context result = %v, want 'context ok'", result)
		}
	})
	
	t.Run("Context cancelado", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		cancel() // Cancelar inmediatamente
		
		request := &MockRequest{Value: "test"}
		
		result, err := m.Send(ctx, request)
		if err == nil {
			t.Error("Send() with cancelled context expected error, got nil")
			return
		}
		
		if err != context.Canceled {
			t.Errorf("Send() with cancelled context error = %v, want %v", err, context.Canceled)
		}
		
		if result != nil {
			t.Errorf("Send() with cancelled context result = %v, want nil", result)
		}
	})
}