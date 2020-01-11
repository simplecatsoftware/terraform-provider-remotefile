package remotefiles

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/stretchr/testify/suite"
	"terraform-provider-remotefiles/remotefiles/fetch"
	"testing"
)

type DataSourceRemotefilesHttpFileTestSuite struct {
	suite.Suite
}

func (suite *DataSourceRemotefilesHttpFileTestSuite) TestDataSourceHttpWithGivenPath() {
	uri := "https://github.com/simplecatsoftware/terraform-provider-rfile/archive/master.zip"
	path := fetch.TempFile("data.remotefiles_http.test")

	err := path.Close()
	if err != nil {
		suite.T().Fatal(err)
	}

	resource.Test(suite.T(), resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
data "remotefiles_http" "test" {
  uri = "%s"
  path = "%s"
}`, uri, path.Name()),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.remotefiles_http.test", "uri", uri),
					resource.TestCheckResourceAttr("data.remotefiles_http.test", "path", path.Name()),
				),
			},
		},
	})
}

func (suite *DataSourceRemotefilesHttpFileTestSuite) TestDataSourceHttpWithoutGivenPath() {
	uri := "https://github.com/simplecatsoftware/terraform-provider-rfile/archive/master.zip"

	resource.Test(suite.T(), resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
data "remotefiles_http" "test" {
  uri = "%s"
}`, uri),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.remotefiles_http.test", "uri", uri),
					resource.TestCheckResourceAttrSet("data.remotefiles_http.test", "path"),
				),
			},
		},
	})
}

func TestDataSourceRemotefilesHttpFileTestSuite(t *testing.T) {
	suite.Run(t, new(DataSourceRemotefilesHttpFileTestSuite))
}
