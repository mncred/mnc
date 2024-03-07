// package zhelper allows read zipped files.
package zhelper

import (
	"archive/zip"
	"io"
)

// TOMLDeserializer describes interface for deserializing TOML file.
type TOMLDeserializer interface {
	DeserializeTOML(data []byte) error
}

// JSONDeserializer describes interface for deserializing JSON file.
type JSONDeserializer interface {
	DeserializeJSON(data []byte) error
}

// DeserializeTOML unmarshals TOML file.
func DeserializeTOML(zrd *zip.Reader, v TOMLDeserializer, filename string) error {
	file, err := zrd.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	if err := v.DeserializeTOML(data); err != nil {
		return err
	}
	return nil
}

// DeserializeJSON unmarshals JSON file.
func DeserializeJSON(zrd *zip.Reader, v JSONDeserializer, filename string) error {
	file, err := zrd.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	if err := v.DeserializeJSON(data); err != nil {
		return err
	}
	return nil
}
