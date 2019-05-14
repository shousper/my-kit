package store

import (
	"github.com/shousper/my-kit/store/object"
	"github.com/shousper/my-kit/store/object/encoded"
	"github.com/shousper/my-kit/store/object/generic"
	"github.com/shousper/my-kit/store/raw"
	"github.com/shousper/my-kit/store/raw/local"
	"github.com/shousper/my-kit/store/raw/remote"
)

type Store = raw.Store

var (
	/* Store implementations */

	NewMapRawStore           = local.NewDefaultStore
	NewBigCacheRawStore      = local.NewBigCacheStore
	NewConcurrentMapRawStore = local.NewConcurrentMapStore
	NewFreeCacheRawStore     = local.NewFreeCacheStore

	NewRedisRawStore = remote.NewRedisStore

	/* Feature wrappers */
	NewLazyRawStore    = raw.NewLazyStore
	NewEventedRawStore = raw.NewEventedStore
)

type ObjectStore = object.Store

var (
	/* Store implementations */
	NewGenericObjectStore = generic.NewStore
	NewLRU2QObjectStore   = generic.NewLRU2QStore

	NewEncodedObjectStore          = encoded.NewStore
	NewGobEncodedObjectStore       = encoded.NewGobStore
	NewGogoProtoEncodedObjectStore = encoded.NewGogoProtoStore
	NewProtoEncodedObjectStore     = encoded.NewProtoStore
	NewJSONEncodedObjectStore      = encoded.NewJSONStore

	/* Feature wrappers */
	NewLazyObjectStore    = object.NewLazyStore
	NewEventedObjectStore = object.NewEventedStore
	NewIndexedObjectStore = object.NewIndexedStore
)

type Event = raw.Event

var (
	EventBeforeGet    = raw.EventBeforeGet
	EventAfterGet     = raw.EventAfterGet
	EventBeforeSet    = raw.EventBeforeSet
	EventAfterSet     = raw.EventAfterSet
	EventBeforeDelete = raw.EventBeforeDelete
	EventAfterDelete  = raw.EventAfterDelete
	EventBeforeReset  = raw.EventBeforeReset
	EventAfterReset   = raw.EventAfterReset
)

type (
	RawEventCallback    = raw.EventCallback
	ObjectEventCallback = object.EventCallback
)
