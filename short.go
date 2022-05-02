package versioninfo

import "strings"

// Short provides a short string summarizing available version information.
func Short() string {
	parts := make([]string, 0, 3)
	if Version != "unknown" && Version != "(devel)" {
		parts = append(parts, Version)
	}
	if Revision != "unknown" && Revision != "" {
		commit := Revision
		if len(commit) > 7 {
			commit = commit[:7]
		}
		parts = append(parts, commit)
	}
	if len(parts) == 0 {
		return "devel"
	}
	if DirtyBuild {
		parts = append(parts, "dirty")
	}
	return strings.Join(parts, "-")
}
