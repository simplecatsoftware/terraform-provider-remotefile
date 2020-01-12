package types

import (
	"strings"
)

type Temporary struct {
	T
	*Type
	Uri  string
	File File
}

func (t Temporary) GetUri() string {
	return t.Uri
}

func (t Temporary) GetProtocols() []string {
	return []string{"tmp"}
}

func (t Temporary) Read() ([]byte, error) {
	return t.File.Read()
}

func (t Temporary) Write(content []byte) error {
	return t.File.Write(content)
}

func (t Temporary) Validate() error {
	return t.validateUri(t)
}

func (t Temporary) Sha256() (string, error) {
	return t.sha256(t.File)
}

func (t Temporary) GetFileName() string {
	return t.getFileName(t)
}

func (t Temporary) filePath() string {
	return strings.Replace(t.GetUri(), "tmp://", "", -1)
}
