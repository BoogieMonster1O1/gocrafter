package lib

import (
	"os"
)

func GetChildItems(path string, includeFiles bool) ([]string, error) {
	var children []string

	files, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		if file.IsDir() || includeFiles {
			children = append(children, file.Name())
		}
	}

	return children, nil
}
