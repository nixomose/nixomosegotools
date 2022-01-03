package tools

import (
	"os"
	"path/filepath"
	"syscall"
)

func File_exists(logger *Nixomosetools_logger, path string) (Ret, bool) {
	_, err := os.Stat(path)
	if err == nil {
		return nil, true
	}
	if os.IsNotExist(err) {
		return nil, false
	}
	return Error(logger, "error checking for existence of file: ", path, " err: ", err), false
}

func Is_mounted(logger *Nixomosetools_logger, mountpoint string) (Ret, bool) {
	mntpoint, err := os.Stat(mountpoint)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, false
		}
		return Error(logger, "error from stat on ", mountpoint, " err: ", err), false
	}
	parent, err := os.Stat(filepath.Join(mountpoint, ".."))
	if err != nil {
		return Error(logger, "error from stat on ", mountpoint, "/.. err: ", err), false
	}
	mntpointSt := mntpoint.Sys().(*syscall.Stat_t)
	parentSt := parent.Sys().(*syscall.Stat_t)
	return nil, mntpointSt.Dev != parentSt.Dev
}
