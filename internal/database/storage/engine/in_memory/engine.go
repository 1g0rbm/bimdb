package in_memory

type Engine struct {
	store map[string]string
}

func NewEngine() *Engine {
	return &Engine{
		store: make(map[string]string),
	}
}

func (e *Engine) Set(key string, value string) {
	e.store[key] = value
}

func (e *Engine) Get(key string) (string, bool) {
	value, ok := e.store[key]

	return value, ok
}

func (e *Engine) Del(key string) {
	delete(e.store, key)
}
