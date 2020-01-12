package use_case

import "terraform-provider-remotefiles/remotefile/use_case/types"

func Copy(source types.T, destination types.T) error {
	sourceContents, err := source.Read()
	if err != nil {
		return err
	}

	err = destination.Write(sourceContents)
	if err != nil {
		return err
	}

	return nil
}
