package utils

import "sync"

type Map[K comparable, T any] struct {
	_map map[K]T
	lock *sync.RWMutex
}

func NewMap[K comparable, T any]() *Map[K, T] {
	return &Map[K, T]{
		_map: make(map[K]T),
		lock: &sync.RWMutex{},
	}
}

func FromMap[K comparable, T any](m map[K]T) *Map[K, T] {
	return &Map[K, T]{
		_map: m,
		lock: &sync.RWMutex{},
	}
}

func (m *Map[K, T]) Get(key K) (T, bool) {
	m.lock.RLock()
	defer m.lock.RUnlock()
	v, ok := m._map[key]
	return v, ok
}

func (m *Map[K, T]) Set(key K, value T) {
	m.lock.Lock()
	defer m.lock.Unlock()
	m._map[key] = value
}

func (m *Map[K, T]) Delete(key K) {
	m.lock.Lock()
	defer m.lock.Unlock()
	delete(m._map, key)
}

func (m *Map[K, T]) Clear() {
	m.lock.Lock()
	defer m.lock.Unlock()
	m._map = make(map[K]T)
}

func (m *Map[K, T]) Len() int {
	m.lock.RLock()
	defer m.lock.RUnlock()
	return len(m._map)
}

// Do not use this to modify the map
func (m *Map[K, T]) ForEach(f func(key K, value T)) {
	m.lock.RLock()
	defer m.lock.RUnlock()
	for k, v := range m._map {
		f(k, v)
	}
}

// Use this if the function modifies the map
func (m *Map[K, T]) ForEachModifySafe(f func(key K, value T)) {
	m.lock.Lock()
	defer m.lock.Unlock()
	for k, v := range m._map {
		m.lock.Unlock()
		f(k, v)
		m.lock.Lock()
	}
}

func (m *Map[K, T]) Find(f func(key K, value T) bool) (*K, *T) {
	m.lock.RLock()
	defer m.lock.RUnlock()
	for k, v := range m._map {
		if f(k, v) {
			return &k, &v
		}
	}
	return nil, nil
}
