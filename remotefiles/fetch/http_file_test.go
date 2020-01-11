package fetch

import (
	"github.com/stretchr/testify/suite"
	"io/ioutil"
	"strings"
	"testing"
)

type HttpFileTestSuite struct {
	suite.Suite
}

func (suite *HttpFileTestSuite) TestFetchingAnHttpFile() {
	file := TempFile("example")
	localFile := LocalFile{Path: file.Name()}
	remoteFile := RemoteFile{Uri: "http://example.com/index.html"}

	err := HttpFile(remoteFile, localFile)

	if err != nil {
		suite.T().Fatal(err)
	}

	fileContent, err := ioutil.ReadFile(file.Name())
	if err != nil {
		suite.T().Fatal(err)
	}

	if !strings.Contains(string(fileContent), "Example Domain") {
		suite.T().Fatal("File", file.Name(), "does not contain the words 'Example Domain'")
	}
}

func TestHttpFileTestSuite(t *testing.T) {
	suite.Run(t, new(HttpFileTestSuite))
}