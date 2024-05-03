package utils

import "path/filepath"

func GetCommitStorePath(dataDir string) string {
	return filepath.Join(dataDir, "committer.db")
}

func GetStateStorePath(dataDir string, backend string) string {
	return filepath.Join(dataDir, backend)
}

func GetChangelogPath(commitStorePath string) string {
	return filepath.Join(commitStorePath, "changelog")
}
