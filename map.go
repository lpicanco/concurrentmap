package concurrentmap

import "sync"

// Map is a thread-safe map
type Map struct {
	items map[interface{}]interface{}
	mu    sync.RWMutex
}

// New returns a new thread-safe map
func New() Map {
	return Map{items: make(map[interface{}]interface{})}
}

// Get returns the value to which the specified key is mapped.
func (m *Map) Get(key interface{}) (value interface{}, found bool) {
	m.mu.RLock()
	value, found = m.items[key]
	m.mu.RUnlock()

	return
}

// Put associates the specified value with the specified key.
func (m *Map) Put(key interface{}, value interface{}) {
	m.mu.Lock()
	m.items[key] = value
	m.mu.Unlock()
}

// Remove the entry associated with the specified key.
func (m *Map) Remove(key interface{}) (found bool) {
	_, found = m.Get(key)

	if found {
		m.mu.Lock()
		delete(m.items, key)
		defer m.mu.Unlock()
	}
	return
}

// Size returns the number of items in this map
func (m *Map) Size() (size int) {
	m.mu.RLock()
	size = len(m.items)
	m.mu.RUnlock()
	return
}
