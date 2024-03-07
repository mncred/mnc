package describer

import (
	"archive/zip"
	"errors"
	"io"
	"io/fs"
	"os"

	"github.com/mncred/mnc/internal/describer/models"
	"github.com/mncred/mnc/internal/describer/zhelper"
)

// FabricInfo contains info about fabric mod file.
type FabricInfo struct {
	// Generic models.GenericModInfo `json:"generic"`
	Sources struct {
		FabricModJson models.FabricModJsonV1 `json:"fabric.mod.json"`
	} `json:"sources"`
}

// DescribeFabric tries to describe fabric mod.
// Returns null without errors if no fabric signatures found.
func DescribeFabric(file *os.File) (*FabricInfo, error) {
	info := FabricInfo{}

	size, err := file.Seek(0, io.SeekEnd)
	if err != nil {
		return nil, err
	}

	// try open file as archive
	zFile, err := zip.NewReader(file, size)
	if err != nil {
		// return nothing if it is not archive - no fabric signatures
		if errors.Is(err, zip.ErrFormat) {
			return nil, nil
		}
		return nil, err
	}

	// try find and load fabric.mod.json
	var fabricModJsonV1 models.FabricModJsonV1
	if err := zhelper.DeserializeJSON(zFile, &fabricModJsonV1, "fabric.mod.json"); err != nil {
		// return nothing if it is not fabric mod file - no fabric signatures
		if errors.Is(err, fs.ErrNotExist) {
			return nil, nil
		}
		return nil, err
	}
	info.Sources.FabricModJson = fabricModJsonV1

	return &info, nil
}
