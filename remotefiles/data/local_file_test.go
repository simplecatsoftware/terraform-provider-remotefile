package data

import (
	"terraform-provider-remotefiles/test"
	"testing"
)

func TestLocalFileHasAUri(t *testing.T) {
	file := test.TempFile("test")
	localFile := LocalFile{Path: file.Name()}

	if file.Name() != localFile.Path {
		t.Fatal("remoteFile.Uri does not match uri", localFile.Path, file.Name())
	}
}
