package fs

import (
	"io/fs"
	"io/ioutil"
	"log"
	"sort"
)

func ReadDir(cwd string) []fs.FileInfo {
	files, err := ioutil.ReadDir(cwd)
	if err != nil {
		log.Fatal(err)
	}
	return files
}

func SortByName(files []fs.FileInfo, desc bool) {
	sort.Slice(files, func(i, j int) bool {
		if desc {
			return files[i].Name() < files[j].Name()
		}

		return files[i].Name() > files[j].Name()
	})
}

func SortDirFirst(files []fs.FileInfo) {
	sort.Slice(files, func(i, _ int) bool {
		return files[i].IsDir()
	})
}
