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

package cfg

import (
	"os"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	text := `# Test Config
[device]
name = "Drive A"
serial_number = "4-8-15-16-23-42"

[server]
address = "127.0.0.1"
port = 8010

[makemkv]
output_directory = "."
makemkv_exe = "makemkvcon"

[db]
connection_string = "dbname=test-db"
`

	tmpFile, err := os.CreateTemp("", "test_load_config.*.toml")
	if err != nil {
		t.Error("Failed to create test file")
		return
	}
	defer os.Remove(tmpFile.Name())

	if _, err := tmpFile.WriteString(text); err != nil {
		t.Error("Failed to write test file")
		return
	}

	if err = LoadConfig(tmpFile.Name()); err != nil {
		t.Error("LoadConfig returned an error: ", err)
		return
	}

	if Device.Name != "Drive A" {
		t.Errorf("Device.Name = '%s', expected 'Drive A'", Device.Name)
	}

	if Device.Serial != "4-8-15-16-23-42" {
		t.Errorf("Device.Serial = '%s', expected '4-8-15-16-23-42'", Device.Serial)
	}

	if Server.Address != "127.0.0.1" {
		t.Errorf("Server.Address = '%s', expected '127.0.0.1'", Server.Address)
	}

	if Server.Port != 8010 {
		t.Errorf("Server.Port = '%d', expected 8010", Server.Port)
	}

	if MakeMkv.OutDir != "." {
		t.Errorf("MakeMkv.OutDir = '%s', expected '.'", MakeMkv.OutDir)
	}

	if MakeMkv.MakeMKV != "makemkvcon" {
		t.Errorf("MakeMkv.MakeMKV = '%s', expected 'makemkvcon'", MakeMkv.MakeMKV)
	}

	if Db.ConnStr != "dbname=test-db" {
		t.Errorf("Db.ConnStr = '%s', expected 'dbname=test-db'", Db.ConnStr)
	}
}

func TestDeviceConfigValidation(t *testing.T) {
	valid := DeviceConfig{
		Name:   "Valid Drive",
		Serial: "123-456-789",
	}

	if err := valid.Validate(); err != nil {
		t.Error("Expected valid device config.")
	}

	invalid := []DeviceConfig{
		{Name: "", Serial: "123-456-789"},
		{Name: "Valid Drive", Serial: ""},
		{Name: "", Serial: ""},
	}

	for _, cfg := range invalid {
		if err := cfg.Validate(); err == nil {
			t.Errorf("Expected invalid device config: %+v", cfg)
		}
	}
}

func TestServerConfigValidation(t *testing.T) {
	valid := ServerConfig{
		Address: "127.0.0.1",
		Port:    8080,
	}

	if err := valid.Validate(); err != nil {
		t.Error("Expected valid server config.")
	}

	invalid := []ServerConfig{
		{Address: "", Port: 8080},
		{Address: "127.0.0.1", Port: -1},
		{Address: "127.0.0.1"},
		{Address: ""},
	}

	for _, cfg := range invalid {
		if err := cfg.Validate(); err == nil {
			t.Errorf("Expected invalid server config: %+v", cfg)
		}
	}
}

func TestMakeMkvConfigValidation(t *testing.T) {
	valid := MakeMkvConfig{
		OutDir:  ".",
		MakeMKV: "makemkvcon",
	}

	if err := valid.Validate(); err != nil {
		t.Error("Expected valid MakeMKV config.")
	}

	invalid := []MakeMkvConfig{
		{OutDir: "", MakeMKV: "makemkvcon"},
		{OutDir: ".", MakeMKV: ""},
		{OutDir: "", MakeMKV: ""},
		{OutDir: "unlikely/to/exist", MakeMKV: ""},
	}

	for _, cfg := range invalid {
		if err := cfg.Validate(); err == nil {
			t.Errorf("Expected invalid MakeMKV config: %+v", cfg)
		}
	}
}

func TestDbConfigValidation(t *testing.T) {
	valid := DatabaseConfig{
		ConnStr: "dbname=test-db",
	}

	if err := valid.Validate(); err != nil {
		t.Error("Expected valid database config.")
	}

	invalid := []DatabaseConfig{
		{ConnStr: ""},
	}

	for _, cfg := range invalid {
		if err := cfg.Validate(); err == nil {
			t.Errorf("Expected invalid database config: %+v", cfg)
		}
	}
}
