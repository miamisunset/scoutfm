package fs

import (
	"io/fs"
	"io/ioutil"
	"log"
)

func ReadDir(cwd string) []fs.FileInfo {
	files, err := ioutil.ReadDir(cwd)
	if err != nil {
		log.Fatal(err)
	}
	return files
}
