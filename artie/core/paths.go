/*
  Onix Config Manager - Artie
  Copyright (c) 2018-2020 by www.gatblau.org
  Licensed under the Apache License, Version 2.0 at http://www.apache.org/licenses/LICENSE-2.0
  Contributors to this project, hereby assign copyright in this code to the project,
  to be licensed under the same terms as the rest of the code.
*/
package core

import (
	"fmt"
	"log"
	"os"
	"os/user"
	"path/filepath"
)

const CliName = "artie"

// gets the user home directory
func HomeDir() string {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	return usr.HomeDir
}

func WorkDir() string {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	return wd
}

// gets the root path of the local registry
func RegistryPath() string {
	return filepath.Join(HomeDir(), fmt.Sprintf(".%s", CliName))
}

// temporary path for file operations
func TmpPath() string {
	return filepath.Join(RegistryPath(), "tmp")
}

func TmpExists() {
	tmp := TmpPath()
	// ensure tmp folder exists for temp file operations
	_, err := os.Stat(tmp)
	if os.IsNotExist(err) {
		_ = os.MkdirAll(tmp, os.ModePerm)
	}
}
