// package model provides models helper functions.
package model

// BasicType describes basic data type.
type BasicType interface {
	~int64 | ~uint64 | ~float64 | ~string | ~bool
}
