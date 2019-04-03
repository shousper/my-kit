package store_test

import (
	"my-kit/store"
	"my-kit/store/object"
	"testing"

	"github.com/golang/protobuf/proto"
	"github.com/stretchr/testify/assert"
)

type Foo struct {
	ID string    `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (m *Foo) ProtoMessage()  {}
func (m *Foo) Reset()         { *m = Foo{} }
func (m *Foo) String() string { return proto.CompactTextString(m) }

type FooStore struct {
	*object.IndexedStore
}

func NewFooStore(t *testing.T) *FooStore {
	// Create the base, raw key-value string-[]byte store
	rawStore := store.NewStore() // TODO("Support other backends and configuration")

	// Wrap with object store that uses protobuf encoding
	objectStore := object.NewProtoStore(rawStore)

	// Wrap with event hooks for get, set, delete and reset
	hookedStore := store.NewHookedObjectStore(objectStore)
	hookedStore.On(object.HookBeforeGet, func(key string, value interface{}) {
		t.Logf("HookBeforeGet: %v = %#v", key, value)
	})
	hookedStore.On(object.HookAfterGet, func(key string, value interface{}) {
		t.Logf("HookAfterGet: %v = %#v", key, value)
	})
	hookedStore.On(object.HookBeforeSet, func(key string, value interface{}) {
		t.Logf("HookBeforeSet: %v = %#v", key, value)
	})
	hookedStore.On(object.HookAfterSet, func(key string, value interface{}) {
		t.Logf("HookAfterSet: %v = %#v", key, value)
	})

	// Wrap with lazy getter when entries are missing
	lazyStore := store.NewLazyObjectStore(hookedStore, func(key string) (i interface{}, e error) {
		return nil, nil
	})

	// Wrap with indexing
	indexStore := store.NewIndexedObjectStore(lazyStore).
		Index("name", func(entity interface{}) (key string) {
			return entity.(*Foo).Name
		})

	// Finally, return concrete store implementation
	return &FooStore{indexStore}
}

func (s *FooStore) GetByName(name string) (*Foo, error) {
	out := Foo{Name: name}
	return &out, s.IndexedStore.GetBy("name", &out)
}

func (s *FooStore) Get(id string) (*Foo, error) {
	var out Foo
	return &out, s.IndexedStore.Get(id, &out)
}

func (s *FooStore) Set(in *Foo) error {
	return s.IndexedStore.Set(in.ID, in)
}

func TestNewFooStore(t *testing.T) {
	s := NewFooStore(t)

	assert.NoError(t, s.Set(&Foo{ID: "1", Name: "Bar"}))
	v, err := s.GetByName("Bar")
	assert.NoError(t, err)
	assert.NotNil(t, v)
	assert.Equal(t, "1", v.ID)
}
