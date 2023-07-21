package lib

import (
	"gocrafter/models"
	"os"
)

func GetChildItems(path string) ([]models.ChildItem, error) {
	var children []models.ChildItem

	files, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		child := models.ChildItem{
			Path:  file.Name(),
			IsDir: file.IsDir(),
		}
		children = append(children, child)
	}

	return children, nil
}
