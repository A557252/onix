//   Onix Config DatabaseProvider - Dbman
//   Copyright (c) 2018-2020 by www.gatblau.org
//   Licensed under the Apache License, Version 2.0 at http://www.apache.org/licenses/LICENSE-2.0
//   Contributors to this project, hereby assign copyright in this code to the project,
//   to be licensed under the same terms as the rest of the code.
package cmd

import (
	"fmt"
	. "github.com/gatblau/onix/dbman/core"
	"github.com/spf13/cobra"
	"io/ioutil"
	"strings"
)

type ConfigListCmd struct {
	cmd *cobra.Command
}

func NewConfigListCmd() *ConfigListCmd {
	c := &ConfigListCmd{
		cmd: &cobra.Command{
			Use:     "list",
			Short:   "list all available configuration sets",
			Example: `dbman config list`,
		}}
	c.cmd.Run = c.Run
	return c
}

func (c *ConfigListCmd) Run(cmd *cobra.Command, args []string) {
	// get the files in the current path
	files, err := ioutil.ReadDir(DM.GetConfigSetDir())
	if err != nil {
		fmt.Printf("!!! I cannot read from directory %v: %v", DM.GetConfigSetDir(), err)
		return
	}
	// print a list
	for _, file := range files {
		if strings.HasPrefix(file.Name(), ".dbman_") {
			if strings.Contains(DM.Cfg.ConfigFileUsed(), file.Name()) {
				fmt.Print("* ")
			} else {
				fmt.Print("  ")
			}
			fmt.Println(file.Name()[7 : len(file.Name())-5])
		}
	}
}
