package generic

import (
	"reflect"
	"sync"

	"github.com/shousper/my-kit/store/object"
)

type DefaultStore struct {
	mu sync.RWMutex
	m  map[string]object.Value
}

var _ object.Store = (*DefaultStore)(nil)

func NewStore() *DefaultStore {
	return &DefaultStore{
		m: make(map[string]object.Value),
	}
}

func (s *DefaultStore) Get(key string, out object.Value) error {
	if key == "" {
		return object.ErrInvalidKey
	}

	s.mu.RLock()
	defer s.mu.RUnlock()
	if v, ok := s.m[key]; ok {
		// While reflect is a little "dirty", this is what the encoding package does to facilitate
		// type safety for consumers.
		outV := reflect.ValueOf(out)
		newV := reflect.ValueOf(v)
		if outV.Kind() == reflect.Ptr {
			outV = outV.Elem()
			if newV.Kind() == reflect.Ptr {
				newV = newV.Elem()
			}
		}
		outV.Set(newV)
		return nil
	}

	return object.ErrNotFound
}

func (s *DefaultStore) Set(key string, in object.Value) error {
	if key == "" {
		return object.ErrInvalidKey
	}

	s.mu.Lock()
	defer s.mu.Unlock()
	s.m[key] = in
	return nil
}

func (s *DefaultStore) Delete(key string) error {
	if key == "" {
		return object.ErrInvalidKey
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.m, key)
	return nil
}

func (s *DefaultStore) Keys() ([]string, error) {
	out := make([]string, len(s.m))
	i := 0
	for k := range s.m {
		out[i] = k
		i++
	}
	return out, nil
}

func (s *DefaultStore) Iterate() object.Iterator {
	return object.NewIterator(s)
}

func (s *DefaultStore) Close() error {
	return nil
}

func (s *DefaultStore) Reset() error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.m = make(map[string]object.Value)
	return nil
}

func (s *DefaultStore) Len() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return len(s.m)
}

func (s *DefaultStore) Capacity() int {
	// Doesn't really work for maps.. so we'll just return Len?
	return s.Len()
}
