package remotefiles

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"terraform-provider-remotefiles/remotefiles/fetch"
	"testing"
)

func TestDataSourceHttpWithGivenPath(t *testing.T) {
	uri := "https://github.com/simplecatsoftware/terraform-provider-rfile/archive/master.zip"
	path := fetch.TempFile("data.remotefiles_http.test")

	err := path.Close()
	if err != nil {
		t.Fatal(err)
	}

	resource.Test(t, resource.TestCase{
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

func TestDataSourceHttpWithoutGivenPath(t *testing.T) {
	uri := "https://github.com/simplecatsoftware/terraform-provider-rfile/archive/master.zip"

	resource.Test(t, resource.TestCase{
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
