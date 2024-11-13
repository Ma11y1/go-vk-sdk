package events

type EventListener[T interface{}] struct {
	Listener func(T)
}

func NewEventListener[T interface{}](listener func(T)) *EventListener[T] {
	return &EventListener[T]{listener}
}

type EventEmitter[T, U comparable] struct {
	listeners map[T][]*EventListener[U]
}

func NewEventEmitter[T, U comparable]() *EventEmitter[T, U] {
	return &EventEmitter[T, U]{
		listeners: make(map[T][]*EventListener[U]),
	}
}

func (e *EventEmitter[T, U]) On(name T, handler *EventListener[U]) {
	e.listeners[name] = append(e.listeners[name], handler)
}

func (e *EventEmitter[T, U]) Off(name T, handler *EventListener[U]) {
	if _, ok := e.listeners[name]; !ok {
		return
	}

	handlers := e.listeners[name]
	for i := 0; i < len(handlers); i++ {
		if handlers[i] == handler {
			handlers = append(handlers[:i], handlers[i+1:]...)
			e.listeners[name] = handlers
			return
		}
	}
}

func (e *EventEmitter[T, U]) Emit(name T, event U) {
	if handlers, ok := e.listeners[name]; ok {
		for _, handler := range handlers {
			handler.Listener(event)
		}
	}
}

func (e *EventEmitter[T, U]) Clear(name T) {
	delete(e.listeners, name)
}

func (e *EventEmitter[T, U]) Keys() []T {
	keys := make([]T, 0, len(e.listeners))
	for k := range e.listeners {
		keys = append(keys, k)
	}
	return keys
}
