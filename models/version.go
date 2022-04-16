package models

type Version struct {
	Version        string `json:"version"`
	CommitHash     string `json:"commit_hash"`
	BuildTimestamp string `json:"build_timestamp"`
}
