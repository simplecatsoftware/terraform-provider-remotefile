package fetch

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type RemoteFileTypeTestSuite struct {
	suite.Suite
}

func (suite *RemoteFileTypeTestSuite) TestRemoteFileHasAUri() {
	uri := "http://example.com/index.html"
	remoteFile := RemoteFile{Uri: uri}

	if uri != remoteFile.Uri {
		suite.T().Fatal("remoteFile.Uri does not match uri", remoteFile.Uri, uri)
	}
}

func TestRemoteFileTypeTestSuite(t *testing.T) {
	suite.Run(t, new(RemoteFileTypeTestSuite))
}
