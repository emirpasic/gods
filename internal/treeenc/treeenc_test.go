package treeenc

import (
	"encoding/json"
	"fmt"
	"math"
	"strings"
	"testing"
)

type customType struct {
	value string
}

func (c customType) MarshalText() ([]byte, error) {
	return []byte(fmt.Sprintf("customType(%s)", c.value)), nil
}

func (c *customType) UnmarshalText(text []byte) error {
	value := strings.TrimPrefix(string(text), "customType(")
	value = strings.TrimSuffix(value, ")")
	c.value = value
	return nil
}

func TestKeyMarshaler(t *testing.T) {
	strMap := make(map[*KeyMarshaler[string]]string)
	strMap[&KeyMarshaler[string]{Key: "key"}] = "value"
	data, err := json.Marshal(strMap)
	if err != nil {
		t.Fatalf("Got error: %v", err)
	}
	expected := "{\"key\":\"value\"}"
	if string(data) != expected {
		t.Fatalf("Expected %q, got %q", expected, string(data))
	}

	intMap := make(map[*KeyMarshaler[int]]int)
	intMap[&KeyMarshaler[int]{Key: math.MinInt}] = math.MinInt
	data, err = json.Marshal(&intMap)
	if err != nil {
		t.Fatalf("Got error: %v", err)
	}
	expected = "{\"-9223372036854775808\":-9223372036854775808}"
	if string(data) != expected {
		t.Fatalf("Expected %q, got %q", expected, string(data))
	}

	uintMap := make(map[*KeyMarshaler[uint]]uint)
	uintMap[&KeyMarshaler[uint]{Key: math.MaxUint}] = math.MaxUint
	data, err = json.Marshal(&uintMap)
	if err != nil {
		t.Fatalf("Got error: %v", err)
	}
	expected = "{\"18446744073709551615\":18446744073709551615}"
	if string(data) != expected {
		t.Fatalf("Expected %q, got %q", expected, string(data))
	}

	customMap := make(map[*KeyMarshaler[customType]]string)
	customMap[&KeyMarshaler[customType]{Key: customType{value: "key"}}] = "value"
	data, err = json.Marshal(&customMap)
	if err != nil {
		t.Fatalf("Got error: %v", err)
	}
	expected = "{\"customType(key)\":\"value\"}"
	if string(data) != expected {
		t.Fatalf("Expected %q, got %q", expected, string(data))
	}
}

func TestKeyUnmarshaler(t *testing.T) {
	strMap := make(map[KeyUnmarshaler[string]]string)
	data := []byte("{\"key\":\"value\"}")
	if err := json.Unmarshal(data, &strMap); err != nil {
		t.Fatalf("Got error: %v", err)
	}
	if len(strMap) != 1 {
		t.Fatalf("Expected 1 key, got %d", len(strMap))
	}
	for k, v := range strMap {
		if *k.Key != "key" {
			t.Fatalf("Expected key %q, got %q", "key", *k.Key)
		}
		if v != "value" {
			t.Fatalf("Expected value %q, got %q", "value", v)
		}
	}

	intMap := make(map[KeyUnmarshaler[int]]int)
	data = []byte("{\"-9223372036854775808\":-9223372036854775808}")
	if err := json.Unmarshal(data, &intMap); err != nil {
		t.Fatalf("Got error: %v", err)
	}
	if len(intMap) != 1 {
		t.Fatalf("Expected 1 key, got %d", len(intMap))
	}
	for k, v := range intMap {
		if *k.Key != math.MinInt {
			t.Fatalf("Expected key %d, got %d", math.MinInt, *k.Key)
		}
		if v != math.MinInt {
			t.Fatalf("Expected value %d, got %d", math.MinInt, v)
		}
	}

	uintMap := make(map[KeyUnmarshaler[uint]]uint)
	data = []byte("{\"18446744073709551615\":18446744073709551615}")
	if err := json.Unmarshal(data, &uintMap); err != nil {
		t.Fatalf("Got error: %v", err)
	}
	if len(uintMap) != 1 {
		t.Fatalf("Expected 1 key, got %d", len(uintMap))
	}
	for k, v := range uintMap {
		if *k.Key != math.MaxUint {
			t.Fatalf("Expected key %d, got %d", uint(math.MaxUint), *k.Key)
		}
		if v != math.MaxUint {
			t.Fatalf("Expected value %d, got %d", uint(math.MaxUint), v)
		}
	}

	customMap := make(map[KeyUnmarshaler[customType]]string)
	data = []byte("{\"customType(key)\":\"value\"}")
	if err := json.Unmarshal(data, &customMap); err != nil {
		t.Fatalf("Got error: %v", err)
	}
	if len(customMap) != 1 {
		t.Fatalf("Expected 1 key, got %d", len(customMap))
	}
	for k, v := range customMap {
		if (*k.Key).value != "key" {
			t.Fatalf("Expected key %q, got %q", "key", (*k.Key).value)
		}
		if v != "value" {
			t.Fatalf("Expected value %q, got %q", "value", v)
		}
	}
}
