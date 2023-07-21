package models

type ChildItem struct {
	Path  string `json:"path"`
	IsDir bool   `json:"isDir"`
}
