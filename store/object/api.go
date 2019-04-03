package object

import "io"

type IteratorFunc func(key string, value interface{}) (bool, error)

type Store interface {
	io.Closer

	Iterate(referenceType interface{}, fn IteratorFunc) error
	Get(key string, out interface{}) error
	Set(key string, in interface{}) error
	Delete(key string) error
	Reset() error
	Len() int
	Capacity() int
}

