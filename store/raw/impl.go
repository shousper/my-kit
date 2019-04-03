package raw

import (
	"sync"
)

type DefaultStore struct {
	mu sync.RWMutex
	m map[string][]byte
}

var _ Store = (*DefaultStore)(nil)

func NewDefaultStore() *DefaultStore {
	return &DefaultStore{
		m: make(map[string][]byte),
	}
}

func (s *DefaultStore) Get(key string) ([]byte, error) {
	if key == "" {
		return nil, ErrInvalidKey
	}

	s.mu.RLock()
	defer s.mu.RUnlock()
	if v, ok := s.m[key]; ok {
		return v, nil
	}
	return nil, ErrNotFound
}

func (s *DefaultStore) Set(key string, in []byte) error {
	if key == "" {
		return ErrInvalidKey
	}

	s.mu.Lock()
	defer s.mu.Unlock()
	s.m[key] = in
	return nil
}

func (s *DefaultStore) Delete(key string) error {
	if key == "" {
		return ErrInvalidKey
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.m, key)
	return nil
}

func (s *DefaultStore) Iterate(fn IteratorFunc) error {
	s.mu.RLock()
	defer s.mu.RUnlock()
	for k, v := range s.m {
		done, err := fn(k, v)
		if err != nil {
			return err
		}

		if done {
			return nil
		}
	}
	return nil
}

func (s *DefaultStore) Close() error {
	return nil
}

func (s *DefaultStore) Reset() error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.m = make(map[string][]byte)
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
