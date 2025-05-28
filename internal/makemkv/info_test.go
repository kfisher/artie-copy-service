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

func TestAddAttributeToStreamInfo(t *testing.T) {
	attr := Attribute{Id: AI_BITRATE, Value: "224 Kb/s"}

	stream := StreamInfo{}

	if err := stream.AddAttribute(attr); err != nil {
		t.Errorf("AddAttribute returned an unexpected error. err = %s", err)
		return
	}

	if v, ok := stream.Attributes[attr.Id]; ok {
		if v != attr.Value {
			t.Errorf("Attribute %s has value %s, expected %s", attr.Id, v, attr.Value)
		}
	} else {
		t.Errorf("Attribute %s not found in stream", attr.Id)
	}

	if err := stream.AddAttribute(attr); err == nil {
		t.Error("AddAttribute did not return an error for duplicate attribute")
	}
}

func TestAddAttributeToTitleInfo(t *testing.T) {
	attr := Attribute{Id: AI_DISK_SIZE, Value: "26.4 GB"}

	title := TitleInfo{}

	if err := title.AddAttribute(attr); err != nil {
		t.Errorf("AddAttribute returned an unexpected error. err = %s", err)
		return
	}

	if v, ok := title.Attributes[attr.Id]; ok {
		if v != attr.Value {
			t.Errorf("Attribute %s has value %s, expected %s", attr.Id, v, attr.Value)
		}
	} else {
		t.Errorf("Attribute %s not found in stream", attr.Id)
	}

	if err := title.AddAttribute(attr); err == nil {
		t.Error("AddAttribute did not return an error for duplicate attribute")
	}
}

func TestAddStreamAttributeToTitleInfo(t *testing.T) {
	attr := Attribute{Id: AI_BITRATE, Value: "224 Kb/s"}

	title := TitleInfo{}

	if err := title.AddStreamAttribute(1, attr); err != nil {
		t.Errorf("AddStreamAttribute returned an unexpected error. err = %s", err)
		return
	}

	if len(title.Streams) != 2 {
		t.Errorf("Expected 2 streams, got %d", len(title.Streams))
		return
	}

	if title.Streams[0].Attributes != nil {
		t.Errorf("Stream 0 should not have attributes, but found %d", len(title.Streams[0].Attributes))
	}

	if v, ok := title.Streams[1].Attributes[attr.Id]; ok {
		if v != attr.Value {
			t.Errorf("Attribute %s has value %s, expected %s", attr.Id, v, attr.Value)
		}
	} else {
		t.Errorf("Attribute %s not found in stream", attr.Id)
	}

	if err := title.AddStreamAttribute(MAX_STREAM_COUNT+1, Attribute{AI_ANGLE_INFO, ""}); err == nil {
		t.Error("AddStreamAttribute did not return an error for out of bounds index")
	}
}

func TestAddAttributeToDiscInfo(t *testing.T) {
	attr := Attribute{Id: AI_METADATA_LANGUAGE_NAME, Value: "English"}

	disc := DiscInfo{}

	if err := disc.AddAttribute(attr); err != nil {
		t.Errorf("AddAttribute returned an unexpected error. err = %s", err)
		return
	}

	if v, ok := disc.Attributes[attr.Id]; ok {
		if v != attr.Value {
			t.Errorf("Attribute %s has value %s, expected %s", attr.Id, v, attr.Value)
		}
	} else {
		t.Errorf("Attribute %s not found in stream", attr.Id)
	}

	if err := disc.AddAttribute(attr); err == nil {
		t.Error("AddAttribute did not return an error for duplicate attribute")
	}
}

func TestAddTitleAttributeToDiscInfo(t *testing.T) {
	attr := Attribute{Id: AI_DISK_SIZE, Value: "26.4 GB"}

	disc := DiscInfo{}

	if err := disc.AddTitleAttribute(1, attr); err != nil {
		t.Errorf("AddTitleAttribute returned an unexpected error. err = %s", err)
		return
	}

	if len(disc.Titles) != 2 {
		t.Errorf("Expected 2 titles, got %d", len(disc.Titles))
		return
	}

	if disc.Titles[0].Attributes != nil {
		t.Errorf("Title 0 should not have attributes, but found %d", len(disc.Titles[0].Attributes))
	}

	if v, ok := disc.Titles[1].Attributes[attr.Id]; ok {
		if v != attr.Value {
			t.Errorf("Attribute %s has value %s, expected %s", attr.Id, v, attr.Value)
		}
	} else {
		t.Errorf("Attribute %s not found in title", attr.Id)
	}

	if err := disc.AddTitleAttribute(MAX_TITLE_COUNT+1, Attribute{AI_ANGLE_INFO, ""}); err == nil {
		t.Error("AddTitleAttribute did not return an error for out of bounds index")
	}
}

func TestAddStreamAttributeToDiscInfo(t *testing.T) {
	attr := Attribute{Id: AI_BITRATE, Value: "224 Kb/s"}

	disc := DiscInfo{}

	if err := disc.AddStreamAttribute(0, 1, attr); err != nil {
		t.Errorf("AddStreamAttribute returned an unexpected error. err = %s", err)
		return
	}

	if len(disc.Titles) != 2 {
		t.Errorf("Expected 2 titles, got %d", len(disc.Titles))
		return
	}

	if len(disc.Titles[1].Streams) != 1 {
		t.Errorf("Expected 1 stream, got %d", len(disc.Titles))
		return
	}

	if v, ok := disc.Titles[1].Streams[0].Attributes[attr.Id]; ok {
		if v != attr.Value {
			t.Errorf("Attribute %s has value %s, expected %s", attr.Id, v, attr.Value)
		}
	} else {
		t.Errorf("Attribute %s not found in stream", attr.Id)
	}

	if err := disc.AddStreamAttribute(MAX_STREAM_COUNT+1, 0, Attribute{AI_ANGLE_INFO, ""}); err == nil {
		t.Error("AddStreamAttribute did not return an error for out of bounds index")
	}

	if err := disc.AddStreamAttribute(0, MAX_TITLE_COUNT+1, Attribute{AI_ANGLE_INFO, ""}); err == nil {
		t.Error("AddStreamAttribute did not return an error for out of bounds index")
	}
}
