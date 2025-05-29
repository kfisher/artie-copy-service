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

// Package cfg handles managing configuration information. Configuration is
// expected to be set during startup and not change afterwards. Therefore, they
// should be safe to read in any routine without the need for a synchronization
// method such as mutexes.
package cfg

import (
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/pelletier/go-toml/v2"
)

var (
	Device  DeviceConfig
	Server  ServerConfig
	MakeMkv MakeMkvConfig
	Db      DatabaseConfig
)

// LoadConfig loads the configuration options provided by the TOML file `path`
// and updates the service's global config.
func LoadConfig(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("failed to open config file: %w", err)
	}
	defer file.Close()

	bs, err := io.ReadAll(file)
	if err != nil {
		return fmt.Errorf("failed to read config file: %w", err)
	}

	var config serviceConfig
	if err := toml.Unmarshal(bs, &config); err != nil {
		return fmt.Errorf("failed to parse config file: %w", err)
	}

	if err = config.Device.Validate(); err != nil {
		return fmt.Errorf("invalid device configuration: %w", err)
	}

	if err = config.Server.Validate(); err != nil {
		return fmt.Errorf("invalid server configuration: %w", err)
	}

	if err = config.MakeMKV.Validate(); err != nil {
		return fmt.Errorf("invalid makemkv configuration: %w", err)
	}

	if err = config.Db.Validate(); err != nil {
		return fmt.Errorf("invalid db configuration: %w", err)
	}

	Device = config.Device
	Server = config.Server
	MakeMkv = config.MakeMKV
	Db = config.Db

	return nil
}

type DeviceConfig struct {
	Name   string `toml:"name"`
	Serial string `toml:"serial_number"`
}

func (d *DeviceConfig) Validate() error {
	if d.Name == "" {
		return errors.New("name is missing or empty")
	}

	if d.Serial == "" {
		return errors.New("serial_number is missing or empty")
	}

	return nil
}

type ServerConfig struct {
	Address string `toml:"address"`
	Port    int32  `toml:"port"`
}

func (s *ServerConfig) Validate() error {
	if s.Address == "" {
		return errors.New("address is missing or empty")
	}

	if s.Port <= 0 {
		return errors.New("port missing or invalid")
	}

	return nil
}

type MakeMkvConfig struct {
	OutDir  string `toml:"output_directory"`
	MakeMKV string `toml:"makemkv_exe"`
}

func (m *MakeMkvConfig) Validate() error {
	if m.OutDir == "" {
		return errors.New("output_directory is missing or empty")
	}

	if _, err := os.Stat(m.OutDir); os.IsNotExist(err) {
		return errors.New("output_directory does not exist or is not a directory")
	}

	// TODO[MED] Determine how to verify that the executable exists and is
	// executable.

	if m.MakeMKV == "" {
		return errors.New("makemkv_exe is missing or empty")
	}

	return nil
}

type DatabaseConfig struct {
	ConnStr string `toml:"connection_string"`
}

func (d *DatabaseConfig) Validate() error {
	if d.ConnStr == "" {
		return errors.New("db is missing or empty")
	}

	return nil
}

type serviceConfig struct {
	Device  DeviceConfig
	Server  ServerConfig
	MakeMKV MakeMkvConfig
	Db      DatabaseConfig
}
