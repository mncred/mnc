// package hasher provides ability to calc multi-hash.
package hasher

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"io"
)

// Calc accepts reader and calculates multi-hash.
func Calc(rd io.Reader) (Hash, error) {
	hSha512 := sha512.New()
	hSha256 := sha256.New()
	hSha1 := sha1.New()
	hMd5 := md5.New()

	mwr := io.MultiWriter(hSha512, hSha256, hSha1, hMd5)
	if _, err := io.Copy(mwr, rd); err != nil {
		return Hash{}, err
	}

	return Hash{
		Sha512: hex.EncodeToString(hSha512.Sum(nil)),
		Sha256: hex.EncodeToString(hSha256.Sum(nil)),
		Sha1:   hex.EncodeToString(hSha1.Sum(nil)),
		Md5:    hex.EncodeToString(hMd5.Sum(nil)),
	}, nil
}

// Hash provides multi-hash info.
type Hash struct {
	Sha512 string `json:"sha512"`
	Sha256 string `json:"sha256"`
	Sha1   string `json:"sha1"`
	Md5    string `json:"md5"`
}
