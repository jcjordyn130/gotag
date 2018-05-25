// Here lives the decoders for gotag, this is all the format specific stuff.
package decoders

import (
	"os"
	"crypto/sha256"
	"io"
	"encoding/hex"
	"github.com/jcjordyn130/gotag/types"
)

// Sum returns the sha256 sum of a track.
// You can either use this or use metadata.Sum().
func Sum(m types.Metadata) (string, error) {
	// Open the file.
	file, err := os.Open(m.File())
	if err != nil {
		return "", err
	}

	// Defer closing it.
	defer file.Close()

	// Make a new hash object.
	hash := sha256.New()

	// Copy the file to the hasher object.
	_, err = io.Copy(hash, file)
	if err != nil {
		return "", err
	}

	// Calculate the hash and return it.
	return hex.EncodeToString(hash.Sum(nil)), err
}
