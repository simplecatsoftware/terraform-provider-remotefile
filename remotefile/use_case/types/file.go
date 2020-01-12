package types

import (
	"io/ioutil"
	"strings"
)

type File struct {
	*Type
	Uri string
}

func (f File) GetProtocols() []string {
	return []string{"file"}
}

func (f File) GetUri() string {
	return f.Uri
}

func (f File) Read() ([]byte, error) {
	path := f.filePath()

	return ioutil.ReadFile(path)
}

func (f File) Write(content []byte) error {
	err := ioutil.WriteFile(f.filePath(), content, 0644)
	if err != nil {
		return err
	}

	return nil
}

func (f File) Validate() error {
	return f.validateUri(f)
}

func (f File) Sha256() (string, error) {
	return f.sha256(f)
}

func (f File) GetFileName() string {
	return f.getFileName(f)
}

func (f File) filePath() string {
	return strings.Replace(f.GetUri(), "file://", "", -1)
}
