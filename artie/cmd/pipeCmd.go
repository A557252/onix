/*
  Onix Config Manager - Artie
  Copyright (c) 2018-2020 by www.gatblau.org
  Licensed under the Apache License, Version 2.0 at http://www.apache.org/licenses/LICENSE-2.0
  Contributors to this project, hereby assign copyright in this code to the project,
  to be licensed under the same terms as the rest of the code.
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// list local artefacts
type PipeCmd struct {
	cmd *cobra.Command
}

func NewPipeCmd() *PipeCmd {
	c := &PipeCmd{
		cmd: &cobra.Command{
			Use:   "pipe",
			Short: "provides functions to manage artie tekton-based pipelines in K8S",
			Long:  `artie pipelines are built on tekton`,
		},
	}
	return c
}
