package utils

import (
	"bytes"

	"github.com/opencontainers/go-digest"
)

// GetDigestJSONFilename generates a filename with a json
// extension that is a hash of the path passed in
func GetDigestJSONFilename(path string) (string, error) {
	hash, err := digest.Parse(path)
	if err != nil {
		return "", err
	}

	str := hash.String()
	var buffer bytes.Buffer
	buffer.WriteString(str)
	buffer.WriteString(".json")

	return buffer.String(), nil
}
