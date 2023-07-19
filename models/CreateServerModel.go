package models

import (
	"encoding/json"
	"net/http"
	"time"
)

type VersionManifest struct {
	Latest   LatestVersion  `json:"latest"`
	Versions []VersionEntry `json:"versions"`
}

type LatestVersion struct {
	Release  string `json:"release"`
	Snapshot string `json:"snapshot"`
}

type VersionEntry struct {
	ID              string    `json:"id"`
	Type            string    `json:"type"`
	URL             string    `json:"url"`
	Time            time.Time `json:"time"`
	ReleaseTime     time.Time `json:"releaseTime"`
	SHA1            string    `json:"sha1"`
	ComplianceLevel int       `json:"complianceLevel"`
}

var CachedVersionManifest = GetManifest()

func GetManifest() VersionManifest {
	resp, _ := http.Get("https://piston-meta.mojang.com/mc/game/version_manifest_v2.json")
	defer resp.Body.Close()

	var manifest VersionManifest
	json.NewDecoder(resp.Body).Decode(&manifest)

	return manifest
}
