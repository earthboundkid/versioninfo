package versioninfo

import (
	"flag"
	"fmt"
	"os"
	"time"
)

// AddFlag adds -v and -version flags to the FlagSet.
// If triggered, the flags print version information and call os.Exit(0).
// If FlagSet is nil, it adds the flags to flag.CommandLine.
func AddFlag(f *flag.FlagSet) {
	if f == nil {
		f = flag.CommandLine
	}
	f.Var(flagVersion{}, "v", "short alias for -version")
	f.Var(flagVersion{}, "version", "print version information and exit")
}

type flagVersion struct{}

func (v flagVersion) IsBoolFlag() bool {
	return true
}

func (v flagVersion) String() string {
	return ""
}

func (v flagVersion) Set(s string) error {
	if s == "true" {
		fmt.Println("Version:", Version)
		fmt.Println("Revision:", Revision)
		if Revision != "unknown" {
			fmt.Println("Committed:", LastCommit.Format(time.RFC1123))
			if DirtyBuild {
				fmt.Println("Dirty Build")
			}
		}
		os.Exit(0)
	}
	return nil
}
