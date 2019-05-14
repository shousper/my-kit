package local

import (
	"github.com/orcaman/concurrent-map"
	"github.com/shousper/my-kit/store/raw"
)

type ConcurrentMapStore struct {
	store cmap.ConcurrentMap
}

var _ raw.Store = (*ConcurrentMapStore)(nil)

func NewConcurrentMapStore() *ConcurrentMapStore {
	return &ConcurrentMapStore{
		cmap.New(),
	}
}

func (s *ConcurrentMapStore) Close() error {
	return nil
}

func (s *ConcurrentMapStore) Iterate() raw.Iterator {
	return raw.NewIterator(s)
}

func (s *ConcurrentMapStore) Get(key string) (raw.Value, error) {
	v, ok := s.store.Get(key)
	if !ok {
		return nil, raw.ErrNotFound
	}
	return v.(raw.Value), nil
}

func (s *ConcurrentMapStore) Set(key string, data raw.Value) error {
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

func (s *ConcurrentMapStore) Keys() ([]string, error) {
	return s.store.Keys(), nil
}
