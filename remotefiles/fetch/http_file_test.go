package fetch

import (
	"io/ioutil"
	"strings"
	"testing"
)

func TestFetchHttp(t *testing.T) {
	file := TempFile("example")
	localFile := LocalFile{Path: file.Name()}
	remoteFile := RemoteFile{Uri: "http://example.com/index.html"}

	err := HttpFile(remoteFile.Uri, localFile.Path)

	if err != nil {
		t.Fatal(err)
	}

	fileContent, err := ioutil.ReadFile(file.Name())
	if err != nil {
		t.Fatal(err)
	}

	if !strings.Contains(string(fileContent), "Example Domain") {
		t.Fatal("File", file.Name(), "does not contain the words 'Example Domain'")
	}
}
