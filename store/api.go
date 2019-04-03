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
	NewHookedStore = raw.NewHookedStore
)

type ObjectStore = object.Store

var (
	/* Store implementations */
	NewObjectStore = object.NewStore

	/* Feature wrappers */
	NewLazyObjectStore = object.NewLazyStore
	NewHookedObjectStore = object.NewHookedStore
	NewIndexedObjectStore = object.NewIndexedStore
)
