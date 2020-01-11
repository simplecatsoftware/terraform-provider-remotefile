package fetch

import (
	"io/ioutil"
	"log"
	"os"
)

func TempFile(name string) *os.File {
	file, err := ioutil.TempFile("", name)
	if err != nil {
		log.Fatal(err)
	}

	return file
}
