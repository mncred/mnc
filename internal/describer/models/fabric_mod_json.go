package models

import (
	"encoding/json"

	"github.com/mncred/mnc/internal/model"
)

// FabricModJsonV1 describes fabric.mod.json
// https://fabricmc.net/wiki/documentation:fabric_mod_json
type FabricModJsonV1 struct {
	// SchemaVersion - Needed for internal mechanisms. Must always be 1.
	SchemaVersion int `json:"schemaVersion"`
	// Id - Defines the mod's identifier - a string of Latin letters, digits, underscores with length from 2 to 64.
	Id string `json:"id"`
	// Version - Defines the mod's version - a string value, optionally matching the
	// Semantic Versioning 2.0.0 specification.
	Version string `json:"version"`
	// Provides - Defines the list of ids of mod. It can be seen as the aliases of the mod.
	// Fabric Loader will treat these ids as mods that exist. If there are other mods using that id,
	// they will not be loaded.
	Provides *[]string `json:"provides"`
	// Environment - Defines where mod runs: only on the client side (client mod), only on the server side (plugin)
	// or on both sides (regular mod).
	// Can be: "*", "client", "server"
	Environment *string `json:"environment"`
	// Entrypoints - Defines main classes of your mod, that will be loaded.
	// There are 3 default entry points for your mod: main, client, server.
	Entrypoints *map[string][]string `json:"entrypoints"`
	// Jars - A list of nested JARs inside your mod's JAR to load. Before using the field, check out the guidelines
	// on the usage of the nested JARs. Each entry is an object containing file key. That should be a path inside
	// your mod's JAR to the nested JAR
	Jars *[]struct {
		File string `json:"file"`
	} `json:"jars"`
	// languageAdapters - A dictionary of adapters for used languages to their adapter classes full names.
	LanguageAdapters *map[string]string `json:"languageAdapters"`
	// Mixins - A list of mixin configuration files. Each entry is the path to the mixin configuration file inside
	// your mod's JAR or an object containing following fields: config, environment
	Mixins *[]model.BasicOrComplex[string, struct {
		Config      string `json:"config"`
		Environment string `json:"environment"`
	}] `json:"mixins"`
	// Depends - For dependencies required to run. Without them a game will crash.
	Depends *map[string]string `json:"depends"`
	// Recommends - For dependencies not required to run. Without them a game will log a warning.
	Recommends *map[string]string `json:"recommends"`
	// Suggests - For dependencies not required to run. Use this as a kind of metadata.
	Suggests *map[string]string `json:"suggests"`
	// Breaks - For mods whose together with yours might cause a game crash. With them a game will crash.
	Breaks *map[string]string `json:"breaks"`
	// Conflicts - For mods whose together with yours cause some kind of bugs, etc. With them a game will log a warning.
	Conflicts *map[string]string `json:"conflicts"`
	// Name - Defines the user-friendly mod's name. If not present, assume it matches id.
	Name *string `json:"name"`
	// Description - Defines the mod's description. If not present, assume empty string.
	Description *string `json:"description"`
	// Contact - Defines the contact information for the project. It is an object of the following fields
	Contact *map[string]string `json:"contact"`
	// Authors - A list of authors of the mod. Each entry is a single name or an object.
	Authors *[]model.BasicOrComplex[string, struct {
		Name    string            `json:"name"`
		Contact map[string]string `json:"contact"`
	}] `json:"authors"`
	// Contributors - A list of contributors to the mod. Each entry is the same as in author field.
	Contributors *[]model.BasicOrComplex[string, struct {
		Name    string            `json:"name"`
		Contact map[string]string `json:"contact"`
	}] `json:"contributors"`
	// License - Defines the licensing information. Can either be a single license string or a list of them.
	License *model.SingleOrMany[string] `json:"license"`
	// Icon - Defines the mod's icon. Can be provided in one of two forms:
	// A path to a single PNG file.
	// A dictionary of images widths to their files' paths.
	Icon *model.BasicOrComplex[string, map[string]string] `json:"icon"`
}

// DeserializeJSON unmarshals file.
func (f *FabricModJsonV1) DeserializeJSON(data []byte) error {
	return json.Unmarshal(data, f)
}

// SerializeJSON marshals file.
func (f FabricModJsonV1) SerializeJSON() ([]byte, error) {
	return json.MarshalIndent(f, "", "  ")
}
