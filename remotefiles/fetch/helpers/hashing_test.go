package helpers

import (
	"github.com/stretchr/testify/suite"
	"terraform-provider-remotefiles/remotefiles/fetch"
	"testing"
)

type HashSha256TestSuite struct {
	suite.Suite
}

func (suite *HashSha256TestSuite) TestHashSha256() {
	knownFile := fetch.LocalFile{Path: "hash_known.txt"}
	knownHash := "df2e8c587f36dee996ace336501547402d7ab6529dacb1e9e8003d5a1974efc8"
	computedHash, err := HashSha256(knownFile)

	if err != nil {
		suite.T().Fatal("Could not generate computed hash", err)
	}

	if knownHash != computedHash {
		suite.T().Fatal("Computed hash does not match known hash", knownHash, "!=", computedHash)
	}
}

func TestHashSha256TestSuite(t *testing.T) {
	suite.Run(t, new(HashSha256TestSuite))
}