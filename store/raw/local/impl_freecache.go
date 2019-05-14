package local

import (
	"time"

	"github.com/coocood/freecache"
	"github.com/shousper/my-kit/store/raw"
)

type FreeCacheStore struct {
	cache *freecache.Cache

	expiration time.Duration
}

var _ raw.Store = (*FreeCacheStore)(nil)

func NewFreeCacheStore(size int, expiration time.Duration) *FreeCacheStore {
	return &FreeCacheStore{
		freecache.NewCache(size),
		expiration,
	}
}

func (s *FreeCacheStore) Get(key string) (raw.Value, error) {
	if key == "" {
		return nil, raw.ErrInvalidKey
	}

	v, err := s.cache.Get(raw.Value(key))
	return v, s.normalizeError(err)
}

func (s *FreeCacheStore) Set(key string, in raw.Value) error {
	if key == "" {
		return raw.ErrInvalidKey
	}
	err := s.cache.Set(raw.Value(key), in, int(s.expiration.Seconds()))
	return s.normalizeError(err)
}

func (s *FreeCacheStore) Delete(key string) error {
	if key == "" {
		return raw.ErrInvalidKey
	}
	s.cache.Del(raw.Value(key))
	return nil
}

func (s *FreeCacheStore) Keys() ([]string, error) {
	iterator := s.cache.NewIterator()

	out := make([]string, s.cache.EntryCount())
	i := 0
	for {
		if e := iterator.Next(); e != nil {
			out[i] = string(e.Key)
			i++
			continue
		}
		break
	}
	return out, nil
}

func (s *FreeCacheStore) Iterate() raw.Iterator {
	return &freeCacheIterator{
		iter: s.cache.NewIterator(),
	}
}

func (s *FreeCacheStore) Close() error {
	return nil
}

func (s *FreeCacheStore) Reset() error {
	s.cache.Clear()
	return nil
}

func (s *FreeCacheStore) Len() int {
	return int(s.cache.EntryCount())
}

func (s *FreeCacheStore) Capacity() int {
	return -1
}

func (s *FreeCacheStore) normalizeError(err error) error {
	if err == freecache.ErrNotFound {
		return raw.ErrNotFound
	}
	return err
}

type freeCacheIterator struct {
	iter *freecache.Iterator
}

func (i *freeCacheIterator) Next() (raw.Value, bool) {
	if e := i.iter.Next(); e != nil {
		return e.Value, true
	}
	return nil, false
}

func (i *freeCacheIterator) Err() error {
	return nil
}
