package object

import (
	"encoding/json"
	"my-kit/store/raw"
)

func NewJSONStore(store raw.Store) *DefaultStore {
	return NewStore(store, json.Marshal, json.Unmarshal)
}
