// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package linkedhashmap

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
	"unicode/utf8"
)

func TestFromJSONPreservesObjectOrder(t *testing.T) {
	tests := []struct {
		name, data string
		keys       []string
	}{
		{"value matches later key", `{"a":"c","b":"x","c":"y"}`, []string{"a", "b", "c"}},
		{"nested key matches later key", `{"z":{"b":1},"a":2,"b":3}`, []string{"z", "a", "b"}},
		{"escaped key", `{"z":1,"\u0061":2,"b":3}`, []string{"z", "a", "b"}},
		{"escaped less than", `{"z":1,"<":2,"a":3}`, []string{"z", "<", "a"}},
		{"duplicate keys", `{"a":1,"b":2,"a":3}`, []string{"a", "b"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, unmarshal := range []bool{false, true} {
				m := New[string, any]()
				m.Put("old", false)
				var err error
				if unmarshal {
					err = json.Unmarshal([]byte(tt.data), m)
				} else {
					err = m.FromJSON([]byte(tt.data))
				}
				if err != nil {
					t.Fatal(err)
				}
				if got := m.Keys(); !reflect.DeepEqual(got, tt.keys) {
					t.Fatalf("keys = %v, want %v", got, tt.keys)
				}
			}
		})
	}
}

func TestFromJSONOrderValuesAndErrors(t *testing.T) {
	m := New[string, map[string]int]()
	if err := m.FromJSON([]byte(`{"z":{"a":1},"b":{"b":2},"z":{"c":3}}`)); err != nil {
		t.Fatal(err)
	}
	if got := m.Keys(); !reflect.DeepEqual(got, []string{"z", "b"}) {
		t.Fatal(got)
	}
	got, _ := m.Get("z")
	if !reflect.DeepEqual(got, map[string]int{"c": 3}) {
		t.Fatalf("duplicate value = %v", got)
	}
	for _, bad := range []string{`[]`, `{"a":`, `{"a":1} {"b":2}`, `{"a":true}`} {
		before := m.Keys()
		if err := m.FromJSON([]byte(bad)); err == nil {
			t.Fatalf("accepted %s", bad)
		}
		if !reflect.DeepEqual(m.Keys(), before) {
			t.Fatalf("error changed keys for %s", bad)
		}
	}
	for _, empty := range []string{`{}`, `null`} {
		if err := m.FromJSON([]byte(empty)); err != nil {
			t.Fatal(err)
		}
		if !m.Empty() {
			t.Fatalf("%s did not clear map", empty)
		}
	}
}

type jsonOrderKey int

func (k *jsonOrderKey) UnmarshalText(text []byte) error {
	switch string(text) {
	case "first":
		*k = 1
	case "second":
		*k = 2
	default:
		*k = 3
	}
	return nil
}

func TestFromJSONNonStringKeyOrder(t *testing.T) {
	m := New[int, string]()
	if err := m.FromJSON([]byte(`{"30":"3","2":"2","10":"1","02":"updated"}`)); err != nil {
		t.Fatal(err)
	}
	if got := m.Keys(); !reflect.DeepEqual(got, []int{30, 2, 10}) {
		t.Fatal(got)
	}
	if got, _ := m.Get(2); got != "updated" {
		t.Fatal(got)
	}
	custom := New[jsonOrderKey, int]()
	if err := custom.FromJSON([]byte(`{"second":2,"first":1}`)); err != nil {
		t.Fatal(err)
	}
	if got := custom.Keys(); !reflect.DeepEqual(got, []jsonOrderKey{2, 1}) {
		t.Fatal(got)
	}
}

func TestFromJSONOrderLargeObject(t *testing.T) {
	original := New[string, string]()
	for i := 0; i < 200; i++ {
		original.Put(fmt.Sprintf("key%03d", i), fmt.Sprintf("key%03d", 199-i))
	}
	data, err := original.ToJSON()
	if err != nil {
		t.Fatal(err)
	}
	decoded := New[string, string]()
	if err := decoded.FromJSON(data); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(decoded.Keys(), original.Keys()) {
		t.Fatal("round trip changed order")
	}
}

func BenchmarkFromJSONOrdered(b *testing.B) {
	for _, size := range []int{10, 100, 1000} {
		b.Run(fmt.Sprint(size), func(b *testing.B) {
			original := New[string, int]()
			for i := 0; i < size; i++ {
				original.Put(fmt.Sprint(i), i)
			}
			data, _ := original.ToJSON()
			m := New[string, int]()
			b.ReportAllocs()
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				if err := m.FromJSON(data); err != nil {
					b.Fatal(err)
				}
			}
		})
	}
}

func FuzzFromJSONKeyOrder(f *testing.F) {
	f.Add("a", "b", "c")
	f.Add("z", "<", "a")
	f.Fuzz(func(t *testing.T, a, b, c string) {
		if !utf8.ValidString(a) || !utf8.ValidString(b) || !utf8.ValidString(c) {
			return
		}
		original := New[string, string]()
		original.Put(a, c)
		original.Put(b, a)
		original.Put(c, b)
		data, err := original.ToJSON()
		if err != nil {
			t.Fatal(err)
		}
		decoded := New[string, string]()
		if err := decoded.FromJSON(data); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(decoded.Keys(), original.Keys()) || !reflect.DeepEqual(decoded.Values(), original.Values()) {
			t.Fatalf("round trip changed map: %q", data)
		}
	})
}
