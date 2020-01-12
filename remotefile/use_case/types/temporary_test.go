package types

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

type TemporaryTestSuite struct {
	suite.Suite
}

func (suite *TemporaryTestSuite) TestTemporaryGetProtocols() {
	tmp := Temporary{Uri: "tmp://tmp-file-protocols"}
	err := tmp.Validate()

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), []string{"tmp"}, tmp.GetProtocols(), "protocols do not match")
}

func (suite *TemporaryTestSuite) TestTemporaryWriteRead() {
	tmp, err := MakeTestTemporaryFile("tmp://tmp-file-protocols")
	assert.NoError(suite.T(), err)

	err = tmp.Write([]byte("NEW FILE"))
	assert.NoError(suite.T(), err)

	retrievedFile, err := tmp.Read()
	assert.NoError(suite.T(), err)

	assert.Equal(suite.T(), "NEW FILE", string(retrievedFile), "newly created file content \"%s\" does not match expected \"%s\"", retrievedFile, "NEW FILE")
}

func (suite *TemporaryTestSuite) TestTemporarySha256() {
	tmp, err := MakeTestTemporaryFile("tmp://tmp-file-sha256")
	assert.NoError(suite.T(), err)

	err = tmp.Write([]byte("KNOWN HASH"))
	assert.NoError(suite.T(), err)

	knownHash := "50bc30d15bd9d6b6c1160fb03d5216fc5b047d23eccd1a4add8f51accc44e3e5"
	computedHash, err := tmp.Sha256()
	assert.NoError(suite.T(), err)

	assert.Equal(suite.T(), knownHash, computedHash)
}

func MakeTestTemporaryFile(uri string) (Temporary, error) {
	info, err := ioutil.TempFile(os.TempDir(), strings.Replace(uri, "tmp://", "", -1))
	if err != nil {
		return Temporary{}, err
	}

	file := File{Uri: fmt.Sprintf("file://%s", info.Name())}
	err = file.Validate()
	if err != nil {
		return Temporary{}, err
	}

	tmp := Temporary{Uri: uri, File: file}

	err = tmp.Validate()

	return tmp, err
}

func TestTemporaryType(t *testing.T) {
	suite.Run(t, new(TemporaryTestSuite))
}
