package objectstore

import (
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/shousper/my-kit/store"
	"github.com/shousper/my-kit/store/object"
	"github.com/shousper/my-kit/store/object/encoded"
)

type Item struct {
	ID   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (m *Item) ProtoMessage()  {}
func (m *Item) Reset()         { *m = Item{} }
func (m *Item) String() string { return proto.CompactTextString(m) }

type ItemStore struct {
	*object.IndexedStore

	gets, sets uint64
}

func NewItemStore() *ItemStore {
	// Create the base, raw key-value string-[]byte store
	rawStore := store.NewBigCacheRawStore(1 * time.Second) // TODO("Support other backends and configuration")

	// Wrap with object store that uses protobuf encoding
	objectStore := encoded.NewProtoStore(rawStore)

	// Wrap with event hooks for get, set, delete and reset
	hookedStore := store.NewEventedObjectStore(objectStore)

	// Wrap with lazy getter when entries are missing
	lazyStore := store.NewLazyObjectStore(hookedStore, func(key string) (i object.Value, e error) {
		return nil, nil
	})

	// Wrap with indexing
	indexStore := store.NewIndexedObjectStore(lazyStore).
		Index("name", func(entity object.Value) (key string) {
			return entity.(*Item).Name
		})

	// Create the concrete store implementation
	s := &ItemStore{IndexedStore: indexStore}

	// Bind event handlers
	hookedStore.On(store.EventAfterGet, s.onAfterGet)
	hookedStore.On(store.EventAfterSet, s.onAfterSet)

	// Return composed store
	return s
}

func (s *ItemStore) onAfterGet(key string, value object.Value) {
	s.gets++
}

func (s *ItemStore) onAfterSet(key string, value object.Value) {
	s.sets++
}

func (s *ItemStore) GetByName(name string) (*Item, error) {
	out := Item{Name: name}
	return &out, s.IndexedStore.GetBy("name", &out)
}

func (s *ItemStore) Get(id string) (*Item, error) {
	var out Item
	return &out, s.IndexedStore.Get(id, &out)
}

func (s *ItemStore) Set(in *Item) error {
	return s.IndexedStore.Set(in.ID, in)
}
