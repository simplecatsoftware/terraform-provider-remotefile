package use_case

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type TemporaryFileTestSuite struct {
	suite.Suite
}

func (suite *TemporaryFileTestSuite) TestTemporaryFile() {
	tmp, err := TemporaryFile("temporary-file", []byte("TEST FILE"))
	assert.NoError(suite.T(), err)

	content, err := tmp.Read()
	assert.NoError(suite.T(), err)

	assert.Equal(suite.T(), "TEST FILE", string(content), "temporary contents do not match")
}

func TestTemporaryFile(t *testing.T) {
	suite.Run(t, new(TemporaryFileTestSuite))
}
