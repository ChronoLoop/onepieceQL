package utils

import (
	"os"
	"path"
	"path/filepath"
)

func GetExePath() string {
	ex, err := os.Executable()

	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)

	//air runs executable in tmp
	if path.Base(exPath) == "tmp" {
		return filepath.Join(exPath, "../")
	}

	wd, err := os.Getwd()

	if err != nil {
		panic(err)
	}
	return wd
}
