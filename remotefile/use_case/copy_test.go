package use_case

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"terraform-provider-remotefiles/remotefile/use_case/types"
	"testing"
)

type CopyTestSuite struct {
	suite.Suite
}

func (suite *CopyTestSuite) TestCopy() {
	file1, err := types.TestFactory("test://test-address-1")
	assert.NoError(suite.T(), err)

	err = file1.Write([]byte("TEST FILE"))
	assert.NoError(suite.T(), err)

	file2, err := types.TestFactory("test://test-address-2")
	assert.NoError(suite.T(), err)

	err = Copy(file1, file2)
	assert.NoError(suite.T(), err, "error copying files")

	content1, err := file1.Read()
	assert.NoError(suite.T(), err, "error reading file 1")

	content2, err := file2.Read()
	assert.NoError(suite.T(), err, "error reading file 2")

	assert.Equal(suite.T(), content1, content2, "contents do not match", string(content1), string(content2))
}

func (suite *CopyTestSuite) TestCopyHttp() {
	httpFile, err := HttpFactory("https://github.com/simplecatsoftware/lambda-http-example/archive/master.zip")
	assert.NoError(suite.T(), err)

	testFile, err := TemporaryFactory(fmt.Sprintf("tmp://%s", httpFile.GetFileName()))
	assert.NoError(suite.T(), err)

	err = Copy(httpFile, testFile)
	assert.NoError(suite.T(), err, "error copying files")

	httpContent, err := httpFile.Read()
	assert.NoError(suite.T(), err)

	testContent, err := testFile.Read()
	assert.NoError(suite.T(), err)

	assert.Equal(suite.T(), httpContent, testContent)
}

func TestCopy(t *testing.T) {
	suite.Run(t, new(CopyTestSuite))
}
