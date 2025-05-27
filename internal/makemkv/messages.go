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

import (
	"errors"
	"strconv"
	"strings"
)

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
	Type rune
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
	Attribute  Attribute
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

// ParseMessage parses a line of output from MakeMKV and returns the
// corresponding message struct instance.
func ParseMessage(line string) (any, error) {
	parts := strings.SplitN(line, ":", 2)
	if len(parts) != 2 {
		return nil, errors.New("failed to get message id and data")
	}

	data := strings.Split(parts[1], ",")

	// NOTE: Some data is ignored here because it seems to be data only
	//       relevant to MakeMKV. In fact, some of the data we are parsing
	//       we probably don't actually need.

	switch parts[0] {
	case "CINFO":
		id, err := strconv.ParseInt(data[0], 10, 32)
		if err != nil {
			return nil, errors.New("[CINFO] failed to parse attribute id")
		}
		value := strings.Trim(data[2], "\"")
		return DiscInfoMessage{Attribute{AttributeId(id), value}}, nil
	case "DRV":
		index, err := strconv.ParseInt(data[0], 10, 32)
		if err != nil {
			return nil, errors.New("[DRV] failed to parse index")
		}
		state, err := strconv.ParseInt(data[1], 10, 32)
		if err != nil {
			return nil, errors.New("[DRV] failed to parse state")
		}
		flags, err := strconv.ParseInt(data[3], 10, 32)
		if err != nil {
			return nil, errors.New("[DRV] failed to parse flags")
		}
		driveName := strings.Trim(data[4], "\"")
		discName := strings.Trim(data[5], "\"")
		device := strings.Trim(data[6], "\"")
		return DriveMessage{
			Index:     int32(index),
			State:     DriveState(state),
			Flags:     MediaFlag(flags),
			DriveName: driveName,
			DiscName:  discName,
			Device:    device,
		}, nil
	case "MSG":
		code, err := strconv.ParseInt(data[0], 10, 32)
		if err != nil {
			return nil, errors.New("[MSG] failed to parse code")
		}
		message := strings.Trim(data[3], "\"")
		return GeneralMessage{int32(code), message}, nil
	case "PRGT":
		id, err := strconv.ParseInt(data[0], 10, 32)
		if err != nil {
			return nil, errors.New("[PRGT] failed to parse id")
		}
		code, err := strconv.ParseInt(data[1], 10, 32)
		if err != nil {
			return nil, errors.New("[PRGT] failed to parse code")
		}
		name := strings.Trim(data[2], "\"")
		return ProgressTitleMessage{int32(code), int32(id), name, 'T'}, nil
	case "PRGC":
		id, err := strconv.ParseInt(data[0], 10, 32)
		if err != nil {
			return nil, errors.New("[PRGT] failed to parse id")
		}
		code, err := strconv.ParseInt(data[1], 10, 32)
		if err != nil {
			return nil, errors.New("[PRGT] failed to parse code")
		}
		name := strings.Trim(data[2], "\"")
		return ProgressTitleMessage{int32(code), int32(id), name, 'C'}, nil
	case "PRGV":
		current, err := strconv.ParseInt(data[0], 10, 32)
		if err != nil {
			return nil, errors.New("[PRGV] failed to parse current value")
		}
		total, err := strconv.ParseInt(data[1], 10, 32)
		if err != nil {
			return nil, errors.New("[PRGV] failed to parse total value")
		}
		max, err := strconv.ParseInt(data[2], 10, 32)
		if err != nil {
			return nil, errors.New("[PRGV] failed to parse max value")
		}
		return ProgressValueMessage{int32(current), int32(total), int32(max)}, nil
	case "SINFO":
		title, err := strconv.ParseInt(data[0], 10, 32)
		if err != nil {
			return nil, errors.New("[SINFO] failed to parse title index")
		}
		index, err := strconv.ParseInt(data[1], 10, 32)
		if err != nil {
			return nil, errors.New("[SINFO] failed to parse stream index")
		}
		id, err := strconv.ParseInt(data[2], 10, 32)
		if err != nil {
			return nil, errors.New("[SINFO] failed to parse attribute id")
		}
		value := strings.Trim(data[4], "\"")
		return StreamInfoMessage{int32(index), int32(title), Attribute{AttributeId(id), value}}, nil
	case "TCOUNT":
		count, err := strconv.ParseInt(data[0], 10, 32)
		if err != nil {
			return nil, errors.New("[TCOUNT] failed to parse count")
		}
		return TitleCountMessage{int32(count)}, nil
	case "TINFO":
		index, err := strconv.ParseInt(data[0], 10, 32)
		if err != nil {
			return nil, errors.New("[TINFO] failed to parse index")
		}
		id, err := strconv.ParseInt(data[1], 10, 32)
		if err != nil {
			return nil, errors.New("[CINFO] failed to parse attribute id")
		}
		value := strings.Trim(data[3], "\"")
		return TitleInfoMessage{int32(index), Attribute{AttributeId(id), value}}, nil
	default:
		return nil, errors.New("unrecognized message received")
	}
}
