package remotefiles

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type ProviderTestSuite struct {
	suite.Suite
}

func (suite *ProviderTestSuite) TestProviderReturnsAProviderSchema() {
	provider := Provider()

	assert.IsType(suite.T(), &schema.Provider{}, provider, "Provider() returns an instance of *schema.Provider")
}

func TestProviderTestSuite(t *testing.T) {
	suite.Run(t, new(ProviderTestSuite))
}

var testAccProviders map[string]terraform.ResourceProvider
var testAccProvider *schema.Provider

func init() {
	testAccProvider = Provider()
	testAccProviders = map[string]terraform.ResourceProvider{
		"remotefiles": testAccProvider,
	}
}
