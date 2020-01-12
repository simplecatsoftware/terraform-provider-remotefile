package remotefile

import (
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
			"local_path": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"actual_sha256": {
				Type:     schema.TypeString,
				Computed: true,
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

	_, err = sourceFile.Read()
	if HandleError(err) {
		return err
	}

	filePath := sourceFile.GetFilePath()

	err = data.Set("local_path", filePath)
	if HandleError(err) {
		return err
	}

	sha256, err := sourceFile.Sha256()
	if HandleError(err) {
		return err
	}

	err = data.Set("actual_sha256", sha256)
	if HandleError(err) {
		return err
	}

	data.SetId(sourceUri)

	return nil
}
