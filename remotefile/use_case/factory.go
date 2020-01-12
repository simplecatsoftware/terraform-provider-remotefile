package use_case

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"terraform-provider-remotefiles/remotefile/use_case/types"
)

func Factory(uri string) (types.T, error) {
	obj := types.T(nil)
	err := error(nil)

	if strings.Contains(uri, "file://") {
		obj, err = FileFactory(uri)
	}

	if strings.Contains(uri, "tmp://") {
		obj, err = TemporaryFactory(uri)
	}

	if strings.Contains(uri, "http://") {
		obj, err = HttpFactory(uri)
	}

	if strings.Contains(uri, "https://") {
		obj, err = HttpFactory(uri)
	}

	return obj, err
}

func FileFactory(uri string) (types.File, error) {
	file := types.File{Uri: uri}

	return file, file.Validate()
}

func HttpFactory(uri string) (types.Http, error) {
	http := types.Http{Uri: uri}

	return http, http.Validate()
}

func TemporaryFactory(uri string) (types.Temporary, error) {
	info, err := ioutil.TempFile(os.TempDir(), strings.Replace(uri, "tmp://", "", -1))
	if err != nil {
		return types.Temporary{}, err
	}

	file, err := FileFactory(fmt.Sprintf("file://%s", info.Name()))
	if err != nil {
		return types.Temporary{}, err
	}

	tmp := types.Temporary{Uri: uri, File: file}

	return tmp, tmp.Validate()
}
