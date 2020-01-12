package types

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"strings"
)

type T interface {
	GetUri() string
	GetFilePath() string
	GetProtocols() []string
	Read() ([]byte, error)
	Sha256() (string, error)
	Write([]byte) error
	Validate() error
	GetFileName() string
}

type Type struct {
	T
	Uri string
}

func (t *Type) Write([]byte) error {
	return errors.New("write is not implemented")
}

func (t *Type) GetFileName([]byte) error {
	return errors.New("get file name is not implemented")
}

func (t *Type) validateUri(item T) error {
	if item.GetUri() == "" {
		return errors.New("uri is not set")
	}

	typeHasProtocol := false
	for _, protocol := range item.GetProtocols() {
		if strings.HasPrefix(item.GetUri(), fmt.Sprintf("%s://", protocol)) {
			typeHasProtocol = true
		}
	}

	if !typeHasProtocol {
		return errors.New(fmt.Sprintf("uri is not one of %s:// %s", strings.Join(item.GetProtocols(), ", "), item.GetUri()))
	}

	return nil
}

func (t *Type) sha256(item T) (string, error) {
	fileData, err := item.Read()
	if err != nil {
		return "", err
	}

	computedHash := sha256.Sum256(fileData)

	return hex.EncodeToString(computedHash[:]), nil
}

func (t *Type) getFileName(item T) string {
	splitUri := strings.Split(item.GetUri(), "/")
	filename := splitUri[len(splitUri)-1]

	return filename
}
