package fetch

import (
	"testing"
)

func TestLocalFileHasAUri(t *testing.T) {
	file := TempFile("test")
	localFile := LocalFile{Path: file.Name()}

	if file.Name() != localFile.Path {
		t.Fatal("remoteFile.Uri does not match uri", localFile.Path, file.Name())
	}
}
