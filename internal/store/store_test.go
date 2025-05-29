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

package store_test

import (
	"testing"

	"github.com/kfisher/artie-copy-service/internal/models"
	"github.com/kfisher/artie-copy-service/internal/store"
)

func TestState(t *testing.T) {
	store.Set(models.OpticalDrive{
		Id:           1337,
		Name:         "Drive A",
		Host:         "localhost",
		DeviceName:   "/dev/sr0",
		SerialNumber: "4-8-15-16-23-42",
		State:        models.DriveStateIdle,
		DiscLabel:    "LOST_S1",
	})

	od := store.GetOpticalDrive()
	if od.Id != 1337 {
		t.Error("Expected ID to be 1337, got:", od.Id)
	}
	if od.Name != "Drive A" {
		t.Error("Expected Name to be 'Drive A', got:", od.Name)
	}
	if od.Host != "localhost" {
		t.Error("Expected Host to be 'localhost', got:", od.Host)
	}
	if od.DeviceName != "/dev/sr0" {
		t.Error("Expected DeviceName to be '/dev/sr0', got:", od.DeviceName)
	}
	if od.SerialNumber != "4-8-15-16-23-42" {
		t.Error("Expected SerialNumber to be '4-8-15-16-23-42', got:", od.SerialNumber)
	}
	if od.State != models.DriveStateIdle {
		t.Error("Expected State to be Idle, got:", od.State)
	}
	if od.DiscLabel != "LOST_S1" {
		t.Error("Expected DiscLabel to be 'LOST_S1', got:", od.DiscLabel)
	}

	if store.GetState() != models.DriveStateIdle {
		t.Error("Expected state to be Idle, got:", store.GetState())
	}

	store.SetState(models.DriveStateCopying)
	if store.GetOpticalDrive().State != models.DriveStateCopying {
		t.Error("Expected state to be Copying, got:", store.GetOpticalDrive().State)
	}
	if store.GetState() != models.DriveStateCopying {
		t.Error("Expected state to be Copying, got:", store.GetState())
	}
}
