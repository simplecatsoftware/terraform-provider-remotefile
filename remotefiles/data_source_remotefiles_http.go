package remotefiles

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"terraform-provider-remotefiles/remotefiles/fetch"
)

func dataSourceHttp() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceHttpRead,
		Schema: map[string]*schema.Schema{
			"uri": {
				Type:     schema.TypeString,
				Required: true,
			},
			"path": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func dataSourceHttpRead(data *schema.ResourceData, meta interface{}) error {
	uri := data.Get("uri").(string)
	path := data.Get("path").(string)

	if path == "" {
		file := fetch.TempFile("")
		path = file.Name()

		err := data.Set("path", path)

		if err != nil {
			panic(err)
		}
	}

	data.SetId(uri)

	return nil
}
