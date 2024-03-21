package models

// Project provides entire mnc.yml file.
type Project struct {
	// Schema provides version of the mnc.yml file.
	// It must be present to correctly parse mnc.yml file.
	Schema string `yaml:"$schema,omitempty"`
}
