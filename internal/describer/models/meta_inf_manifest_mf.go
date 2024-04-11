package models

import (
	"encoding/json"

	"gopkg.in/yaml.v3"
)

// MetInfManifestMf describes META-INF/MANIFEST.MF file.
// Spec: https://docs.oracle.com/javase/8/docs/technotes/guides/jar/jar.html#JARManifest
// TODO: Implement full Java manifest parsing, including repeating fields like Sealed etc.
type MetaInfManifestMfV1 struct {
	ManifestVersion         string `yaml:"Manifest-Version" json:"manifestVersion"`
	CreatedBy               string `yaml:"Created-By" json:"createdBy"`
	MainClass               string `yaml:"Main-Class" json:"mainClass"`
	ClassPath               string `yaml:"Class-Path" json:"classPath"`
	SpecificationTitle      string `yaml:"Specification-Title" json:"specificationTitle"`
	SpecificationVersion    string `yaml:"Specification-Version" json:"specificationVersion"`
	SpecificationVendor     string `yaml:"Specification-Vendor" json:"specificationVendor"`
	ImplementationTitle     string `yaml:"Implementation-Title" json:"implementationTitle"`
	ImplementationVendor    string `yaml:"Implementation-Vendor" json:"implementationVendor"`
	ImplementationVersion   string `yaml:"Implementation-Version" json:"implementationVersion"`
	ImplementationTimestamp string `yaml:"Implementation-Timestamp" json:"implementationTimestamp"`
	Sealed                  bool   `yaml:"Sealed" json:"sealed"`
}

// DeserializeTOML unmarshals file.
func (f *MetaInfManifestMfV1) Unmarshal(data []byte) error {
	return yaml.Unmarshal(data, f)
}

// SerializeJSON marshals file.
func (f MetaInfManifestMfV1) AsJSON() ([]byte, error) {
	return json.MarshalIndent(f, "", "  ")
}
