package app

import (
	"io/ioutil"
	"path"
	"path/filepath"

	"github.com/otiai10/copy"
)

func SaveData(targetBasePath, sourcePath string) (string, error) {
	targetDirectory, err := ioutil.TempDir(targetBasePath, "build-files-*")
	if err != nil {
		return "", err
	}
	name := filepath.Base(sourcePath)
	copy.Copy(sourcePath, path.Join(targetDirectory, name))
	return path.Join(path.Base(targetDirectory), name), nil
}
