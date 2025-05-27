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
	"strings"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	tmpFile, err := os.CreateTemp("", "test_load_config.*.toml")
	if err != nil {
		t.Error("Failed to create test file")
		return
	}
	defer os.Remove(tmpFile.Name())

	lines := []string{
		"name = \"Drive A\"",
		"serial_number = \"4-8-15-16-23-42\"",
		"address = \"127.0.0.1\"",
		"port = 8010",
		"working_directory = \"/artie/working_dir\"",
		"makemkv_exe = \"makemkvcon\"",
	}

	if _, err := tmpFile.WriteString(strings.Join(lines, "\n") + "\n"); err != nil {
		t.Error("Failed to write test file")
		return
	}

	if err = LoadConfig(tmpFile.Name()); err != nil {
		t.Error("LoadConfig returned an error")
		return
	}

	if Config.Name != "Drive A" {
		t.Errorf("Config.Name = %s, expected \"Drive A\"", Config.Name)
	}

	if Config.Serial != "4-8-15-16-23-42" {
		t.Errorf("Config.Serial = %s, expected \"4-8-15-16-23-42\"", Config.Serial)
	}

	if Config.Address != "127.0.0.1" {
		t.Errorf("Config.Address = %s, expected \"127.0.0.1\"", Config.Address)
	}

	if Config.Port != 8010 {
		t.Errorf("Config.Port = %d, expected 8010", Config.Port)
	}

	if Config.Directory != "/artie/working_dir" {
		t.Errorf("Config.Directory = %s, expected \"/artie/working_dir\"", Config.Directory)
	}
}

func TestIsValidate(t *testing.T) {
	validConfig := ServiceConfig{
		Name:      "Valid Drive",
		Serial:    "123-456-789",
		Address:   "127.0.0.1",
		Port:      8080,
		Directory: "..",
		MakeMKV:   "makemkvcon",
	}

	if !validConfig.IsValid() {
		t.Error("IsValid returned false for a valid configuration")
		validConfig.LogValidationErrors()
	}

	invalidConfigs := []ServiceConfig{
		{Serial: "123-456-789", Address: "127.0.0.1", Port: 8080, Directory: "..", MakeMKV: "makemkvcon"},
		{Name: "Valid Drive", Address: "127.0.0.1", Port: 8080, Directory: "..", MakeMKV: "makemkvcon"},
		{Name: "Valid Drive", Serial: "123-456-789", Port: 8080, Directory: "..", MakeMKV: "makemkvcon"},
		{Name: "Valid Drive", Serial: "123-456-789", Address: "127.0.0.1", Directory: "..", MakeMKV: "makemkvcon"},
		{Name: "Valid Drive", Serial: "123-456-789", Address: "127.0.0.1", Port: 8080, MakeMKV: "makemkvcon"},
		{Name: "Valid Drive", Serial: "123-456-789", Address: "127.0.0.1", Port: 8080, Directory: "/probably/doesnt/exist", MakeMKV: "makemkvcon"},
		{Name: "Valid Drive", Serial: "123-456-789", Address: "127.0.0.1", Port: 8080, Directory: ".."},
	}

	for _, config := range invalidConfigs {
		if config.IsValid() {
			t.Error("IsValid returned true for an invalid configuration")
		}
	}
}
