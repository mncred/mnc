// package describe provides info about any minecraft related file.
// Generally it designed to identify .jar files of mods and plugins.
package describer

import (
	"os"
)

// Describe accepts any file and tries to determine what it is.
// It will modify file's cursor.
func Describe(file *os.File) (*Info, error) {
	var err error
	info := Info{}

	info.General, err = DescribeGeneral(file)
	if err != nil {
		return nil, err
	}

	info.Loader.Fabric, err = DescribeFabric(file)
	if err != nil {
		return nil, err
	}

	info.Loader.Forge, err = DescribeForge(file)
	if err != nil {
		return nil, err
	}

	return &info, nil
}

// Info describes general file info.
type Info struct {
	General *GeneralInfo `json:"general"`
	Loader  struct {
		Fabric *FabricInfo `json:"fabric"`
		Forge  *ForgeInfo  `json:"forge"`
	} `json:"loader"`
}
