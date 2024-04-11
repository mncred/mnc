// package describe provides info about any minecraft related file.
// Generally it designed to identify .jar files of mods and plugins.
package describer

import (
	"archive/zip"
	"io"
	"os"

	"github.com/mncred/mnc/internal/describer/models"
)

// Info describes general file info.
type Info struct {
	General *GeneralInfo `json:"general"`
	Loaders struct {
		Fabric   *FabricInfo   `json:"fabric"`
		Forge    *ForgeInfo    `json:"forge"`
		NeoForge *NeoForgeInfo `json:"neoForge"`
	} `json:"loaders"`
}

// Describe accepts any file and tries to determine what it is.
// It will modify file's cursor.
func Describe(file *os.File) (*Info, error) {
	var err error
	info := Info{}

	info.General, err = DescribeGeneral(file)
	if err != nil {
		return nil, err
	}

	info.Loaders.Fabric, err = DescribeFabric(file)
	if err != nil {
		return nil, err
	}

	info.Loaders.Forge, err = DescribeForge(file)
	if err != nil {
		return nil, err
	}

	info.Loaders.NeoForge, err = DescribeNeoForge(file)
	if err != nil {
		return nil, err
	}

	return &info, nil
}

// ModInfo provides unified description for any mod from loaders.
type ModInfo struct {
	ModId     string `json:"id"`
	Version   string `json:"version"`
	Author    string `json:"author"`
	Title     string `json:"title"`
	Icon      string `json:"icon"`
	McVersion string `json:"mcVersion"`
	// Dependencies - list of dependency ID and version range, like "forge>=42.0.0"
	Dependencies map[string]string `json:"dependencies"`
	// Provides - list of provides ID and version, like: "voice-chat>=1.19.2"
	Provides map[string]string `json:"provides"`
}

// UnmarshalZip unmarshals any available model structure from zip file.
func UnmarshalZip(zrd *zip.Reader, v models.Model, filename string) error {
	file, err := zrd.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	if err := v.Unmarshal(data); err != nil {
		return err
	}
	return nil
}
