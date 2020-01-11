package fetch

import "testing"

func TestRemoteFileHasAUri(t *testing.T) {
	uri := "http://example.com/index.html"
	remoteFile := RemoteFile{Uri: uri}

	if uri != remoteFile.Uri {
		t.Fatal("remoteFile.Uri does not match uri", remoteFile.Uri, uri)
	}
}
