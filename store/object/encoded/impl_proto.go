package encoded

import (
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/pkg/errors"
	"github.com/shousper/my-kit/store/raw"
)

func NewProtoStore(store raw.Store) *DefaultStore {
	return NewStore(store, protoMarshal, protoUnmarshal)
}

func protoMarshal(in interface{}) ([]byte, error) {
	switch v := in.(type) {
	case proto.Message:
		return proto.Marshal(v)
	case *bool:
		return proto.Marshal(&wrappers.BoolValue{Value: *v})
	case bool:
		return proto.Marshal(&wrappers.BoolValue{Value: v})
	case *uint32:
		return proto.Marshal(&wrappers.UInt32Value{Value: *v})
	case uint32:
		return proto.Marshal(&wrappers.UInt32Value{Value: v})
	case *uint64:
		return proto.Marshal(&wrappers.UInt64Value{Value: *v})
	case uint64:
		return proto.Marshal(&wrappers.UInt64Value{Value: v})
	case *int32:
		return proto.Marshal(&wrappers.Int32Value{Value: *v})
	case int32:
		return proto.Marshal(&wrappers.Int32Value{Value: v})
	case *int64:
		return proto.Marshal(&wrappers.Int64Value{Value: *v})
	case int64:
		return proto.Marshal(&wrappers.Int64Value{Value: v})
	case *float32:
		return proto.Marshal(&wrappers.FloatValue{Value: *v})
	case float32:
		return proto.Marshal(&wrappers.FloatValue{Value: v})
	case *float64:
		return proto.Marshal(&wrappers.DoubleValue{Value: *v})
	case float64:
		return proto.Marshal(&wrappers.DoubleValue{Value: v})
	case *string:
		return proto.Marshal(&wrappers.StringValue{Value: *v})
	case string:
		return proto.Marshal(&wrappers.StringValue{Value: v})
	case *[]byte:
		return proto.Marshal(&wrappers.BytesValue{Value: *v})
	case []byte:
		return proto.Marshal(&wrappers.BytesValue{Value: v})
	case nil:
		return nil, nil
	}
	return nil, errors.Errorf("invalid proto: %t", in)
}

func protoUnmarshal(data []byte, out interface{}) error {
	switch v := out.(type) {
	case proto.Message:
		return proto.Unmarshal(data, v)
	case *bool:
		wrapper := new(wrappers.BoolValue)
		err := proto.Unmarshal(data, wrapper)
		*v = wrapper.Value
		return err
	case *uint32:
		wrapper := new(wrappers.UInt32Value)
		err := proto.Unmarshal(data, wrapper)
		*v = wrapper.Value
		return err
	case *uint64:
		wrapper := new(wrappers.UInt64Value)
		err := proto.Unmarshal(data, wrapper)
		*v = wrapper.Value
		return err
	case *int32:
		wrapper := new(wrappers.Int32Value)
		err := proto.Unmarshal(data, wrapper)
		*v = wrapper.Value
		return err
	case *int64:
		wrapper := new(wrappers.Int64Value)
		err := proto.Unmarshal(data, wrapper)
		*v = wrapper.Value
		return err
	case *float32:
		wrapper := new(wrappers.FloatValue)
		err := proto.Unmarshal(data, wrapper)
		*v = wrapper.Value
		return err
	case *float64:
		wrapper := new(wrappers.DoubleValue)
		err := proto.Unmarshal(data, wrapper)
		*v = wrapper.Value
		return err
	case *string:
		wrapper := new(wrappers.StringValue)
		err := proto.Unmarshal(data, wrapper)
		*v = wrapper.Value
		return err
	case *[]byte:
		wrapper := new(wrappers.BytesValue)
		err := proto.Unmarshal(data, wrapper)
		*v = wrapper.Value
		return err
	case nil:
		return nil
	}
	return errors.Errorf("invalid proto: %t", out)
}
