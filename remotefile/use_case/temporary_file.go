package use_case

import (
	"fmt"
	"terraform-provider-remotefiles/remotefile/use_case/types"
)

func TemporaryFile(name string, contents []byte) (types.T, error) {
	file, err := Factory(fmt.Sprintf("tmp://%s", name))
	if err != nil {
		return nil, err
	}

	err = file.Write(contents)
	if err != nil {
		return nil, err
	}

	return file, nil
}
