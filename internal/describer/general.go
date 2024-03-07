package describer

import (
	"io"
	"os"

	"github.com/mncred/mnc/internal/hasher"
)

// GeneralInfo contains general file info.
type GeneralInfo struct {
	// Size of the file.
	Size int64 `json:"size"`
	// Hash provides multi-hash info.
	Hash hasher.Hash `json:"hash"`
}

// DescribeGeneral tries to describe general file info.
func DescribeGeneral(file *os.File) (*GeneralInfo, error) {
	info := GeneralInfo{}

	// calculate file hashes
	file.Seek(0, io.SeekStart)
	if hash, err := hasher.Calc(file); err != nil {
		return nil, err
	} else {
		info.Hash = hash
	}

	// calculate file size
	if size, err := file.Seek(0, io.SeekEnd); err != nil {
		return nil, err
	} else {
		info.Size = size
	}

	return &info, nil
}
