package util

import (
	"io/fs"
	"os/user"
	"strconv"
	"syscall"
)

func GetFileOwner(file fs.FileInfo) (string, string) {
	stat := file.Sys().(*syscall.Stat_t)

	// TODO: handle error
	u, _ := user.LookupId(strconv.FormatUint(uint64(stat.Uid), 10))
	g, _ := user.LookupGroupId(strconv.FormatUint(uint64(stat.Gid), 10))

	return u.Username, g.Name
}
