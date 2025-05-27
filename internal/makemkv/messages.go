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

package makemkv

// Attribute is a key/value pair for an attribute of a disc, title or stream
// reported by MakeMKV when extracting disc information.
type Attribute struct {
	Id    AttributeId
	Value string
}

// DiscInfoMessage represents a 'CINFO' message from MakeMKV which contains
// information about a disc inserted into the drive.
type DiscInfoMessage struct {
	Attribute Attribute
}

// DriveMessage represents a 'DRV' message from MakeMKV which contains
// information about a disc drive.
type DriveMessage struct {
	Index     int32
	State     DriveState
	Flags     MediaFlag
	DriveName string
	DiscName  string
	Device    string
}

// GeneralMessage represents a 'MSG' message from MakeMKV which is a general
// information message.
type GeneralMessage struct {
	Code    int32
	Message string
}

// ProgressTitleMessage represents either a 'PRGT' or 'PRGC' message from
// MakeMKV. These messages contain the label for the current operation (PRGT)
// and the current sub-operation (PTRC) progress.
type ProgressTitleMessage struct {
	Code int32
	Id   int32
	Name string
}

// ProgressValueMessage represents the 'PRGV' message from MakeMKV which
// contains the progress values for the current operation and current
// sub-operation.
type ProgressValueMessage struct {
	Current int32
	Total   int32
	Max     int32
}

// StreamInfoMessage represents a 'SINFO' message from MakeMKV which contains
// information about an audio, subtitle, or video stream within a title on the
// disc.
type StreamInfoMessage struct {
	Index      int32
	TitleIndex int32
}

// TitleCountMessage represents a'TCOUNT' message from MakeMKV which contains
// the number of titles. This may not match the number of MKV files created
// because some may be omitted based on configuration options like minimum
// length.
type TitleCountMessage struct {
	Count int32
}

// TitleInfoMessage represents a 'TINFO' message from MakeMKV which contains
// information about a title on the disc.
type TitleInfoMessage struct {
	Index     int32
	Attribute Attribute
}

// UnknownMessage doesn't correspond to anything outputted by MakeMKV. It is
// used as a fallback if something that can be identified is encountered.
type UnknownMessage struct {
	RawMessage string
}
