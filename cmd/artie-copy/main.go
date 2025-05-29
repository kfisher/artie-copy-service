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
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/kfisher/artie-copy-service/internal/blk"
	"github.com/kfisher/artie-copy-service/internal/cfg"
	"github.com/kfisher/artie-copy-service/internal/db"
	"github.com/kfisher/artie-copy-service/internal/models"
	"github.com/kfisher/artie-copy-service/internal/service"
	"github.com/kfisher/artie-copy-service/internal/store"
)

func main() {
	// TODO: Add a command line flag to set the log level.
	slog.SetLogLoggerLevel(slog.LevelInfo)

	if len(os.Args) < 2 {
		fmt.Println("usage: artie-copy CONFIG")
		return
	}

	cfgPath := os.Args[1]
	slog.Info("Loading config", "path", cfgPath)
	if err := cfg.LoadConfig(cfgPath); err != nil {
		fmt.Printf("Failed to load configuration at %s\n", cfgPath)
		fmt.Printf("error: %s\n", err)
		os.Exit(1)
	}

	slog.Info("Initializing database pool.")
	if err := db.InitPool(); err != nil {
		fmt.Printf("Failed to initialize the database pool.\n")
		fmt.Printf("error: %s\n", err)
		os.Exit(1)
	}
	defer db.Close()

	slog.Info("Loading device information.")
	device, err := blk.GetBlockDevice(cfg.Device.Serial)
	if err != nil {
		fmt.Printf("Failed to get device information.\n")
		fmt.Printf("error: %s", err)
		os.Exit(1)
	}

	id, err := db.InitOpticalDriveInfo(context.Background(), cfg.Device.Serial)
	if err != nil {
		fmt.Printf("Failed to update drive info.\n")
		fmt.Printf("error: %s\n", err)
		os.Exit(1)
	}

	hostname, err := os.Hostname()
	if err != nil {
		fmt.Printf("Failed to get hostname.\n")
		fmt.Printf("error: %s\n", err)
		os.Exit(1)
	}

	od := models.OpticalDrive{
		Id:           id,
		Name:         cfg.Device.Name,
		Host:         hostname,
		DeviceName:   device.Name,
		SerialNumber: cfg.Device.Serial,
		State:        models.DriveStateIdle,
		DiscLabel:    device.Label,
	}

	store.Set(od)

	slog.Info("Starting service.", "serial", cfg.Device.Serial, "device", device.Name, "address", cfg.Server.Address, "port", cfg.Server.Port)

	if err = service.Run(); err != nil {
		fmt.Printf("Failed to run server\n")
		fmt.Printf("error: %s\n", err)
		os.Exit(1)
	}
}
