package misc

import (
	"os"
	"testing"
)

const devRootDir = "/Users/lawrence/Projects/Learn WebDev/functional-server"

func TestGetRootDir(t *testing.T) {
	os.Chdir("../../")
	// NOTE: in testing, the supposedly cmd dir becomes runtime dir in this file
	runtimeDir := GetCmdDir()
	if runtimeDir != devRootDir {
		t.Error("incorrect root directory: expected dev root /Users/lawrence/Projects/Learn WebDev/functional-server, got", runtimeDir)
	}
}
