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
)

const (
	// MAX_STREAM_COUNT is maximum number of streams that can be extracted from a
	// title. This isn't a limit imposed by MakeMKV, but rather a limit imposed by
	// the application to protect against unexpected data causing problems.
	MAX_STREAM_COUNT int = 100

	// MAX_TITLE_COUNT is maximum number of titles that can be extracted from a
	// disc. This isn't a limit imposed by MakeMKV, but rather a limit imposed by
	// the application to protect against unexpected data causing problems.
	MAX_TITLE_COUNT int = 100
)

// StreamInfo is stream (video, audio, or subtitle) information extracted by
// MakeMKV's info command from a DVD or Blu-ray.
type StreamInfo struct {
	Attributes map[AttributeId]string
}

// AddAttribute adds attribute `attr` to the stream returning an error if an
// attribute already exists in the stream.
func (s *StreamInfo) AddAttribute(attr Attribute) error {
	if s.Attributes == nil {
		s.Attributes = make(map[AttributeId]string)
	}

	if _, ok := s.Attributes[attr.Id]; !ok {
		s.Attributes[attr.Id] = attr.Value
		return nil
	} else {
		return errors.New("attribute already exists")
	}
}

// TitleInfo is title information extracted by MakeMKV's info command from a DVD
// or Blu-ray.
type TitleInfo struct {
	Attributes map[AttributeId]string
	Streams    []StreamInfo
}

// AddAttribute adds attribute `attr` to the title returning an error if an
// attribute already exists in the title.
func (t *TitleInfo) AddAttribute(attr Attribute) error {
	if t.Attributes == nil {
		t.Attributes = make(map[AttributeId]string)
	}

	if _, ok := t.Attributes[attr.Id]; !ok {
		t.Attributes[attr.Id] = attr.Value
		return nil
	} else {
		return errors.New("attribute already exists")
	}
}

// AddStreamAttribute adds attribute `attr` to the stream at index `stream` in
// the title returning an error if the index exceeds the maximum stream count or
// if the attribute already exists in the stream.
//
// If the stream at index `stream` doesn't exist, it will be created as well as
// any indexes in between the existing streams and the new stream index.
func (t *TitleInfo) AddStreamAttribute(stream int, attr Attribute) error {
	if stream < 0 {
		return errors.New("title index cannot be negative")
	}

	if stream >= MAX_STREAM_COUNT {
		return errors.New("stream index exceeds stream count limit")
	}

	if t.Streams == nil {
		t.Streams = make([]StreamInfo, 0)
	}

	if stream >= int(len(t.Streams)) {
		for len(t.Streams) <= int(stream) {
			t.Streams = append(t.Streams, StreamInfo{})
		}
	}

	return t.Streams[stream].AddAttribute(attr)
}

// DiscInfo is disc information extracted by MakeMKV's info command from a DVD
// or Blu-ray.
type DiscInfo struct {
	TitleCount int
	Attributes map[AttributeId]string
	Titles     []TitleInfo
}

// AddAttribute adds attribute `attr` to the title returning an error if an
// attribute already exists in the title.
func (d *DiscInfo) AddAttribute(attr Attribute) error {
	if d.Attributes == nil {
		d.Attributes = make(map[AttributeId]string)
	}

	if _, ok := d.Attributes[attr.Id]; !ok {
		d.Attributes[attr.Id] = attr.Value
		return nil
	} else {
		return errors.New("attribute already exists")
	}
}

// AddTitleAttribute adds attribute `attr` to the stream at index `title` in the
// title returning an error if the index exceeds the maximum stream count or if
// the attribute already exists in the stream.
func (d *DiscInfo) AddTitleAttribute(title int, attr Attribute) error {
	if err := d.ensureTitleExists(title); err != nil {
		return err
	} else {
		return d.Titles[title].AddAttribute(attr)
	}
}

// AddStreamAttribute adds attribute `attr` to the stream at index `stream` in
// the title at index `title` returning an error if the index exceeds the maximum
// stream count or if the attribute already exists in the stream.
func (d *DiscInfo) AddStreamAttribute(stream, title int, attr Attribute) error {
	if err := d.ensureTitleExists(title); err != nil {
		return err
	} else {
		return d.Titles[title].AddStreamAttribute(stream, attr)
	}
}

func (d *DiscInfo) ensureTitleExists(title int) error {
	if title < 0 {
		return errors.New("title index cannot be negative")
	}

	if title >= MAX_TITLE_COUNT {
		return errors.New("title index exceeds stream count limit")
	}

	if d.Titles == nil {
		d.Titles = make([]TitleInfo, 0)
	}

	if title >= len(d.Titles) {
		for len(d.Titles) <= int(title) {
			d.Titles = append(d.Titles, TitleInfo{})
		}
	}

	return nil
}
