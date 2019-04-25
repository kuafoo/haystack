package util

import (
	"github.com/pkg/errors"
	"os"
)

func IsWriteable(folder string) (err error) {
	fileInfo, err := os.Stat(folder)

	if err != nil {
		return err
	}

	if !fileInfo.IsDir() {
		return errors.New("Not a valid folder!")
	}

	perm := fileInfo.Mode().Perm()
	if 0200&perm != 0 {
		return nil
	}

	return errors.New("Not writeable!")
}

func GetFileSize(file *os.File) (size int64, err error) {
	var fi os.FileInfo
	if fi, err = file.Stat(); err == nil {
		size = fi.Size()
	}
	return
}
