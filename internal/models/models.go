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

// Package models defines the data models. This includes both the models used
// in the database and as data transfer objects between services.
package models

// TODO: Most of these models will be moved to a common or core project so that
//       they can be used across multiple projects. So there may be some data
//       fields that aren't required for the copy service.

// Specifies the different states of an optical drive in the context of the
// application, not the OS level. For example, the application doesn't care if
// the drive is open, opening, or closing; It just cares if there is a disc
// inserted or not and if a copy operation is in progress.
type OpticalDriveState string

const (
	// TODO: DOC
	DriveStateIdle OpticalDriveState = "idle"

	// TODO: DOC
	DriveStateCopying OpticalDriveState = "copying"
)

// OpticalDrive represents an optical drive.
type OpticalDrive struct {
	// Id is the unique identifier associated with the drive. This is mainly
	// just used as the database identifier. Outside of the database, drives
	// will usually be identified by their serial number.
	Id int

	// Name is the name assigned to the drive via configuration. It is used
	// to make it easier to distinguish between multiple drives. Its easier
	// to know what drive is being referred to in a UI if its called Top
	// Drive vs some cryptic serial number.
	Name string

	// Host is the hostname or ip address of the computer the drive is
	// connected to or installed in.
	Host string

	// DeviceName the name or path of the device set by the system. On
	// Linux, this will be something like /dev/sr0.
	DeviceName string

	// SerialNumber is the serial number assigned by the manufacturer as
	// reported by the system. Although not a guarantee, it is assumed that
	// these will be unique per drive.
	SerialNumber string

	// State is the current state of the optical drive service instance.
	State OpticalDriveState

	// DiscLabel is the label of the disc reported by the system. If a disc,
	// is not inserted into the drive, it will be an empty string.
	DiscLabel string
}
