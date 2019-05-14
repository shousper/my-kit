package local

import (
	"sync"

	"github.com/shousper/my-kit/store/raw"
)

type DefaultStore struct {
	mu sync.RWMutex
	m  map[string]raw.Value
}

var _ raw.Store = (*DefaultStore)(nil)

func NewDefaultStore() *DefaultStore {
	return &DefaultStore{
		m: make(map[string]raw.Value),
	}
}

func (s *DefaultStore) Get(key string) (raw.Value, error) {
	if key == "" {
		return nil, raw.ErrInvalidKey
	}

	s.mu.RLock()
	defer s.mu.RUnlock()
	if v, ok := s.m[key]; ok {
		return v, nil
	}
	return nil, raw.ErrNotFound
}

func (s *DefaultStore) Set(key string, in raw.Value) error {
	if key == "" {
		return raw.ErrInvalidKey
	}
	if in == nil {
		return raw.ErrInvalidValue
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	s.m[key] = in
	return nil
}

func (s *DefaultStore) Delete(key string) error {
	if key == "" {
		return raw.ErrInvalidKey
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.m, key)
	return nil
}

func (s *DefaultStore) Keys() ([]string, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	out := make([]string, len(s.m))
	i := 0
	for key := range s.m {
		out[i] = key
		i++
	}
	return out, nil
}

func (s *DefaultStore) Iterate() raw.Iterator {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return &raw.DefaultIterator{
		Store: s,
	}
}

func (s *DefaultStore) Close() error {
	return nil
}

func (s *DefaultStore) Reset() error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.m = make(map[string]raw.Value)
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
