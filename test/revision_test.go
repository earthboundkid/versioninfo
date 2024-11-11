package main_test

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/rogpeppe/go-internal/testscript"
)

func Test(t *testing.T) {
	testscript.Run(t, testscript.Params{
		Dir: "testdata",
		Setup: func(e *testscript.Env) error {
			gomod := `
module main

require github.com/earthboundkid/versioninfo/v2 v2.0.0

replace github.com/earthboundkid/versioninfo/v2 => WORKDIR/..
`
			wd, _ := os.Getwd()
			gomod = strings.Replace(gomod, "WORKDIR", wd, 1)
			os.MkdirAll(filepath.Join(e.WorkDir, "main"), 0755)
			return os.WriteFile(filepath.Join(e.WorkDir, "main", "go.mod"), []byte(gomod), 0644)
		},
	})
}
