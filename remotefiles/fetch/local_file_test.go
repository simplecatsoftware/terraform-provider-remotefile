package fetch

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type LocalFileTypeTestSuite struct {
	suite.Suite
}

func (suite *LocalFileTypeTestSuite) TestLocalFileHasAPath() {
	file := TempFile("test")
	localFile := LocalFile{Path: file.Name()}

	if file.Name() != localFile.Path {
		suite.T().Fatal("remoteFile.Uri does not match uri", localFile.Path, file.Name())
	}
}

func TestLocalFileTypeTestSuite(t *testing.T) {
	suite.Run(t, new(LocalFileTypeTestSuite))
}
