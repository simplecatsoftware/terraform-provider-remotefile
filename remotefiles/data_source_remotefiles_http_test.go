package remotefiles

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"testing"
)

func TestDataSourceHttp(t *testing.T) {
	uri := "https://github.com/simplecatsoftware/terraform-provider-rfile/archive/master.zip"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceHttp(uri),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.remotefiles_http.test", "uri", uri),
				),
			},
		},
	})
}

func testAccPreCheck(t *testing.T) {

}

func testAccDataSourceHttp(uri string) string {
	return fmt.Sprintf(`
data "remotefiles_http" "test" {
  uri = "%s"
}`, uri)
}