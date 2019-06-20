package raw

import (
	"github.com/orcaman/concurrent-map"
)

type ConcurrentMapStore struct {
	store cmap.ConcurrentMap
}
var _ Store = (*ConcurrentMapStore)(nil)

func NewConcurrentMapStore() *ConcurrentMapStore {
	return &ConcurrentMapStore{
		cmap.New(),
	}
}

func (s *ConcurrentMapStore) Close() error {
	return nil
}

func (s *ConcurrentMapStore) Iterate(fn IteratorFunc) error {
	s.store.IterCb(func(key string, v interface{}) {
		_, _ = fn(key, v.([]byte))
	})
	return nil
}

func (s *ConcurrentMapStore) Get(key string) ([]byte, error) {
	v, ok := s.store.Get(key)
	if !ok {
		return nil, ErrNotFound
	}
	return v.([]byte), nil
}

func (s *ConcurrentMapStore) Set(key string, data []byte) error {
	s.store.Set(key, data)
	return nil
}

func (s *ConcurrentMapStore) Delete(key string) error {
	s.store.Remove(key)
	return nil
}

func (s *ConcurrentMapStore) Reset() error {
	s.store = cmap.New()
	return nil
}

func (s *ConcurrentMapStore) Len() int {
	return s.store.Count()
}

func (s *ConcurrentMapStore) Capacity() int {
	return s.store.Count()
}

