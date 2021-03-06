/*
   Onix Config Manager - OxTerra - Terraform Http Backend for Onix
   Copyright (c) 2018-2020 by www.gatblau.org

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0
   Unless required by applicable law or agreed to in writing, software distributed under
   the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
   either express or implied.
   See the License for the specific language governing permissions and limitations under the License.

   Contributors to this project, hereby assign copyright in this code to the project,
   to be licensed under the same terms as the rest of the code.
*/
package main

import (
	. "github.com/gatblau/oxc"
	"github.com/rs/zerolog"
	"strings"
)

type Backend struct {
	config *Config
	client *Client
	ready  bool
}

// start the backend process
func (backend *Backend) start() error {
	var err error

	// load the configuration file
	if cfg, err := NewConfig(); err == nil {
		backend.config = cfg
	} else {
		return err
	}

	// initialises the logger
	backend.setLogger(backend.config.LogLevel)

	client, err := NewClient(backend.config.Ox)
	if err != nil {
		return err
	}

	// initialises the Onix REST client
	backend.client = client

	// checks if a meta model for Terraform is defined in Onix
	model := NewTerraModel(backend.client)
	err = model.create()
	if err != nil {
		return err
	}

	// the backend is now ready to receive http connections
	backend.ready = true

	// start the service listener
	svc := NewTerraService(*backend)
	svc.Start()
	return nil
}

func (backend *Backend) setLogger(logLevel string) {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	switch strings.ToLower(logLevel) {
	case "info":
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	case "debug":
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	case "error":
		zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	case "fatal":
		zerolog.SetGlobalLevel(zerolog.FatalLevel)
	case "trace":
		zerolog.SetGlobalLevel(zerolog.TraceLevel)
	}
}
