package store

import (
	"my-kit/store/object"
	"my-kit/store/raw"
)

type Store = raw.Store

var (
	// NewStore defaults to bigcache 🤷‍
	NewStore = NewBigCacheStore

	/* Store implementations */
	NewMapStore = raw.NewDefaultStore
	NewBigCacheStore = raw.NewBigCacheStore

	/* Feature wrappers */
	NewLazyStore = raw.NewLazyStore
	NewEventedStore = raw.NewEventedStore
	NewChanneledStore = raw.NewChanneledStore
)

type ObjectStore = object.Store

var (
	/* Store implementations */
	NewObjectStore = object.NewStore

	/* Feature wrappers */
	NewLazyObjectStore = object.NewLazyStore
	NewEventedObjectStore = object.NewEventedStore
	NewChanneledObjectStore = object.NewChanneledStore
	NewIndexedObjectStore = object.NewIndexedStore
)

type Event = raw.Event

var (
	EventBeforeGet = raw.EventBeforeGet
	EventAfterGet = raw.EventAfterGet
	EventBeforeSet = raw.EventBeforeSet
	EventAfterSet = raw.EventAfterSet
	EventBeforeDelete = raw.EventBeforeDelete
	EventAfterDelete = raw.EventAfterDelete
	EventBeforeReset = raw.EventBeforeReset
	EventAfterReset = raw.EventAfterReset
)

type RawEventData = raw.EventData
type RawEventCallback = raw.EventCallback

type ObjectEventData = object.EventData
type ObjectEventCallback = object.EventCallback
