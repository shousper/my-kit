package encoded

import (
	"encoding/json"
	"github.com/shousper/my-kit/store/raw"
)

func NewJSONStore(store raw.Store) *DefaultStore {
	return NewStore(store, json.Marshal, json.Unmarshal)
}
