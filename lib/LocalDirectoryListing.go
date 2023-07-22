package lib

import (
	"gocrafter/models/data"
	"os"
)

func GetChildItems(path string) ([]data.ChildItem, error) {
	var children []data.ChildItem

	files, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		child := data.ChildItem{
			Path:  file.Name(),
			IsDir: file.IsDir(),
		}
		children = append(children, child)
	}

	return children, nil
}
