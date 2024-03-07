package model

import "encoding/json"

// SingleOrMany describes single or many data structures.
// It's useful when the field value might be as a single value
// or as an array of values.
type SingleOrMany[T any] Any

// UnmarshalJSON implements json.Unmarshaler interface.
func (s *SingleOrMany[T]) UnmarshalJSON(data []byte) error {
	*s = SingleOrMany[T](data)
	return nil
}

// MarshalJSON implements json.Marshaler interface.
func (s SingleOrMany[T]) MarshalJSON() ([]byte, error) {
	return s, nil
}

// IsMany checks if data is array.
func (s SingleOrMany[T]) IsMany() bool {
	return Any(s).IsArray()
}

// IsSingle checks if data is single.
func (s SingleOrMany[T]) IsSingle() bool {
	return !s.IsMany()
}

// ValuesE returns array of values.
func (s SingleOrMany[T]) ValuesE() ([]T, error) {
	if s.IsSingle() {
		var v T
		return []T{v}, json.Unmarshal(s, &v)
	}
	var v []T
	return v, json.Unmarshal(s, &v)
}

// Values returns array of values.
func (s SingleOrMany[T]) Values() []T {
	v, _ := s.ValuesE()
	return v
}
