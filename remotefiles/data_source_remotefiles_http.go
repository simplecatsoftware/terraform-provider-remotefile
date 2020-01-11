package remotefiles

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
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

	data.SetId(uri)

	return nil
}
