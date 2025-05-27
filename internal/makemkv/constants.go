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
	DS_EMPTY_CLOSED DriveState = 0
	DS_EMPTY_OPEN   DriveState = 1
	DS_INSERTED     DriveState = 2
	DS_LOADING      DriveState = 3
	DS_NO_DRIVE     DriveState = 256
	DS_UNMOUNTING   DriveState = 257
)

// MediaFlag represents a flag that can be applied to media inserted into a
// drive. Multiple flags can be applied to a single disc.
type MediaFlag int32

const (
	// Digital Video Disc (DVD)
	MF_DVD_FILES_PRESENT MediaFlag = 1

	// High-Definition Video Disc (HDVD)
	MF_HDVD_FILES_PRESENT MediaFlag = 2

	// Blu-ray Disc
	MF_BLURAY_FILES_PRESENT MediaFlag = 4

	// Advanced Access Content System (AACS)
	MF_AACS_FILES_PRESENT MediaFlag = 8

	// Blu-ray Disc Secure Video Path (BDSVP)
	MF_BDSVM_FILES_PRESENT MediaFlag = 16
)

// AttributeId is an attribute type identifier for attributes reported by
// MakeMKV when reporting information about a disc, its titles, or a title's
// audio/video streams.
//
// The values for the AttributeId constants come from the MakeMKV v1.17.7
// source code in header file apdefs.h
type AttributeId int32

const (
	AI_UNKNOWN                          AttributeId = 0
	AI_TYPE                             AttributeId = 1
	AI_NAME                             AttributeId = 2
	AI_LANG_CODE                        AttributeId = 3
	AI_LANG_NAME                        AttributeId = 4
	AI_CODEC_ID                         AttributeId = 5
	AI_CODEC_SHORT                      AttributeId = 6
	AI_CODEC_LONG                       AttributeId = 7
	AI_CHAPTER_COUNT                    AttributeId = 8
	AI_DURATION                         AttributeId = 9
	AI_DISK_SIZE                        AttributeId = 10
	AI_DISK_SIZE_BYTES                  AttributeId = 11
	AI_STREAM_TYPE_EXTENSION            AttributeId = 12
	AI_BITRATE                          AttributeId = 13
	AI_AUDIO_CHANNELS_COUNT             AttributeId = 14
	AI_ANGLE_INFO                       AttributeId = 15
	AI_SOURCE_FILE_NAME                 AttributeId = 16
	AI_AUDIO_SAMPLE_RATE                AttributeId = 17
	AI_AUDIO_SAMPLE_SIZE                AttributeId = 18
	AI_VIDEO_SIZE                       AttributeId = 19
	AI_VIDEO_ASPECT_RATIO               AttributeId = 20
	AI_VIDEO_FRAME_RATE                 AttributeId = 21
	AI_STREAM_FLAGS                     AttributeId = 22
	AI_DATE_TIME                        AttributeId = 23
	AI_ORIGINAL_TITLE_ID                AttributeId = 24
	AI_SEGMENTS_COUNT                   AttributeId = 25
	AI_SEGMENTS_MAP                     AttributeId = 26
	AI_OUTPUT_FILE_NAME                 AttributeId = 27
	AI_METADATA_LANGUAGE_CODE           AttributeId = 28
	AI_METADATA_LANGUAGE_NAME           AttributeId = 29
	AI_TREE_INFO                        AttributeId = 30
	AI_PANEL_TITLE                      AttributeId = 31
	AI_VOLUME_NAME                      AttributeId = 32
	AI_ORDER_WEIGHT                     AttributeId = 33
	AI_OUTPUT_FORMAT                    AttributeId = 34
	AI_OUTPUT_FORMAT_DESCRIPTION        AttributeId = 35
	AI_SEAMLESS_INFO                    AttributeId = 36
	AI_PANEL_TEXT                       AttributeId = 37
	AI_MKV_FLAGS                        AttributeId = 38
	AI_MKV_FLAGS_TEXT                   AttributeId = 39
	AI_AUDIO_CHANNEL_LAYOUT_NAME        AttributeId = 40
	AI_OUTPUT_CODEC_SHORT               AttributeId = 41
	AI_OUTPUT_CONVERSION_TYPE           AttributeId = 42
	AI_OUTPUT_AUDIO_SAMPLE_RATE         AttributeId = 43
	AI_OUTPUT_AUDIO_SAMPLE_SIZE         AttributeId = 44
	AI_OUTPUT_AUDIO_CHANNELS_COUNT      AttributeId = 45
	AI_OUTPUT_AUDIO_CHANNEL_LAYOUT_NAME AttributeId = 46
	AI_OUTPUT_AUDIO_CHANNEL_LAYOUT      AttributeId = 47
	AI_OUTPUT_AUDIO_MIX_DESCRIPTION     AttributeId = 48
	AI_COMMENT                          AttributeId = 49
	AI_OFFSET_SEQUENCE_ID               AttributeId = 50
)
