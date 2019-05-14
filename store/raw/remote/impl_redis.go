package remote

import (
	"github.com/go-redis/redis"
	"github.com/shousper/my-kit/store/raw"
)

const (
	redisScanCount = 100
)

type RedisStore struct {
	cl        redis.UniversalClient
	namespace string

	// Pre-computed pattern for matching keys that belong to the namespace of this store
	pattern string
}

var _ raw.Store = (*RedisStore)(nil)

func NewRedisStore(cl redis.UniversalClient, namespace string) *RedisStore {
	return &RedisStore{
		cl:        cl,
		namespace: namespace,

		pattern: namespace + ":*",
	}
}

func (s *RedisStore) namespaceKey(in string) string {
	return s.namespace + ":" + in
}

func (s *RedisStore) Get(key string) (raw.Value, error) {
	if key == "" {
		return nil, raw.ErrInvalidKey
	}

	v, err := s.cl.Get(s.namespaceKey(key)).Bytes()
	if err != nil {
		return nil, err
	}
	if len(v) == 0 {
		return nil, raw.ErrNotFound
	}
	return v, nil
}

func (s *RedisStore) Set(key string, in raw.Value) error {
	if key == "" {
		return raw.ErrInvalidKey
	}
	if in == nil {
		return raw.ErrInvalidValue
	}
	return s.cl.Set(s.namespaceKey(key), in, -1).Err()
}

func (s *RedisStore) Delete(key string) error {
	if key == "" {
		return raw.ErrInvalidKey
	}

	return s.cl.Del(s.namespaceKey(key)).Err()
}

func (s *RedisStore) Keys() ([]string, error) {
	return s.cl.Keys(s.pattern).Result()
}

func (s *RedisStore) Iterate() raw.Iterator {
	return &redisIterator{
		iter: s.cl.Scan(0, s.pattern, redisScanCount).Iterator(),
	}
}

func (s *RedisStore) Close() error {
	return s.cl.Close()
}

func (s *RedisStore) Reset() error {
	_, err := s.cl.TxPipelined(func(p redis.Pipeliner) error {
		keys, err := p.Keys(s.pattern).Result()
		if err != nil {
			return err
		}

		return p.Del(keys...).Err()
	})
	return err
}

func (s *RedisStore) Len() int {
	// TODO("Make more efficient? Maybe use an atomic integer to track this?")
	keys, err := s.cl.Keys(s.pattern).Result()
	if err != nil {
		return 0
	}
	return len(keys)
}

func (s *RedisStore) Capacity() int {
	// Doesn't really work for redis.. so we'll just return Len?
	return s.Len()
}

type redisIterator struct {
	iter *redis.ScanIterator
}

func (i *redisIterator) Next() (raw.Value, bool) {
	if i.iter.Err() != nil {
		return nil, false
	}
	if !i.iter.Next() {
		return nil, false
	}
	return raw.Value(i.iter.Val()), true
}

func (i *redisIterator) Err() error {
	return i.iter.Err()
}
