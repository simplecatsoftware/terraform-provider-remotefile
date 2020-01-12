package remotefile

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"terraform-provider-remotefiles/remotefile/use_case"
)

func dataSourceRemoteFilesRead() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceRemoteFilesReadRead,
		Schema: map[string]*schema.Schema{
			"source": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func dataSourceRemoteFilesReadRead(data *schema.ResourceData, meta interface{}) error {
	sourceUri := data.Get("source").(string)

	sourceFile, err := use_case.Factory(sourceUri)
	if HandleError(err) {
		return err
	}

	fileName := sourceFile.GetFileName()

	localFile, err := use_case.Factory(fmt.Sprintf("tmp://*%s", fileName))
	if HandleError(err) {
		return err
	}

	err = use_case.Copy(sourceFile, localFile)
	if HandleError(err) {
		return err
	}

	data.SetId(sourceUri)

	return nil
}
