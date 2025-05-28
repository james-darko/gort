//go:build testing

package gort

import (
	"os"
	"path"
)

func NearestGitRoot() string {
	dir, err := os.Getwd()
	if err != nil {
		panic("failed to get current working directory: " + err.Error())
	}
	for dir != "/" {
		if _, err := os.Stat(dir + "/.git"); err == nil {
			return dir
		}
		dir = dir[:len(dir)-len("/"+path.Base(dir))]
	}
	panic("git directory not found in current path or any parent directory")
}
