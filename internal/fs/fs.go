package fs

import (
	"io/fs"
	"io/ioutil"
	"log"
	"sort"
)

// ReadDir reads a directory returns all files or horribly fails. Nothing special
func ReadDir(cwd string) []fs.FileInfo {
	files, err := ioutil.ReadDir(cwd)
	if err != nil {
		log.Fatal(err)
	}
	return files
}

// SortByName takes a slice of files and a bool to indicate sorting order
// descending or ascending, and returns same slice sorted by name.
func SortByName(files []fs.FileInfo, desc bool) {
	sort.Slice(files, func(i, j int) bool {
		if desc {
			return files[i].Name() < files[j].Name()
		}

		return files[i].Name() > files[j].Name()
	})
}

// SortDirFirst returns file slice sorted, so directories appear first.
func SortDirFirst(files []fs.FileInfo) {
	sort.Slice(files, func(i, _ int) bool {
		return files[i].IsDir()
	})
}
