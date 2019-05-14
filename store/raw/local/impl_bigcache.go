package local

import (
	"time"

	"github.com/allegro/bigcache"
	"github.com/shousper/my-kit/store/raw"
)

type BigCacheStore struct {
	store *bigcache.BigCache
}

var _ raw.Store = (*BigCacheStore)(nil)

func NewBigCacheStore(eviction time.Duration) *BigCacheStore {
	cfg := bigcache.DefaultConfig(eviction)
	cfg.Verbose = false
	cache, _ := bigcache.NewBigCache(cfg)
	return &BigCacheStore{
		store: cache,
	}
}

func (s *BigCacheStore) Get(key string) (raw.Value, error) {
	if key == "" {
		return nil, raw.ErrInvalidKey
	}

	v, err := s.store.Get(key)
	return v, s.normalizeError(err)
}

func (s *BigCacheStore) Set(key string, in raw.Value) error {
	if key == "" {
		return raw.ErrInvalidKey
	}
	if in == nil {
		return raw.ErrInvalidValue
	}
	err := s.store.Set(key, in)
	return s.normalizeError(err)
}

func (s *BigCacheStore) Delete(key string) error {
	if key == "" {
		return raw.ErrInvalidKey
	}
	err := s.store.Delete(key)
	return s.normalizeError(err)
}

func (s *BigCacheStore) Keys() ([]string, error) {
	out := make([]string, s.store.Len())
	i := 0

	iterator := s.store.Iterator()
	for iterator.SetNext() {
		entry, err := iterator.Value()
		if err != nil {
			return nil, err
		}

		out[i] = entry.Key()
		i++
	}
	return out, nil
}

func (s *BigCacheStore) Iterate() raw.Iterator {
	return &bigCacheIterator{
		iterator: s.store.Iterator(),
	}
}

func (s *BigCacheStore) Close() error {
	return s.store.Close()
}

func (s *BigCacheStore) Reset() error {
	return s.store.Reset()
}

func (s *BigCacheStore) Len() int {
	return s.store.Len()
}

func (s *BigCacheStore) Capacity() int {
	return s.store.Capacity()
}

func (s *BigCacheStore) normalizeError(err error) error {
	if err == bigcache.ErrEntryNotFound {
		return raw.ErrNotFound
	}
	return err
}

type bigCacheIterator struct {
	iterator *bigcache.EntryInfoIterator
	err      error
}

func (i *bigCacheIterator) Next() (raw.Value, bool) {
	if i.err != nil {
		return nil, false
	}
	if !i.iterator.SetNext() {
		return nil, false
	}
	v, err := i.iterator.Value()
	if err != nil {
		i.err = err
		return nil, false
	}
	return v.Value(), true
}

func (i *bigCacheIterator) Err() error {
	return i.err
}
