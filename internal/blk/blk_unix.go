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

//go:build !windows

package blk

import (
	"encoding/json"
	"errors"
	"log/slog"
	"os/exec"
)

// TODO[LOW]: Implement testing for this package. It isn't straightforward to
// test because it relies on external commands whose output can vary based on
// the system configuration. Its not worth the time and complexity to update
// the code to make it testable.

// BlockDeviceList is a list of block devices reported by the os. This is really
// only needed to process the JSON output from the `lsblk` command.
type BlockDeviceList struct {
	Devices []BlockDevice `json:"blockdevices"`
}

// FindBySerial searches the list of block devices for a device with the
// specified serial number `sn`.
func (b *BlockDeviceList) FindBySerial(sn string) (BlockDevice, bool) {
	for _, dev := range b.Devices {
		if dev.Serial == sn {
			return dev, true
		}
	}
	return BlockDevice{}, false
}

// GetBlockDevice gets block device information for the device with serial
// number `sn`.
func GetBlockDevice(sn string) (BlockDevice, bool, error) {
	cmd := exec.Command("lsblk", "--list", "--paths", "--json", "--output", "NAME,SERIAL,LABEL,TYPE")
	out, err := cmd.Output()
	if err != nil {
		slog.Warn("lsblk exited with an error status", "command", cmd.Args)
		return BlockDevice{}, false, errors.New("lsblk command failed")
	}

	var blkList BlockDeviceList
	if err := json.Unmarshal(out, &blkList); err != nil {
		return BlockDevice{}, false, errors.New("failed to parse lsblk output")
	}

	dev, found := blkList.FindBySerial(sn)
	return dev, found, nil
}
