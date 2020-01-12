package types

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type TypeTestSuite struct {
	suite.Suite
}

func (suite *TypeTestSuite) TestTypeValidateUri() {
	testType := TestType{
		Uri:     "test://test-uri",
		Content: []byte{},
	}

	assert.NoError(suite.T(), testType.Validate(), "unable to validate url schema")
}

func (suite *TypeTestSuite) TestTypeSha256() {
	testType := TestType{
		Uri:     "test://test-uri",
		Content: []byte("KNOWN TEST FILE"),
	}

	hash, err := testType.Sha256()

	assert.NoError(suite.T(), err, "unable to generate sha256")
	assert.Equal(suite.T(), "00373683b0ecdccf8cff358d906429b5f45b6977be068b29fcd3dc57a168235a", hash, "unable to validate url schema")
}

func TestTypeTestSuite(t *testing.T) {
	suite.Run(t, new(TypeTestSuite))
}
