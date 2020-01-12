package types

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"io/ioutil"
	"os"
	"testing"
)

type FileTestSuite struct {
	suite.Suite
}

func (suite *FileTestSuite) TestFileGetUri() {
	d, err := os.Getwd()
	assert.NoError(suite.T(), err)

	inputUri := fmt.Sprintf("file://%s/known_file.txt", d)

	file, err := MakeTestFileType(inputUri)

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), inputUri, file.GetUri())
}

func (suite *FileTestSuite) TestFileGetName() {
	d, err := os.Getwd()
	assert.NoError(suite.T(), err)

	inputUri := fmt.Sprintf("file://%s/known_file.txt", d)

	file, err := MakeTestFileType(inputUri)

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), "known_file.txt", file.GetFileName())
}


func (suite *FileTestSuite) TestFileGetProtocols() {
	d, err := os.Getwd()
	assert.NoError(suite.T(), err)

	inputUri := fmt.Sprintf("file://%s/known_file.txt", d)

	file, err := MakeTestFileType(inputUri)
	assert.NoError(suite.T(), err)

	assert.Equal(suite.T(), []string{"file"}, file.GetProtocols(), "input does not match output %s %s", inputUri, file.GetUri())
}

func (suite *FileTestSuite) TestFileRead() {
	dir, err := os.Getwd()
	assert.NoError(suite.T(), err)

	inputUri := fmt.Sprintf("file://%s/known_file.txt", dir)
	file, err := MakeTestFileType(inputUri)
	assert.NoError(suite.T(), err)

	fileContent, err := file.Read()
	assert.NoError(suite.T(), err)

	assert.Equal(suite.T(), string(fileContent), "KNOWN TEST FILE")
}

func (suite *FileTestSuite) TestFileWrite() {
	info, err := ioutil.TempFile(os.TempDir(), "*test-file")
	assert.NoError(suite.T(), err)

	path := info.Name()
	uri := fmt.Sprintf("file://%s", path)

	f, err := MakeTestFileType(uri)
	assert.NoError(suite.T(), err)

	err = f.Write([]byte("NEW FILE"))
	assert.NoError(suite.T(), err)

	assert.FileExists(suite.T(), path)

	retrievedFile, err := ioutil.ReadFile(path)
	assert.NoError(suite.T(), err)

	assert.Equal(suite.T(), "NEW FILE", string(retrievedFile), "newly created file content \"%s\" does not match expected \"%s\"", retrievedFile, "NEW FILE")

	err = os.Remove(path)
	assert.NoError(suite.T(), err)
}

func (suite *FileTestSuite) TestFileSha256() {
	cwd, err := os.Getwd()
	assert.NoError(suite.T(), err)

	path := fmt.Sprintf("%s/known_file.txt", cwd)
	uri := fmt.Sprintf("file://%s", path)

	f := File{Uri: uri}

	knownHash := "00373683b0ecdccf8cff358d906429b5f45b6977be068b29fcd3dc57a168235a"
	computedHash, err := f.Sha256()

	assert.NoError(suite.T(), err)

	if knownHash != computedHash {
		suite.T().Fatal("Computed hash does not match known hash", knownHash, "!=", computedHash)
	}
}

func MakeTestFileType(uri string) (File, error) {
	file := File{Uri: uri}
	err := file.Validate()

	return file, err
}

func TestFileType(t *testing.T) {
	suite.Run(t, new(FileTestSuite))
}
