package helpers

import (
	"crypto/sha256"
	"encoding/hex"
	"io/ioutil"
	"terraform-provider-remotefiles/remotefiles/fetch"
)

func HashSha256(localFile fetch.LocalFile) (hash string, err error) {
	fileData, err := ioutil.ReadFile(localFile.Path)

	if err != nil {
		return "", err
	}

	computedHash := sha256.Sum256(fileData)

	return hex.EncodeToString(computedHash[:]), nil
}
