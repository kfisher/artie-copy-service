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
	"testing"
)

func TestParseDiscInfoMessage(t *testing.T) {
	msg, err := ParseMessage("CINFO:2,0,\"The A-Team\"")
	if err != nil {
		t.Error("ParseMessage returned an error")
	}

	cinfo, ok := msg.(DiscInfoMessage)
	if !ok {
		t.Error("ParseMessage returned an incorrect type")
	}

	if cinfo.Attribute.Id != AI_NAME {
		t.Errorf("Attribute.Id = %s, expected %s", cinfo.Attribute.Id, AI_NAME)
	}

	if cinfo.Attribute.Value != "The A-Team" {
		t.Errorf("Attribute.Value = \"%s\", expected \"The A-Team\"", cinfo.Attribute.Value)
	}
}

func TestParseDriveMessage(t *testing.T) {
	msg, err := ParseMessage("DRV:2,1,999,12,\"4815162342\",\"A_TEAM\",\"/dev/sr1\"")
	if err != nil {
		t.Error("ParseMessage returned an error")
	}

	drv, ok := msg.(DriveMessage)
	if !ok {
		t.Error("ParseMessage returned an incorrect type")
	}

	if drv.Index != 2 {
		t.Errorf("Index = %d, expected 2", drv.Index)
	}

	if drv.State != DS_EMPTY_OPEN {
		t.Errorf("State = %d, expected 1", drv.State)
	}

	if drv.Flags != MF_BLURAY_FILES_PRESENT|MF_AACS_FILES_PRESENT {
		t.Errorf("Flags = %d, expected 12", drv.Flags)
	}

	if drv.DriveName != "4815162342" {
		t.Errorf("DriveName = %s, expected \"4815162342\"", drv.DriveName)
	}

	if drv.DiscName != "A_TEAM" {
		t.Errorf("DiscName = %s, expected \"A_TEAM\"", drv.DiscName)
	}

	if drv.Device != "/dev/sr1" {
		t.Errorf("Device = %s, expected \"/dev/sr1\"", drv.Device)
	}
}

func TestParseGeneralMessage(t *testing.T) {
	message, err := ParseMessage("MSG:3007,0,0,\"Using direct disc access mode\",\"Using direct disc access mode\"")
	if err != nil {
		t.Error("ParseMessage returned an error")
	}

	msg, ok := message.(GeneralMessage)
	if !ok {
		t.Error("ParseMessage returned an incorrect type")
	}

	if msg.Code != 3007 {
		t.Errorf("Code = %d, expected 3007", msg.Code)
	}

	if msg.Message != "Using direct disc access mode" {
		t.Errorf("Message = %s, expected \"Using direct disc access mode\"", msg.Message)
	}
}

func TestParseProgressTitleMessage(t *testing.T) {
	msg, err := ParseMessage("PRGC:3400,7,\"Processing AV clips\"")
	if err != nil {
		t.Error("ParseMessage returned an error")
	}

	prgc, ok := msg.(ProgressTitleMessage)
	if !ok {
		t.Error("ParseMessage returned an incorrect type")
	}

	if prgc.Id != 3400 {
		t.Errorf("Id = %d, expected 3400 ", prgc.Id)
	}

	if prgc.Code != 7 {
		t.Errorf("Code = %d, expected 7 ", prgc.Code)
	}

	if prgc.Name != "Processing AV clips" {
		t.Errorf("Name = %s, expected \"Processing AV clips\"", prgc.Name)
	}

	if prgc.Type != 'C' {
		t.Errorf("Type = %d, expected 'C'", prgc.Type)
	}

	msg, err = ParseMessage("PRGT:3404,9,\"Opening Blu-ray disc\"")
	if err != nil {
		t.Error("ParseMessage returned an error")
	}

	prgt, ok := msg.(ProgressTitleMessage)
	if !ok {
		t.Error("ParseMessage returned an incorrect type")
	}

	if prgt.Id != 3404 {
		t.Errorf("Id = %d, expected 3404 ", prgt.Id)
	}

	if prgt.Code != 9 {
		t.Errorf("Code = %d, expected 9 ", prgt.Code)
	}

	if prgt.Name != "Opening Blu-ray disc" {
		t.Errorf("Name = %s, expected \"Opening Blu-ray disc\"", prgt.Name)
	}

	if prgt.Type != 'T' {
		t.Errorf("Type = %d, expected 'C'", prgt.Type)
	}
}

func TestParseProgressValueMessage(t *testing.T) {
	msg, err := ParseMessage("PRGV:30929,21318,65536")
	if err != nil {
		t.Error("ParseMessage returned an error")
	}

	prgv, ok := msg.(ProgressValueMessage)
	if !ok {
		t.Error("ParseMessage returned an incorrect type")
	}

	if prgv.Current != 30929 {
		t.Errorf("Current = %d, expected 30929 ", prgv.Current)
	}

	if prgv.Total != 21318 {
		t.Errorf("Total = %d, expected 21318 ", prgv.Total)
	}

	if prgv.Max != 65536 {
		t.Errorf("Max = %d, expected 65536", prgv.Max)
	}
}

