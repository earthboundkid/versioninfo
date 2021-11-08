// Package versioninfo uses runtime.ReadBuildInfo() to set global executable revision information if possible.
package versioninfo

import "time"

var (
	Version    = "unknown"
	Revision   = "unknown"
	LastCommit time.Time
	DirtyBuild = true
)
