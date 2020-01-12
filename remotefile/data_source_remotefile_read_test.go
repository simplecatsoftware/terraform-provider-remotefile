package remotefile

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/stretchr/testify/suite"
	"testing"
)

type DataSourceRemoteFilesTestSuite struct {
	suite.Suite
}

func (suite *DataSourceRemoteFilesTestSuite) TestDataSourceHttp() {
	resourceName := "data.remotefile_read.http_file"
	sourceUri := "http://example.org/index.html"

	resource.Test(suite.T(), resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
data "remotefile_read" "http_file" {
  source = "%s"
}`, sourceUri),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "source", sourceUri),
					resource.TestCheckResourceAttr(resourceName, "actual_sha256", "ea8fac7c65fb589b0d53560f5251f74f9e9b243478dcb6b3ea79b5e36449c8d9"),
					resource.TestCheckResourceAttrSet(resourceName, "local_path"),
				),
			},
		},
	})
}

func TestDataSourceRemoteFileReadTestSuite(t *testing.T) {
	suite.Run(t, new(DataSourceRemoteFilesTestSuite))
}
