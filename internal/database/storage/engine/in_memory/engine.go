package in_memory

import "sync"

type Engine struct {
	mu    sync.RWMutex
	store map[string]string
}

func NewEngine() *Engine {
	return &Engine{
		store: make(map[string]string),
	}
}

func (e *Engine) Set(key string, value string) {
	e.mu.Lock()
	defer e.mu.Unlock()

	e.store[key] = value
}

func (e *Engine) Get(key string) (string, bool) {
	e.mu.RLock()
	defer e.mu.Unlock()

	value, ok := e.store[key]

	return value, ok
}

func (e *Engine) Del(key string) {
	e.mu.Lock()
	defer e.mu.Unlock()

	delete(e.store, key)
}
