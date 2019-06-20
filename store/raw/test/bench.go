package test

import (
	"fmt"
	"math/rand"
	"testing"

	"my-kit/store/raw"

	"github.com/stretchr/testify/assert"
)

func AddItems(tb testing.TB, s raw.Store, n int, size int) (keys []string, values [][]byte) {
	keys = make([]string, n)
	values = make([][]byte, n)
	for i := 0; i < n; i++ {
		var keyBytes [16]byte
		rand.Read(keyBytes[:])
		key := fmt.Sprintf("%x", keyBytes)
		value := make([]byte, size)
		rand.Read(value)

		keys[i] = key
		values[i] = value
		assert.NoError(tb, s.Set(key, value))
	}
	return
}

var result interface{}

func Benchmark(b *testing.B, fn func() raw.Store) {
	var configurations = []struct {
		capacity, size int
	}{
		{1000, 16},
		{1000, 1024},
		{1000, 262144},
		{1000000, 16},
		{1000000, 1024},
	}

	for _, cfg := range configurations {
		b.Run(fmt.Sprintf("Get %d/%d", cfg.capacity, cfg.size), func(b *testing.B) {
			set := fn()
			keys, _ := AddItems(b, set, cfg.capacity, cfg.size)
			l := len(keys)
			var out interface{}
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				out, _ = set.Get(keys[i%l])
			}
			result = out
		})
	}
	for _, cfg := range configurations {
		b.Run(fmt.Sprintf("Set %d/%d", cfg.capacity, cfg.size), func(b *testing.B) {
			set := fn()
			keys, values := AddItems(b, set, cfg.capacity, cfg.size)
			l := len(keys)

			out := fn()
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = out.Set(keys[i%l], values[i%l])
			}
			result = out
		})
	}
	for _, cfg := range configurations {
		b.Run(fmt.Sprintf("Delete %d/%d", cfg.capacity, cfg.size), func(b *testing.B) {
			set := fn()
			keys, _ := AddItems(b, set, cfg.capacity, cfg.size)
			l := len(keys)

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = set.Delete(keys[i%l])
			}
			result = set
		})
	}
	for _, cfg := range configurations {
		b.Run(fmt.Sprintf("Reset %d/%d", cfg.capacity, cfg.size), func(b *testing.B) {
			sets := make([]raw.Store, b.N)
			for i := 0; i < b.N; i++ {
				sets[i] = fn()
			}

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = sets[i].Reset()
			}
			result = sets
		})
	}
	for _, cfg := range configurations {
		b.Run(fmt.Sprintf("Iterate %d/%d", cfg.capacity, cfg.size), func(b *testing.B) {
			set := fn()
			_, _ = AddItems(b, set, cfg.capacity, cfg.size)

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = set.Iterate(func(key string, value []byte) (b bool, e error) {
					result = value
					return true, nil
				})
			}
		})
	}
}
