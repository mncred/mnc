package describer

import (
	"archive/zip"
	"errors"
	"io"
	"io/fs"
	"os"

	"github.com/mncred/mnc/internal/describer/models"
)

// ForgeInfo contains info about forge mod file.
type ForgeInfo struct {
	Sources struct {
		MetaInfModsToml models.MetaInfModsTomlV1 `json:"META-INF/mods.toml"`
	} `json:"sources"`
}

// DescribeForge tries to describe forge mod.
// Returns null without errors if no forge signatures found.
func DescribeForge(file *os.File) (*ForgeInfo, error) {
	info := ForgeInfo{}

	size, err := file.Seek(0, io.SeekEnd)
	if err != nil {
		return nil, err
	}

	// try open file as archive
	zrd, err := zip.NewReader(file, size)
	if err != nil {
		// return nothing if it is not archive - no forge signature
		if errors.Is(err, zip.ErrFormat) {
			return nil, nil
		}
		return nil, err
	}

	// try find and load META-INF/mods.toml
	var metaInfModsTomlV1 models.MetaInfModsTomlV1
	if err := UnmarshalZip(zrd, &metaInfModsTomlV1, "META-INF/mods.toml"); err != nil {
		// return nothing if it is not forge mod file - no forge signature
		if errors.Is(err, fs.ErrNotExist) {
			return nil, nil
		}
		return nil, err
	}
	info.Sources.MetaInfModsToml = metaInfModsTomlV1

	// check for dependencies, at least one of them must have dependency from forge
	forgeDepFound := false
	for _, mod := range info.Sources.MetaInfModsToml.Mods {
		for _, dep := range info.Sources.MetaInfModsToml.Dependencies[mod.ModId] {
			if dep.ModId == "forge" {
				forgeDepFound = true
			}
		}
	}
	// no forge signatures found
	if !forgeDepFound {
		return nil, nil
	}

	return &info, nil
}
