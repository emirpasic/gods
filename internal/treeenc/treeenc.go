package treeenc

import (
	"encoding"
	"encoding/json"
	"reflect"
	"strconv"
)

// KeyMarshaler is a helper type for marshaling keys of a tree.
// When marshaling a tree, we need first to convert tree key/value
// pairs to a standard Go map, and then marshal the map. However,
// Go maps can only have keys of comparable types, and so we wrap
// the key in a KeyMarshaler and implement the encoding.TextMarshaler
// interface to make it comparable and marshalable.
//
// The map should be declared as map[*KeyMarshaler[K]]V.
type KeyMarshaler[K any] struct {
	Key K
}

var _ encoding.TextMarshaler = &KeyMarshaler[string]{}

func (m *KeyMarshaler[T]) MarshalText() ([]byte, error) {
	kv := reflect.ValueOf(m.Key)
	if tm, ok := kv.Interface().(encoding.TextMarshaler); ok {
		if kv.Kind() == reflect.Pointer && kv.IsNil() {
			return nil, nil
		}
		return tm.MarshalText()
	}

	var text string
	switch kv.Kind() {
	case reflect.String:
		text = kv.String()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		text = strconv.FormatInt(kv.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		text = strconv.FormatUint(kv.Uint(), 10)
	default:
		return nil, &json.UnsupportedTypeError{Type: kv.Type()}
	}
	return []byte(text), nil
}

// KeyUnmarshaler is a helper type for unmarshaling keys of a tree.
// When unmarshaling a tree, we first unmarshal the JSON into a Go
// map, and then convert the map to tree key/value pairs. Similar to
// KeyMarshaler, we wrap the key in a KeyUnmarshaler to make it
// unmarshalable by implementing the encoding.TextUnmarshaler interface.
//
// The map should be declared as map[KeyUnmarshaler[K]]V.
type KeyUnmarshaler[K any] struct {
	Key *K
}

var _ encoding.TextUnmarshaler = &KeyUnmarshaler[string]{}

func (m *KeyUnmarshaler[K]) UnmarshalText(text []byte) error {
	var key K
	m.Key = &key

	kv := reflect.ValueOf(m.Key)
	if tu, ok := kv.Interface().(encoding.TextUnmarshaler); ok {
		if kv.Kind() == reflect.Ptr && kv.IsNil() {
			return nil
		}
		return tu.UnmarshalText(text)
	}

	var err error
	kv = kv.Elem()
	switch kv.Kind() {
	case reflect.String:
		kv.SetString(string(text))
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		var i int64
		i, err = strconv.ParseInt(string(text), 10, 64)
		kv.SetInt(i)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		var u uint64
		u, err = strconv.ParseUint(string(text), 10, 64)
		kv.SetUint(u)
	default:
		err = &json.UnsupportedTypeError{Type: kv.Type()}
	}
	return err
}
