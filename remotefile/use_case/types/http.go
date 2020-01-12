package types

import (
	"net/http"
)

type Http struct {
	*Type
	Uri string
}

func (h Http) GetProtocols() []string {
	return []string{"http", "https"}
}

func (h Http) GetUri() string {
	return h.Uri
}

func (h Http) Read() ([]byte, error) {
	data := []byte{}
	resp, err := http.Get(h.Uri)

	if err != nil {
		return data, err
	}

	defer resp.Body.Close()

	_, err = resp.Body.Read(data)

	return data, nil
}

func (h Http) Validate() error {
	return h.validateUri(h)
}

func (h Http) Sha256() (string, error) {
	return h.sha256(h)
}

func (h Http) GetFileName() string {
	return h.getFileName(h)
}
