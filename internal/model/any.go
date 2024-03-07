package model

import "encoding/json"

// Any describes any available data structure.
type Any json.RawMessage

// UnmarshalJSON implements json.Unmarshaler interface.
func (a *Any) UnmarshalJSON(data []byte) error {
	*a = Any(data)
	return nil
}

// MarshalJSON implements json.Marshaler interface.
func (a Any) MarshalJSON() ([]byte, error) {
	return a, nil
}

// AsE tries to fetch data, returns error otherwise.
func (a Any) AsE(v any) error {
	if err := json.Unmarshal(a, v); err != nil {
		return err
	}
	return nil
}

// As tries to fetch data.
func (a Any) As(v any) {
	a.AsE(v)
}

// Is checks if it's able to fetch data.
func (a Any) Is(v any) bool {
	return a.AsE(v) == nil
}

// AsStringE tries to fetch data as string, returns error otherwise.
func (a Any) AsStringE() (string, error) {
	var str string
	if err := a.AsE(&str); err != nil {
		return "", err
	}
	return str, nil
}

// AsString tries to fetch data as string.
func (a Any) AsString() string {
	str, _ := a.AsStringE()
	return str
}

// IsString checks if data is string.
func (a Any) IsString() bool {
	_, err := a.AsStringE()
	return err == nil
}

// AsFloat64E tries to fetch data as float64, returns error otherwise.
func (a Any) AsFloat64E() (float64, error) {
	var f float64
	if err := json.Unmarshal(a, &f); err != nil {
		return 0, err
	}
	return f, nil
}

// AsFloat64 tries to fetch data as float64.
func (a Any) AsFloat64() float64 {
	f, _ := a.AsFloat64E()
	return f
}

// IsFloat64 checks if data is float64.
func (a Any) IsFloat64() bool {
	_, err := a.AsFloat64E()
	return err == nil
}

// AsFloat32E tries to fetch data as float32, returns error otherwise.
func (a Any) AsFloat32E() (float32, error) {
	var f float32
	if err := json.Unmarshal(a, &f); err != nil {
		return 0, err
	}
	return f, nil
}

// AsFloat32 tries to fetch data as float32.
func (a Any) AsFloat32() float32 {
	f, _ := a.AsFloat32E()
	return f
}

// IsFloat32 checks if data is float32.
func (a Any) IsFloat32() bool {
	_, err := a.AsFloat32E()
	return err == nil
}

// AsInt64E tries to fetch data as int64, returns error otherwise.
func (a Any) AsInt64E() (int64, error) {
	var i int64
	if err := json.Unmarshal(a, &i); err != nil {
		return 0, err
	}
	return i, nil
}

// AsInt64 tries to fetch data as int64.
func (a Any) AsInt64() int64 {
	i, _ := a.AsInt64E()
	return i
}

// IsInt64 checks if data is int64.
func (a Any) IsInt64() bool {
	_, err := a.AsInt64E()
	return err == nil
}

// AsUint64E tries to fetch data as uint64, returns error otherwise.
func (a Any) AsUint64E() (uint64, error) {
	var i uint64
	if err := json.Unmarshal(a, &i); err != nil {
		return 0, err
	}
	return i, nil
}

// AsUint64 tries to fetch data as uint64.
func (a Any) AsUint64() uint64 {
	i, _ := a.AsUint64E()
	return i
}

// IsUint64 checks if data is uint64.
func (a Any) IsUint64() bool {
	_, err := a.AsUint64E()
	return err == nil
}

// AsInt32E tries to fetch data as int32, returns error otherwise.
func (a Any) AsInt32E() (int32, error) {
	var i int32
	if err := json.Unmarshal(a, &i); err != nil {
		return 0, err
	}
	return i, nil
}

// AsInt32 tries to fetch data as int32.
func (a Any) AsInt32() int32 {
	i, _ := a.AsInt32E()
	return i
}

// IsInt32 checks if data is int32.
func (a Any) IsInt32() bool {
	_, err := a.AsInt32E()
	return err == nil
}

// AsUint32E tries to fetch data as uint32, returns error otherwise.
func (a Any) AsUint32E() (uint32, error) {
	var i uint32
	if err := json.Unmarshal(a, &i); err != nil {
		return 0, err
	}
	return i, nil
}

// AsUint32 tries to fetch data as uint32.
func (a Any) AsUint32() uint32 {
	i, _ := a.AsUint32E()
	return i
}

