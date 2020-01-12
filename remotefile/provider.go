package remotefile

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		DataSourcesMap: map[string]*schema.Resource{
			"remotefile_read": dataSourceRemoteFilesRead(),
		},
	}
}

func HandleError(err error) bool {
	return err != nil
}