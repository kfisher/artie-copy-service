// Copyright 2025 Kevin Fisher
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are met:
//
// 1. Redistributions of source code must retain the above copyright notice,
// this list of conditions and the following disclaimer.
//
// 2. Redistributions in binary form must reproduce the above copyright notice,
// this list of conditions and the following disclaimer in the documentation
// and/or other materials provided with the distribution.
//
// 3. Neither the name of the copyright holder nor the names of its contributors
// may be used to endorse or promote products derived from this software without
// specific prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
// AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
// IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE
// ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE
// LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR
// CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF
// SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS
// INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN
// CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE)
// ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE
// POSSIBILITY OF SUCH DAMAGE.

// Package cfg handles managing configuration information.
package cfg

import (
	"errors"
	"io"
	"log/slog"
	"os"

	"github.com/pelletier/go-toml/v2"
)

// Config contains the configuration data for the service.
var Config ServiceConfig

// ServiceConfig defines the configuration options for the service.
type ServiceConfig struct {
	Name      string `toml:"name"`
	Serial    string `toml:"serial_number"`
	Address   string `toml:"address"`
	Port      int32  `toml:"port"`
	Directory string `toml:"working_directory"`
	MakeMKV   string `toml:"makemkv_exe"`
}

// IsValid checks if the configuration is valid.
func (c *ServiceConfig) IsValid() bool {
	if c.Name == "" {
		return false
	}

	if c.Serial == "" {
		return false
	}

	if c.Address == "" {
		return false
	}

	if c.Port <= 0 {
		return false
	}

	if c.Directory == "" {
		return false
	}

	if _, err := os.Stat(c.Directory); os.IsNotExist(err) {
		return false
	}

	// TODO[MED] Determine how to verify that the executable exists and is
	// executable.

	if c.MakeMKV == "" {
		return false
	}

	return true
}

// LogValidationErrors logs any validation errors found in the configuration.
func (c *ServiceConfig) LogValidationErrors() {
	if c.Name == "" {
		slog.Error("name is required and cannot be empty")
	}

	if c.Serial == "" {
		slog.Error("serial_number is required and cannot be empty")
	}

	if c.Address == "" {
		slog.Error("address is required and cannot be empty")
	}

	if c.Port <= 0 {
		slog.Error("port must be a positive integer")
	}

	if c.Directory == "" {
		slog.Error("working_directory is required and cannot be empty")
	}

	if _, err := os.Stat(c.Directory); os.IsNotExist(err) {
		slog.Error("working_directory does not exist", "directory", c.Directory)
	}

	if c.MakeMKV == "" {
		slog.Error("makemkv_exe is required and cannot be empty")
	}
}

// Load config loads the configuration options provided by the TOML file `path`
// and updates the service's global config (see Config).
func LoadConfig(path string) error {
	slog.Info("Loading config", "path", path)

	file, err := os.Open(path)
	if err != nil {
		return errors.New("failed to open config file")
	}
	defer file.Close()

	bs, err := io.ReadAll(file)
	if err != nil {
		return errors.New("failed to read config file")
	}

	if err := toml.Unmarshal(bs, &Config); err != nil {
		return errors.New("failed to parse config file")
	}

	slog.Info("Config loaded", "config", Config)

	return nil
}
