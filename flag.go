package versioninfo

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"time"
)

// AddFlag adds -v and -version flags to the FlagSet.
// If triggered, the flags print version information and call os.Exit(0).
// If FlagSet is nil, it adds the flags to flag.CommandLine.
func AddFlag(f *flag.FlagSet) {
	AddFlagWithFunc(f, printVersion)
}

// AddFlagWithFunc is like AddFlag but with a custom version-print function.
func AddFlagWithFunc(f *flag.FlagSet, do func(b bool) error) {
	if f == nil {
		f = flag.CommandLine
	}
	f.Var(BoolFunc(do), "v", "short alias for -version")
	f.Var(BoolFunc(do), "version", "print version information and exit")
}

// printVersion is the default version print.
func printVersion(b bool) error {
	if !b {
		return nil
	}
	fmt.Println("Version:", Version)
	fmt.Println("Revision:", Revision)
	if Revision != "unknown" {
		fmt.Println("Committed:", LastCommit.Format(time.RFC1123))
		if DirtyBuild {
			fmt.Println("Dirty Build")
		}
	}
	os.Exit(0)
	panic("unreachable")
}

// BoolFunc allows to customize the command line flags.
type BoolFunc func(bool) error

func (f BoolFunc) IsBoolFlag() bool {
	return true
}

func (f BoolFunc) String() string {
	return ""
}

func (f BoolFunc) Set(s string) error {
	b, err := strconv.ParseBool(s)
	if err != nil {
		return err
	}
	return f(b)
}
