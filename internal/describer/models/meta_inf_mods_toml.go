package models

import (
	"encoding/json"

	"github.com/pelletier/go-toml/v2"
)

// MetaInfModsToml describes META-INF/mods.toml file.
type MetaInfModsTomlV1 struct {
	ModLoader          string `toml:"modLoader" json:"modLoader"`
	LoaderVersion      string `toml:"loaderVersion" json:"loaderVersion"`
	License            string `toml:"license" json:"license"`
	IssueTrackerUrl    string `toml:"issueTrackerURL" json:"issueTrackerUrl"`
	ShowAsResourcePack *bool  `toml:"showAsResourcePack" json:"showAsResourcePack"`
	Mods               []struct {
		ModId         string  `toml:"modId" json:"modId"`
		Version       string  `toml:"version" json:"version"`
		DisplayName   string  `toml:"displayName" json:"displayName"`
		UpdateJsonUrl *string `toml:"updateJSONURL" json:"updateJsonUrl"`
		DisplayUrl    *string `toml:"displayURL" json:"displayUrl"`
		LogoFile      *string `toml:"logoFile" json:"logoFile"`
		Credits       *string `toml:"credits" json:"credits"`
		Authors       *string `toml:"authors" json:"authors"`
		Description   *string `toml:"description" json:"description"`
		DisplayTest   *string `toml:"displayTest" json:"displayTest"`
	} `toml:"mods" json:"mods"`
	Dependencies map[string][]struct {
		ModId        string `toml:"modId" json:"modId"`
		Mandatory    bool   `toml:"mandatory" json:"mandatory"`
		VersionRange string `toml:"versionRange" json:"versionRange"`
		Ordering     string `toml:"ordering" json:"ordering"`
		Side         string `toml:"side" json:"side"`
	} `toml:"dependencies" json:"dependencies"`
}

// DeserializeTOML unmarshals file.
func (f *MetaInfModsTomlV1) Unmarshal(data []byte) error {
	return toml.Unmarshal(data, f)
}

// SerializeJSON marshals file.
func (f MetaInfModsTomlV1) AsJSON() ([]byte, error) {
	return json.MarshalIndent(f, "", "  ")
}
