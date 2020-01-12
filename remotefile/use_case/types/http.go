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
	HasRead   bool
}

func (h Http) GetProtocols() []string {
	return []string{"http", "https"}
}

func (h Http) GetUri() string {
	return h.Uri
}

func (h Http) Read() ([]byte, error) {
	if !h.HasRead {
		resp, err := http.Get(h.GetUri())
		if err != nil {
			return []byte{}, err
		}

		out, err := os.Create(h.GetFilePath())
		if err != nil {
			return []byte{}, err
		}

		_, err = io.Copy(out, resp.Body)
		if err != nil {
			return []byte{}, err
		}

		err = out.Close()
		if err != nil {
			return []byte{}, err
		}

		err = resp.Body.Close()
		if err != nil {
			return []byte{}, err
		}

		h.HasRead = true
	}

	return h.LocalFile.Read()
}

func (h Http) Validate() error {
	err := h.validateUri(h)
	if err != nil {
		return err
	}

	_, err = h.Read()
	if err != nil {
		return err
	}

	return nil
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
