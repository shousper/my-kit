package object

import "io"

type Value interface{}

type Store interface {
	io.Closer

	Keys() ([]string, error)
	Iterate() Iterator
	Get(key string, out Value) error
	Set(key string, in Value) error
	Delete(key string) error
	Reset() error
	Len() int
	Capacity() int
}
