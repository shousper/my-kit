package generic

import (
	"reflect"

	"github.com/hashicorp/golang-lru"
	"github.com/shousper/my-kit/store/object"
)

type LRU2QStore struct {
	store *lru.TwoQueueCache
}

var _ object.Store = (*LRU2QStore)(nil)

func NewLRU2QStore(size int) *LRU2QStore {
	cache, _ := lru.New2Q(size)
	return &LRU2QStore{
		store: cache,
	}
}

func (s *LRU2QStore) Get(key string, out object.Value) error {
	if key == "" {
		return object.ErrInvalidKey
	}

	v, ok := s.store.Get(key)
	if !ok {
		return object.ErrNotFound
	}

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

func (s *LRU2QStore) Set(key string, in object.Value) error {
	if key == "" {
		return object.ErrInvalidKey
	}
	if in == nil {
		return object.ErrInvalidValue
	}
	s.store.Add(key, in)
	return nil
}

func (s *LRU2QStore) Delete(key string) error {
	if key == "" {
		return object.ErrInvalidKey
	}
	s.store.Remove(key)
	return nil
}

func (s *LRU2QStore) Keys() ([]string, error) {
	keys := s.store.Keys()

	out := make([]string, len(keys))
	i := 0
	for _, key := range keys {
		out[i] = key.(string)
		i++
	}
	return out, nil
}

func (s *LRU2QStore) Iterate() object.Iterator {
	return object.NewIterator(s)
}

func (s *LRU2QStore) Close() error {
	return nil
}

func (s *LRU2QStore) Reset() error {
	s.store.Purge()
	return nil
}

func (s *LRU2QStore) Len() int {
	return s.store.Len()
}

func (s *LRU2QStore) Capacity() int {
	return s.store.Len()
}