// IsUint32 checks if data is uint32.
func (a Any) IsUint32() bool {
	_, err := a.AsUint32E()
	return err == nil
}

// AsInt16E tries to fetch data as int16, returns error otherwise.
func (a Any) AsInt16E() (int16, error) {
	var i int16
	if err := json.Unmarshal(a, &i); err != nil {
		return 0, err
	}
	return i, nil
}

// AsInt16 tries to fetch data as int16.
func (a Any) AsInt16() int16 {
	i, _ := a.AsInt16E()
	return i
}

// IsInt16 checks if data is int16.
func (a Any) IsInt16() bool {
	_, err := a.AsInt16E()
	return err == nil
}

// AsUint16E tries to fetch data as uint16, returns error otherwise.
func (a Any) AsUint16E() (uint16, error) {
	var i uint16
	if err := json.Unmarshal(a, &i); err != nil {
		return 0, err
	}
	return i, nil
}

// AsUint16 tries to fetch data as uint16.
func (a Any) AsUint16() uint16 {
	i, _ := a.AsUint16E()
	return i
}

// IsUint16 checks if data is uint16.
func (a Any) IsUint16() bool {
	_, err := a.AsUint16E()
	return err == nil
}

// AsInt8E tries to fetch data as int8, returns error otherwise.
func (a Any) AsInt8E() (int8, error) {
	var i int8
	if err := json.Unmarshal(a, &i); err != nil {
		return 0, err
	}
	return i, nil
}

// AsInt8 tries to fetch data as int8.
func (a Any) AsInt8() int8 {
	i, _ := a.AsInt8E()
	return i
}

// IsInt8 checks if data is int8.
func (a Any) IsInt8() bool {
	_, err := a.AsInt8E()
	return err == nil
}

// AsUint8E tries to fetch data as uint8, returns error otherwise.
func (a Any) AsUint8E() (uint8, error) {
	var i uint8
	if err := json.Unmarshal(a, &i); err != nil {
		return 0, err
	}
	return i, nil
}

// AsUint8 tries to fetch data as uint8.
func (a Any) AsUint8() uint8 {
	i, _ := a.AsUint8E()
	return i
}

// IsUint8 checks if data is uint8.
func (a Any) IsUint8() bool {
	_, err := a.AsUint8E()
	return err == nil
}

// AsIntE tries to fetch data as int, returns error otherwise.
func (a Any) AsIntE() (int, error) {
	var i int
	if err := json.Unmarshal(a, &i); err != nil {
		return 0, err
	}
	return i, nil
}

// AsInt tries to fetch data as int.
func (a Any) AsInt() int {
	i, _ := a.AsIntE()
	return i
}

// IsInt checks if data is int.
func (a Any) IsInt() bool {
	_, err := a.AsIntE()
	return err == nil
}

// AsUintE tries to fetch data as uint, returns error otherwise.
func (a Any) AsUintE() (uint, error) {
	var i uint
	if err := json.Unmarshal(a, &i); err != nil {
		return 0, err
	}
	return i, nil
}

// AsUint tries to fetch data as uint.
func (a Any) AsUint() uint {
	i, _ := a.AsUintE()
	return i
}

// IsUint checks if data is uint.
func (a Any) IsUint() bool {
	_, err := a.AsUintE()
	return err == nil
}

// AsBoolE tries to fetch data as bool, returns error otherwise.
func (a Any) AsBoolE() (bool, error) {
	var b bool
	if err := json.Unmarshal(a, &b); err != nil {
		return false, err
	}
	return b, nil
}

// AsBool tries to fetch data as bool.
func (a Any) AsBool() bool {
	b, _ := a.AsBoolE()
	return b
}

// IsBool checks if data is bool.
func (a Any) IsBool() bool {
	_, err := a.AsBoolE()
	return err == nil
}

// IsNil checks if data is nil.
func (a Any) IsNil() bool {
	var v *any
	json.Unmarshal(a, &v)
	return v == nil
}

// IsArray checks if data is array.
func (a Any) IsArray() bool {
	var v []json.RawMessage
	json.Unmarshal(a, &v)
	return v != nil
}

// IsStruct checks if data is struct.
func (a Any) IsStruct() bool {
	var v map[string]json.RawMessage
	json.Unmarshal(a, &v)
	return v != nil
}
