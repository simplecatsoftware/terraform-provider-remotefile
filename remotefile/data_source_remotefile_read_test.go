package remotefile

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/stretchr/testify/suite"
	"testing"
)

type DataSourceRemoteFilesTestSuite struct {
	suite.Suite
}

func (suite *DataSourceRemoteFilesTestSuite) TestDataSourceHttpWithGivenPath() {
	resourceName := "data.remotefile_read.http_file"
	sourceUri := "https://github.com/simplecatsoftware/lambda-http-example/archive/master.zip"

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
					testAccCheckPathFileContains(suite.T(), resourceName, "Example Domain"),
				),
			},
		},
	})
}

func testAccCheckPathFileContains(t *testing.T, resourceName string, contains string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		//rs, ok := s.RootModule().Resources[resourceName]
		//
		//if !ok {
		//	t.Fatal("resource not found:", resourceName)
		//}
		//
		//path := rs.Primary.Attributes["destination"]
		//
		//info, err := os.Stat(path)
		//if os.IsNotExist(err) {
		//	t.Fatal("file at path does not exist", path)
		//}
		//if info.IsDir() {
		//	t.Fatal("file at path is a directory", path)
		//}
		//
		//fileContent, err := ioutil.ReadFile(path)
		//if err != nil {
		//	t.Fatal(err)
		//}
		//
		//if !strings.Contains(string(fileContent), contains) {
		//	t.Fatal("lib", path, "does not contain the words", contains)
		//}

		return nil
	}
}

func TestDataSourceRemotefilesHttpFileTestSuite(t *testing.T) {
	suite.Run(t, new(DataSourceRemoteFilesTestSuite))
}
