package use_case

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"os"
	"terraform-provider-remotefiles/remotefile/use_case/types"
	"testing"
)

type FactoryTestSuite struct {
	suite.Suite
	KnownFilePath string
}

func (suite *FactoryTestSuite) SetupTest() {
	cwd, _ := os.Getwd()
	suite.KnownFilePath = fmt.Sprintf("file://%s/%s", cwd, "known_file.txt")
}

func (suite *FactoryTestSuite) TestFileFactory() {
	file, err := FileFactory(suite.KnownFilePath)

	assert.NoError(suite.T(), err)
	assert.IsType(suite.T(), types.File{}, file)
}

func (suite *FactoryTestSuite) TestFileAutoFactory() {
	file, err := Factory(suite.KnownFilePath)

	assert.NoError(suite.T(), err)
	assert.IsType(suite.T(), types.File{}, file)
}

func (suite *FactoryTestSuite) TestTemporaryFactory() {
	file, err := TemporaryFactory("tmp://temporary-factory")

	assert.NoError(suite.T(), err)
	assert.IsType(suite.T(), types.Temporary{}, file)
}

func (suite *FactoryTestSuite) TestTemporaryAutoFactory() {
	tmpUri := fmt.Sprintf("tmp://%s", "")
	file, _ := Factory(tmpUri)

	assert.IsType(suite.T(), types.Temporary{}, file)
}

func TestFactoryTestSuite(t *testing.T) {
	suite.Run(t, new(FactoryTestSuite))
}
