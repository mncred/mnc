package model

import (
	"encoding/json"
	"reflect"
)

// BasicOrComplex describes basic data type of complex like structs and arrays.
// First we try to unmarshal data as basic, then as complex.
type BasicOrComplex[B BasicType, C any] Any

// UnmarshalJSON implements json.Unmarshaler interface.
func (s *BasicOrComplex[B, C]) UnmarshalJSON(data []byte) error {
	*s = BasicOrComplex[B, C](data)
	return nil
}

// MarshalJSON implements json.Marshaler interface.
func (s BasicOrComplex[B, C]) MarshalJSON() ([]byte, error) {
	return s, nil
}

// IsBasic checks if data is basic.
func (s BasicOrComplex[B, C]) IsBasic() bool {
	var v *B
	if err := json.Unmarshal(s, &v); err != nil {
		return false
	}
	return v != nil
}

// IsArray checks if data is array.
func (s BasicOrComplex[B, C]) IsArray() bool {
	var v []json.RawMessage
	if err := json.Unmarshal(s, &v); err != nil {
		return false
	}
	return v != nil
}

// IsStruct checks if data is struct.
func (s BasicOrComplex[B, C]) IsStruct() bool {
	var v map[string]json.RawMessage
	if err := json.Unmarshal(s, &v); err != nil {
		return false
	}
	return v != nil
}

// IsComplex checks if data is complex, e.g. struct, map or array.
func (s BasicOrComplex[B, C]) IsComplex() bool {
	var v C
	if s.IsArray() {
		switch reflect.TypeOf(v).Kind() {
		case reflect.Slice, reflect.Array:
			return true
		}
		return false
	}
	if s.IsStruct() {
		switch reflect.TypeOf(v).Kind() {
		case reflect.Struct, reflect.Map:
			return true
		}
		return false
	}
	return false
}

// AsBasicE tries to fetch data as basic, returns error otherwise.
func (s BasicOrComplex[B, C]) AsBasicE() (B, error) {
	var b B
	if err := json.Unmarshal(s, &b); err != nil {
		return b, err
	}
	return b, nil
}

// AsBasic tries to fetch data as basic.
func (s BasicOrComplex[B, C]) AsBasic() B {
	b, _ := s.AsBasicE()
	return b
}

// AsComplexE tries to fetch data as complex, returns error otherwise.
func (s BasicOrComplex[B, C]) AsComplexE() (C, error) {
	var c C
	if err := json.Unmarshal(s, &c); err != nil {
		return c, err
	}
	return c, nil
}

// AsComplex tries to fetch data as complex.
func (s BasicOrComplex[B, C]) AsComplex() C {
	c, _ := s.AsComplexE()
	return c
}
