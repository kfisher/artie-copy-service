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

package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	for _, arg := range os.Args[1:] {
		if arg == "info" {
			SimulateInfoCommand()
			return
		}
		if arg == "mkv" {
			SimulateMkvCommand()
			return
		}
	}

	fmt.Fprintln(os.Stderr, "Error: No valid command provided. Use 'info' or 'mkv'.")
	os.Exit(1)
}

func SimulateInfoCommand() {
	path := os.Getenv("FAUX_MAKEMKV_INFO_PATH")
	if path == "" {
		fmt.Fprintln(os.Stderr, "Error: FAUX_MAKEMKV_INFO_PATH environment variable must be set")
		os.Exit(1)
	}

	EchoFile(path)
}

func SimulateMkvCommand() {
	path := os.Getenv("FAUX_MAKEMKV_MKV_PATH")

	if path == "" {
		fmt.Fprintln(os.Stderr, "Error: FAUX_MAKEMKV_MKV_PATH environment variable must be set")
		os.Exit(1)
	}

	EchoFile(path)
}

func GetDelay() time.Duration {
	s := os.Getenv("FAUX_MAKEMKV_DELAY")
	if s != "" {
		if a, b := strconv.ParseInt(s, 10, 32); b == nil {
			return time.Duration(a) * time.Millisecond
		}
	}

	return 0
}

func EchoFile(path string) {
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v", err)
		os.Exit(1)
	}

	delay := GetDelay()

	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		if strings.TrimSpace(line) == "" {
			continue // Skip empty lines
		} else {
			fmt.Println(line)
		}

		time.Sleep(delay)
	}
}
