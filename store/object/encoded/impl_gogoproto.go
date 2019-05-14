package encoded

import (
	"github.com/gogo/protobuf/proto"
	"github.com/gogo/protobuf/types"
	"github.com/pkg/errors"
	"github.com/shousper/my-kit/store/raw"
)

func NewGogoProtoStore(store raw.Store) *DefaultStore {
	return NewStore(store, gogoProtoMarshal, gogoProtoUnmarshal)
}

func gogoProtoMarshal(in interface{}) ([]byte, error) {
	switch v := in.(type) {
	case proto.Message:
		return proto.Marshal(v)
	case *bool:
		return proto.Marshal(&types.BoolValue{Value: *v})
	case bool:
		return proto.Marshal(&types.BoolValue{Value: v})
	case *uint32:
		return proto.Marshal(&types.UInt32Value{Value: *v})
	case uint32:
		return proto.Marshal(&types.UInt32Value{Value: v})
	case *uint64:
		return proto.Marshal(&types.UInt64Value{Value: *v})
	case uint64:
		return proto.Marshal(&types.UInt64Value{Value: v})
	case *int32:
		return proto.Marshal(&types.Int32Value{Value: *v})
	case int32:
		return proto.Marshal(&types.Int32Value{Value: v})
	case *int64:
		return proto.Marshal(&types.Int64Value{Value: *v})
	case int64:
		return proto.Marshal(&types.Int64Value{Value: v})
	case *float32:
		return proto.Marshal(&types.FloatValue{Value: *v})
	case float32:
		return proto.Marshal(&types.FloatValue{Value: v})
	case *float64:
		return proto.Marshal(&types.DoubleValue{Value: *v})
	case float64:
		return proto.Marshal(&types.DoubleValue{Value: v})
	case *string:
		return proto.Marshal(&types.StringValue{Value: *v})
	case string:
		return proto.Marshal(&types.StringValue{Value: v})
	case *[]byte:
		return proto.Marshal(&types.BytesValue{Value: *v})
	case []byte:
		return proto.Marshal(&types.BytesValue{Value: v})
	case nil:
		return nil, nil
	}
	return nil, errors.Errorf("invalid proto: %t", in)
}

func gogoProtoUnmarshal(data []byte, out interface{}) error {
	switch v := out.(type) {
	case proto.Message:
		return proto.Unmarshal(data, v)
	case *bool:
		wrapper := new(types.BoolValue)
		err := proto.Unmarshal(data, wrapper)
		*v = wrapper.Value
		return err
	case *uint32:
		wrapper := new(types.UInt32Value)
		err := proto.Unmarshal(data, wrapper)
		*v = wrapper.Value
		return err
	case *uint64:
		wrapper := new(types.UInt64Value)
		err := proto.Unmarshal(data, wrapper)
		*v = wrapper.Value
		return err
	case *int32:
		wrapper := new(types.Int32Value)
		err := proto.Unmarshal(data, wrapper)
		*v = wrapper.Value
		return err
	case *int64:
		wrapper := new(types.Int64Value)
		err := proto.Unmarshal(data, wrapper)
		*v = wrapper.Value
		return err
	case *float32:
		wrapper := new(types.FloatValue)
		err := proto.Unmarshal(data, wrapper)
		*v = wrapper.Value
		return err
	case *float64:
		wrapper := new(types.DoubleValue)
		err := proto.Unmarshal(data, wrapper)
		*v = wrapper.Value
		return err
	case *string:
		wrapper := new(types.StringValue)
		err := proto.Unmarshal(data, wrapper)
		*v = wrapper.Value
		return err
	case *[]byte:
		wrapper := new(types.BytesValue)
		err := proto.Unmarshal(data, wrapper)
		*v = wrapper.Value
		return err
	case nil:
		return nil
	}
	return errors.Errorf("invalid proto: %t", out)
}
