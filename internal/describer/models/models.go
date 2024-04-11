package models

// Model provides methods to read file format and serialize it as JSON.
type Model interface {
	Unmarshal(data []byte) error
	AsJSON() ([]byte, error)
}
