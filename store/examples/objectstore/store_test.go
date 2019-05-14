package objectstore_test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/shousper/my-kit/store/examples/objectstore"
	"github.com/stretchr/testify/assert"
)

func TestItemStore(t *testing.T) {
	s := objectstore.NewItemStore()

	assert.NoError(t, s.Set(&objectstore.Item{ID: "1", Name: "Bar"}))
	v, err := s.GetByName("Bar")
	assert.NoError(t, err)
	assert.NotNil(t, v)
	assert.Equal(t, "1", v.ID)
}

func TestItemStore_Memory(t *testing.T) {
	s, _, _ := makeSet(t, 1000, 16)
	time.Sleep(500 * time.Millisecond)
	addItems(t, s, 1000, 1024)
	for i := 0; i < 100; i++ {
		time.Sleep(10 * time.Millisecond)
		addItems(t, s, 5, 256)
	}
	addItems(t, s, 1000, 1024)
	time.Sleep(1 * time.Second)
}

var result interface{}

func makeSet(tb testing.TB, n int, size int) (store *objectstore.ItemStore, ids, names []string) {
	store = objectstore.NewItemStore()
	ids, names = addItems(tb, store, n, size)
	return
}

func addItems(tb testing.TB, store *objectstore.ItemStore, n int, size int) (ids, names []string) {
	s2, s4 := size/2, size/4

	ids = make([]string, n)
	names = make([]string, n)

	for i := 0; i < n; i++ {
		data := make([]byte, s2)
		rand.Read(data)

		entry := objectstore.Item{
			ID:   fmt.Sprintf("%x", data[:s4]),
			Name: fmt.Sprintf("%x", data[s4:]),
		}
		ids[i] = entry.ID
		names[i] = entry.Name

		assert.NoError(tb, store.Set(&entry))
	}
	return
}

func BenchmarkItemStore(b *testing.B) {
	b.Run("Get 10000/16", func(b *testing.B) {
		set, ids, _ := makeSet(b, 10000, 16)
		l := len(ids)
		var out interface{}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			b.StopTimer()
			id := ids[i%l]
			b.StartTimer()

			out, _ = set.Get(id)
		}
		result = out
	})

	b.Run("Get 10000/1024", func(b *testing.B) {
		set, ids, _ := makeSet(b, 10000, 1024)
		l := len(ids)
		var out interface{}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			b.StopTimer()
			id := ids[i%l]
			b.StartTimer()

			out, _ = set.Get(id)
		}
		result = out
	})

	b.Run("GetByName 10000/16", func(b *testing.B) {
		set, _, names := makeSet(b, 10000, 16)
		l := len(names)
		var out interface{}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			b.StopTimer()
			name := names[i%l]
			b.StartTimer()

			out, _ = set.GetByName(name)
		}
		result = out
	})

	b.Run("GetByName 10000/1024", func(b *testing.B) {
		set, _, names := makeSet(b, 10000, 1024)
		l := len(names)
		var out interface{}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			b.StopTimer()
			name := names[i%l]
			b.StartTimer()

			out, _ = set.GetByName(name)
		}
		result = out
	})

	b.Run("Set 10000/16", func(b *testing.B) {
		set, ids, _ := makeSet(b, 10000, 16)
		l := len(ids)

		out := objectstore.NewItemStore()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			b.StopTimer()
			k := ids[i%l]
			v, _ := set.Get(k)
			b.StartTimer()

			_ = out.Set(v)
		}
		result = out
	})

	b.Run("Set 10000/1024", func(b *testing.B) {
		set, ids, _ := makeSet(b, 10000, 1024)
		l := len(ids)

		out := objectstore.NewItemStore()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			b.StopTimer()
			k := ids[i%l]
			v, _ := set.Get(k)
			b.StartTimer()

			_ = out.Set(v)
		}
		result = out
	})
}