func TestParseStreamInfoMessage(t *testing.T) {
	msg, err := ParseMessage("SINFO:5,1,7,0,\"Dolby Digital\"")
	if err != nil {
		t.Error("ParseMessage returned an error")
	}

	sinfo, ok := msg.(StreamInfoMessage)
	if !ok {
		t.Error("ParseMessage returned an incorrect type")
	}

	if sinfo.Index != 1 {
		t.Errorf("Index = %d, expected 1", sinfo.Index)
	}

	if sinfo.TitleIndex != 5 {
		t.Errorf("TitleIndex = %d, expected 5", sinfo.TitleIndex)
	}

	if sinfo.Attribute.Id != AI_CODEC_LONG {
		t.Errorf("Attribute.Id = %s, expected %s", sinfo.Attribute.Id, AI_CODEC_LONG)
	}

	if sinfo.Attribute.Value != "Dolby Digital" {
		t.Errorf("Attribute.Value = \"%s\", expected \"Dolby Digital\"", sinfo.Attribute.Value)
	}
}

func TestParseTitleCountMessage(t *testing.T) {
	msg, err := ParseMessage("TCOUNT:53")
	if err != nil {
		t.Error("ParseMessage returned an error")
	}

	tcount, ok := msg.(TitleCountMessage)
	if !ok {
		t.Error("ParseMessage returned an incorrect type")
	}

	if tcount.Count != 53 {
		t.Errorf("Count = %d, expected 53", tcount.Count)
	}
}

func TestParseTitleInfoMessage(t *testing.T) {
	msg, err := ParseMessage("TINFO:3,27,0,\"The A-Team_t00.mkv\"")
	if err != nil {
		t.Error("ParseMessage returned an error")
	}

	tinfo, ok := msg.(TitleInfoMessage)
	if !ok {
		t.Error("ParseMessage returned an incorrect type")
	}

	if tinfo.Index != 3 {
		t.Errorf("Index = %d, expected 3", tinfo.Index)
	}

	if tinfo.Attribute.Id != AI_OUTPUT_FILE_NAME {
		t.Errorf("Attribute.Id = %s, expected %s", tinfo.Attribute.Id, AI_OUTPUT_FILE_NAME)
	}

	if tinfo.Attribute.Value != "The A-Team_t00.mkv" {
		t.Errorf("Attribute.Value = \"%s\", expected \"The A-Team_t00.mkv\"", tinfo.Attribute.Value)
	}
}

// TODO: Add bad data for out of range attribute id values (-/+) for CINFO, TINFO, and SINFO
func TestParseMessageErrorHandling(t *testing.T) {
	cases := []string{
		"UNKNOWN:0,0,0",
		"INVALID",
		"CINFO:INVALID,0,\"The A-Team\"",
		"CINFO:5000,0,\"The A-Team\"",
		"CINFO:-500,0,\"The A-Team\"",
		"CINFO:2,0",
		"DRV:INVALID,1,999,12,\"4815162342\",\"A_TEAM\",\"/dev/sr1\"",
		"DRV:2,INVALID,999,12,\"4815162342\",\"A_TEAM\",\"/dev/sr1\"",
		"DRV:2,1,999,INVALID2,\"4815162342\",\"A_TEAM\",\"/dev/sr1\"",
		"DRV:2,1,999,12,\"A_TEAM\",\"/dev/sr1\"",
		"MSG:INVALID,0,0,\"Using direct disc access mode\",\"Using direct disc access mode\"",
		"MSG:3007",
		"PRGC:INVALID,7,\"Processing AV clips\"",
		"PRGC:3400,INVALID,\"Processing AV clips\"",
		"PRGC:3400,7",
		"PRGT:INVALID,9,\"Opening Blu-ray disc\"",
		"PRGT:3404,INVALID,\"Opening Blu-ray disc\"",
		"PRGT:3404,9",
		"PRGV:INVALID,21318,65536",
		"PRGV:30929,INVALID,65536",
		"PRGV:30929,21318,INVALID",
		"PRGV:30929,21318",
		"SINFO:INVALID,1,7,0,\"Dolby Digital\"",
		"SINFO:5,INVALID,7,0,\"Dolby Digital\"",
		"SINFO:5,1,INVALID,0,\"Dolby Digital\"",
		"SINFO:5,1,3000,0,\"Dolby Digital\"",
		"SINFO:5,1,-300,0,\"Dolby Digital\"",
		"SINFO:5",
		"TCOUNT:INVALID",
		"TCOUNT:",
		"TINFO:INVALID,27,0,\"The A-Team_t00.mkv\"",
		"TINFO:3,INVALID,0,\"The A-Team_t00.mkv\"",
		"TINFO:3,2000,0,\"The A-Team_t00.mkv\"",
		"TINFO:3,-200,0,\"The A-Team_t00.mkv\"",
		"TINFO:3",
	}

	for _, line := range cases {
		_, err := ParseMessage(line)
		if err == nil {
			t.Errorf("ParseMessage should have returned an error for input: %s", line)
		}
	}
}
