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

// DriveState represents different states of an optical drive used by MakeMKV.
//
// The values for the DriveState constants come from the MakeMKV v1.17.7 source
// code in header file apdefs.h
type DriveState int32

const (
	EMPTY_CLOSED DriveState = 0
	EMPTY_OPEN   DriveState = 1
	INSERTED     DriveState = 2
	LOADING      DriveState = 3
	NO_DRIVE     DriveState = 256
	UNMOUNTING   DriveState = 257
)

// MediaFlag represents a flag that can be applied to media inserted into a
// drive. Multiple flags can be applied to a single disc.
type MediaFlag int32

const (
	// Digital Video Disc (DVD)
	DVD_FILES_PRESENT MediaFlag = 1

	// High-Definition Video Disc (HDVD)
	HDVD_FILES_PRESENT MediaFlag = 2

	// Blu-ray Disc
	BLURAY_FILES_PRESENT MediaFlag = 4

	// Advanced Access Content System (AACS)
	AACS_FILES_PRESENT MediaFlag = 8

	// Blu-ray Disc Secure Video Path (BDSVP)
	BDSVM_FILES_PRESENT MediaFlag = 16
)

// AttributeId is an attribute type identifier for attributes reported by
// MakeMKV when reporting information about a disc, its titles, or a title's
// audio/video streams.
//
// The values for the AttributeId constants come from the MakeMKV v1.17.7
// source code in header file apdefs.h
type AttributeId int32

const (
	UNKNOWN                          AttributeId = 0
	TYPE                             AttributeId = 1
	NAME                             AttributeId = 2
	LANG_CODE                        AttributeId = 3
	LANG_NAME                        AttributeId = 4
	CODEC_ID                         AttributeId = 5
	CODEC_SHORT                      AttributeId = 6
	CODEC_LONG                       AttributeId = 7
	CHAPTER_COUNT                    AttributeId = 8
	DURATION                         AttributeId = 9
	DISK_SIZE                        AttributeId = 10
	DISK_SIZE_BYTES                  AttributeId = 11
	STREAM_TYPE_EXTENSION            AttributeId = 12
	BITRATE                          AttributeId = 13
	AUDIO_CHANNELS_COUNT             AttributeId = 14
	ANGLE_INFO                       AttributeId = 15
	SOURCE_FILE_NAME                 AttributeId = 16
	AUDIO_SAMPLE_RATE                AttributeId = 17
	AUDIO_SAMPLE_SIZE                AttributeId = 18
	VIDEO_SIZE                       AttributeId = 19
	VIDEO_ASPECT_RATIO               AttributeId = 20
	VIDEO_FRAME_RATE                 AttributeId = 21
	STREAM_FLAGS                     AttributeId = 22
	DATE_TIME                        AttributeId = 23
	ORIGINAL_TITLE_ID                AttributeId = 24
	SEGMENTS_COUNT                   AttributeId = 25
	SEGMENTS_MAP                     AttributeId = 26
	OUTPUT_FILE_NAME                 AttributeId = 27
	METADATA_LANGUAGE_CODE           AttributeId = 28
	METADATA_LANGUAGE_NAME           AttributeId = 29
	TREE_INFO                        AttributeId = 30
	PANEL_TITLE                      AttributeId = 31
	VOLUME_NAME                      AttributeId = 32
	ORDER_WEIGHT                     AttributeId = 33
	OUTPUT_FORMAT                    AttributeId = 34
	OUTPUT_FORMAT_DESCRIPTION        AttributeId = 35
	SEAMLESS_INFO                    AttributeId = 36
	PANEL_TEXT                       AttributeId = 37
	MKV_FLAGS                        AttributeId = 38
	MKV_FLAGS_TEXT                   AttributeId = 39
	AUDIO_CHANNEL_LAYOUT_NAME        AttributeId = 40
	OUTPUT_CODEC_SHORT               AttributeId = 41
	OUTPUT_CONVERSION_TYPE           AttributeId = 42
	OUTPUT_AUDIO_SAMPLE_RATE         AttributeId = 43
	OUTPUT_AUDIO_SAMPLE_SIZE         AttributeId = 44
	OUTPUT_AUDIO_CHANNELS_COUNT      AttributeId = 45
	OUTPUT_AUDIO_CHANNEL_LAYOUT_NAME AttributeId = 46
	OUTPUT_AUDIO_CHANNEL_LAYOUT      AttributeId = 47
	OUTPUT_AUDIO_MIX_DESCRIPTION     AttributeId = 48
	COMMENT                          AttributeId = 49
	OFFSET_SEQUENCE_ID               AttributeId = 50
)
