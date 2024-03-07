package models

// GenericModInfo describes generic mod file.
type GenericModInfo struct {
	// Id - mod id.
	Id string `json:"id"`
	// Version - mod version.
	Version string `json:"version"`
	// Icon - base64 encoded icon.
	Icon string `json:"icon"`
}
