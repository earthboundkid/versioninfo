package main_test

import (
	"testing"

	"github.com/carlmjohnson/versioninfo"
)

func Test(t *testing.T) {
	if versioninfo.Revision == "unknown" {
		t.Fatal("bad versioninfo.Revision")
	}
	if versioninfo.LastCommit.IsZero() {
		t.Fatal("bad versioninfo.LastCommit")
	}
}
