// Copyright © 2014, Roger Peppe

package modpath

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"os/exec"
	"strings"

	"github.com/rogpeppe/go-internal/modfile"
)

var defaultDir = "."

// Run returns the module path from the gomod file underneath
// the given directory or its all ancestor directories.
// "." is populated as dir if empty string provided.
func Run(dir string) (modulePath string, err error) {
	if dir == "" {
		dir = defaultDir
	}
	gomod, err := findModFile(dir)
	if err != nil {
		return
	}
	b, err := ioutil.ReadFile(gomod)
	if err != nil {
		return
	}
	m := modfile.ModulePath(b)
	if m == "" {
		err = fmt.Errorf("failed to find a module path from %s", gomod)
		return
	}
	modulePath = m
	return
}

// findModFile detects the absolute path to the go.mod of
// the main module via GOMOD env var.
func findModFile(dir string) (string, error) {
	out, err := runCmd(dir, "go", "env", "GOMOD")
	if err != nil {
		return "", err
	}
	out = strings.TrimSpace(out)
	if out == "" {
		return "", errors.New("no go.mod file found in any parent directory")
	}
	return strings.TrimSpace(out), nil
}

func runCmd(dir string, name string, args ...string) (string, error) {
	var outData, errData bytes.Buffer
	c := exec.Command(name, args...)
	c.Stdout = &outData
	c.Stderr = &errData
	c.Dir = dir
	err := c.Run()
	if err == nil {
		return outData.String(), nil
	}
	if _, ok := err.(*exec.ExitError); ok && errData.Len() > 0 {
		return "", errors.New(strings.TrimSpace(errData.String()))
	}
	return "", fmt.Errorf("failed to run %q: %v", append([]string{name}, args...), err)
}
