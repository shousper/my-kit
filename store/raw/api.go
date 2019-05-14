package raw

import "io"

type Value []byte

type IteratorFunc func(key string, value Value) (bool, error)

type Store interface {
	io.Closer

	Keys() ([]string, error)
	Iterate() Iterator
	Get(key string) (Value, error)
	Set(key string, data Value) error
	Delete(key string) error
	Reset() error
	Len() int
	Capacity() int
}
