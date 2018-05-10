// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package utils provides common utility functions.
//
// Provided functionalities:
// - sorting
// - comparators
package utils

import (
	"fmt"
	"strconv"
	"time"
)

// ToString converts a value to string.
func ToString(value interface{}) string {
	switch value.(type) {
	case string:
		return value.(string)
	case int8:
		return strconv.FormatInt(int64(value.(int8)), 10)
	case int16:
		return strconv.FormatInt(int64(value.(int16)), 10)
	case int32:
		return strconv.FormatInt(int64(value.(int32)), 10)
	case int64:
		return strconv.FormatInt(int64(value.(int64)), 10)
	case uint8:
		return strconv.FormatUint(uint64(value.(uint8)), 10)
	case uint16:
		return strconv.FormatUint(uint64(value.(uint16)), 10)
	case uint32:
		return strconv.FormatUint(uint64(value.(uint32)), 10)
	case uint64:
		return strconv.FormatUint(uint64(value.(uint64)), 10)
	case float32:
		return strconv.FormatFloat(float64(value.(float32)), 'g', -1, 64)
	case float64:
		return strconv.FormatFloat(float64(value.(float64)), 'g', -1, 64)
	case time.Time:
		return value.(time.Time).Format(time.RFC3339Nano)
	case bool:
		return strconv.FormatBool(value.(bool))
	default:
		return fmt.Sprintf("%+v", value)
	}
}

// FromString converts to string a value.
func FromString(value string, t ComparatorType) interface{} {
	switch t {
	case StringComparatorType:
		return value
	case IntComparatorType:
		typedKey, _ := strconv.ParseInt(value, 10, 64)
		return int(typedKey)
	case Int8ComparatorType:
		typedKey, _ := strconv.ParseInt(value, 10, 64)
		return int8(typedKey)
	case Int16ComparatorType:
		typedKey, _ := strconv.ParseInt(value, 10, 64)
		return int16(typedKey)
	case Int32ComparatorType:
		typedKey, _ := strconv.ParseInt(value, 10, 64)
		return int32(typedKey)
	case Int64ComparatorType:
		typedKey, _ := strconv.ParseInt(value, 10, 64)
		return int64(typedKey)
	case UIntComparatorType:
		typedKey, _ := strconv.ParseInt(value, 10, 64)
		return uint(typedKey)
	case UInt8ComparatorType:
		typedKey, _ := strconv.ParseInt(value, 10, 64)
		return uint8(typedKey)
	case UInt16ComparatorType:
		typedKey, _ := strconv.ParseInt(value, 10, 64)
		return uint16(typedKey)
	case UInt32ComparatorType:
		typedKey, _ := strconv.ParseInt(value, 10, 64)
		return uint32(typedKey)
	case UInt64ComparatorType:
		typedKey, _ := strconv.ParseInt(value, 10, 64)
		return uint64(typedKey)
	case Float32ComparatorType:
		typedKey, _ := strconv.ParseFloat(value, 64)
		return float32(typedKey)
	case Float64ComparatorType:
		typedKey, _ := strconv.ParseFloat(value, 64)
		return float64(typedKey)
	case ByteComparatorType:
		return []byte(value)
	case TimeComparatorType:
		t, _ := time.Parse(time.RFC3339Nano, value)
		return t
	default:
		return value
	}
}
