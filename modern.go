//go:build go1.18

package versioninfo

import (
	"runtime/debug"
	"time"
)

func init() {
	info, ok := debug.ReadBuildInfo()
	if !ok {
		return
	}
	for _, kv := range info.Settings {
		switch kv.Key {
		case "gitrevision":
			Revision = kv.Value
		case "gitcommittime":
			LastCommit, _ = time.Parse(time.RFC3339, kv.Value)
		case "gituncommitted":
			DirtyBuild = kv.Value == "true"
		}
	}
}
