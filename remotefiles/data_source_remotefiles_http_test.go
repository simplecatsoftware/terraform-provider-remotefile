package remotefiles

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/stretchr/testify/suite"
	"io/ioutil"
	"os"
	"strings"
	"terraform-provider-remotefiles/remotefiles/fetch"
	"testing"
)

type DataSourceRemotefilesHttpFileTestSuite struct {
	suite.Suite
}

func (suite *DataSourceRemotefilesHttpFileTestSuite) TestDataSourceHttpWithGivenPath() {
	resourceName := "data.remotefiles_http.test"
	uri := "http://example.com/index.html"
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
					resource.TestCheckResourceAttr(resourceName, "uri", uri),
					resource.TestCheckResourceAttr(resourceName, "path", path.Name()),
					testAccCheckPathFileContains(suite.T(), resourceName, "Example Domain"),
				),
			},
		},
	})
}

func (suite *DataSourceRemotefilesHttpFileTestSuite) TestDataSourceHttpWithoutGivenPath() {
	resourceName := "data.remotefiles_http.test"
	uri := "http://example.com/index.html"

	resource.Test(suite.T(), resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
data "remotefiles_http" "test" {
  uri = "%s"
}`, uri),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "uri", uri),
					resource.TestCheckResourceAttrSet(resourceName, "path"),
					testAccCheckPathFileContains(suite.T(), resourceName, "Example Domain"),
				),
			},
		},
	})
}

func testAccCheckPathFileContains(t *testing.T, resourceName string, contains string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]

		if !ok {
			t.Fatal("resource not found:", resourceName)
		}

		path := rs.Primary.Attributes["path"]

		info, err := os.Stat(path)
		if os.IsNotExist(err) {
			t.Fatal("file at path does not exist", path)
		}
		if info.IsDir() {
			t.Fatal("file at path is a directory", path)
		}

		fileContent, err := ioutil.ReadFile(path)
		if err != nil {
			t.Fatal(err)
		}

		if !strings.Contains(string(fileContent), contains) {
			t.Fatal("file", path, "does not contain the words", contains)
		}

		return nil
	}
}

func TestDataSourceRemotefilesHttpFileTestSuite(t *testing.T) {
	suite.Run(t, new(DataSourceRemotefilesHttpFileTestSuite))
}
