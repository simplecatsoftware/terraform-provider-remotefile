package types

import (
	"io"
	"net/http"
	"os"
)

type Http struct {
	*Type
	Uri       string
	LocalFile Temporary
}

func (h Http) GetProtocols() []string {
	return []string{"http", "https"}
}

func (h Http) GetUri() string {
	return h.Uri
}

func (h Http) Read() ([]byte, error) {
	data := []byte{}

	resp, err := http.Get(h.GetUri())
	if err != nil {
		return data, err
	}

	out, err := os.Create(h.GetFilePath())
	if err != nil {
		return data, err
	}

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return data, err
	}

	err = out.Close()
	if err != nil {
		return data, err
	}

	err = resp.Body.Close()
	if err != nil {
		return data, err
	}

	return data, nil
}

func (h Http) Validate() error {
	return h.validateUri(h)
}

func (h Http) Sha256() (string, error) {
	return h.LocalFile.Sha256()
}

func (h Http) GetFileName() string {
	return h.getFileName(h)
}

func (h Http) GetFilePath() string {
	return h.LocalFile.GetFilePath()
}
