package types

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"os"
	"testing"
)

type HttpTestSuite struct {
	suite.Suite
	TestFile string
}

func (suite *HttpTestSuite) SetupTest() {
	suite.TestFile = "https://github.com/simplecatsoftware/lambda-http-example/archive/master.zip"
}

func (suite *HttpTestSuite) TestHttpGetUri() {
	http, err := MakeTestHttpType(suite.TestFile)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), suite.TestFile, http.GetUri())
}

func (suite *HttpTestSuite) TestHttpGetName() {
	http, err := MakeTestHttpType(suite.TestFile)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), "master.zip", http.GetFileName())
}

func (suite *HttpTestSuite) TestHttpGetProtocols() {
	http, err := MakeTestHttpType(suite.TestFile)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), []string{"http", "https"}, http.GetProtocols())
}

func (suite *HttpTestSuite) TestHttpRead() {
	http, err := MakeTestHttpType(suite.TestFile)
	assert.NoError(suite.T(), err)

	_, err = http.Read()
	assert.NoError(suite.T(), err)
}

func (suite *HttpTestSuite) TestHttpWrite() {
	http, err := MakeTestHttpType(suite.TestFile)
	assert.NoError(suite.T(), err)

	err = http.Write([]byte("FILE"))
	assert.Error(suite.T(), err)
}

func (suite *HttpTestSuite) TestHttpSha256() {
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

func MakeTestHttpType(uri string) (Http, error) {
	remote := Http{Uri: uri}

	local, err := MakeTestTemporaryFile(fmt.Sprintf("tmp://*%s", remote.GetFileName()))
	if err != nil {
		return Http{}, err
	}

	http := Http{Uri: uri, LocalFile: local}

	err = http.Validate()
	if err != nil {
		return Http{}, err
	}

	return http, err
}

func TestHttpType(t *testing.T) {
	suite.Run(t, new(HttpTestSuite))
}
