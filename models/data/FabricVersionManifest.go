package data

import (
	"encoding/json"
	"net/http"
)

type FabricGameVersion struct {
	Version string `json:"version"`
	Stable  bool   `json:"stable"`
}

type FabricGameVersionList []FabricGameVersion

var FabricGameVersions = getVersions()

func getVersions() FabricGameVersionList {
	url := "https://meta.fabricmc.net/v2/versions/game"
	resp, _ := http.Get(url)
	defer resp.Body.Close()
	var versions FabricGameVersionList
	json.NewDecoder(resp.Body).Decode(&versions)
	return versions
}
