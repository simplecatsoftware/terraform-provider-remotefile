package fetch

import (
	"io"
	"net/http"
	"os"
)

func HttpFile(remoteFile RemoteFile, localFile LocalFile) error {
	// Get the data
	resp, err := http.Get(remoteFile.Uri)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(localFile.Path)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}
