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
type AttributeId string

const (
	AI_UNKNOWN                          AttributeId = "UNKNOWN"                          // 0
	AI_TYPE                             AttributeId = "TYPE"                             // 1
	AI_NAME                             AttributeId = "NAME"                             // 2
	AI_LANG_CODE                        AttributeId = "LANG_CODE"                        // 3
	AI_LANG_NAME                        AttributeId = "LANG_NAME"                        // 4
	AI_CODEC_ID                         AttributeId = "CODEC_ID"                         // 5
	AI_CODEC_SHORT                      AttributeId = "CODEC_SHORT"                      // 6
	AI_CODEC_LONG                       AttributeId = "CODEC_LONG"                       // 7
	AI_CHAPTER_COUNT                    AttributeId = "CHAPTER_COUNT"                    // 8
	AI_DURATION                         AttributeId = "DURATION"                         // 9
	AI_DISK_SIZE                        AttributeId = "DISK_SIZE"                        // 10
	AI_DISK_SIZE_BYTES                  AttributeId = "DISK_SIZE_BYTES"                  // 11
	AI_STREAM_TYPE_EXTENSION            AttributeId = "STREAM_TYPE_EXTENSION"            // 12
	AI_BITRATE                          AttributeId = "BITRATE"                          // 13
	AI_AUDIO_CHANNELS_COUNT             AttributeId = "AUDIO_CHANNELS_COUNT"             // 14
	AI_ANGLE_INFO                       AttributeId = "ANGLE_INFO"                       // 15
	AI_SOURCE_FILE_NAME                 AttributeId = "SOURCE_FILE_NAME"                 // 16
	AI_AUDIO_SAMPLE_RATE                AttributeId = "AUDIO_SAMPLE_RATE"                // 17
	AI_AUDIO_SAMPLE_SIZE                AttributeId = "AUDIO_SAMPLE_SIZE"                // 18
	AI_VIDEO_SIZE                       AttributeId = "VIDEO_SIZE"                       // 19
	AI_VIDEO_ASPECT_RATIO               AttributeId = "VIDEO_ASPECT_RATIO"               // 20
	AI_VIDEO_FRAME_RATE                 AttributeId = "VIDEO_FRAME_RATE"                 // 21
	AI_STREAM_FLAGS                     AttributeId = "STREAM_FLAGS"                     // 22
	AI_DATE_TIME                        AttributeId = "DATE_TIME"                        // 23
	AI_ORIGINAL_TITLE_ID                AttributeId = "ORIGINAL_TITLE_ID"                // 24
	AI_SEGMENTS_COUNT                   AttributeId = "SEGMENTS_COUNT"                   // 25
	AI_SEGMENTS_MAP                     AttributeId = "SEGMENTS_MAP"                     // 26
	AI_OUTPUT_FILE_NAME                 AttributeId = "OUTPUT_FILE_NAME"                 // 27
	AI_METADATA_LANGUAGE_CODE           AttributeId = "METADATA_LANGUAGE_CODE"           // 28
	AI_METADATA_LANGUAGE_NAME           AttributeId = "METADATA_LANGUAGE_NAME"           // 29
	AI_TREE_INFO                        AttributeId = "TREE_INFO"                        // 30
	AI_PANEL_TITLE                      AttributeId = "PANEL_TITLE"                      // 31
	AI_VOLUME_NAME                      AttributeId = "VOLUME_NAME"                      // 32
	AI_ORDER_WEIGHT                     AttributeId = "ORDER_WEIGHT"                     // 33
	AI_OUTPUT_FORMAT                    AttributeId = "OUTPUT_FORMAT"                    // 34
	AI_OUTPUT_FORMAT_DESCRIPTION        AttributeId = "OUTPUT_FORMAT_DESCRIPTION"        // 35
	AI_SEAMLESS_INFO                    AttributeId = "SEAMLESS_INFO"                    // 36
	AI_PANEL_TEXT                       AttributeId = "PANEL_TEXT"                       // 37
	AI_MKV_FLAGS                        AttributeId = "MKV_FLAGS"                        // 38
	AI_MKV_FLAGS_TEXT                   AttributeId = "MKV_FLAGS_TEXT"                   // 39
	AI_AUDIO_CHANNEL_LAYOUT_NAME        AttributeId = "AUDIO_CHANNEL_LAYOUT_NAME"        // 40
	AI_OUTPUT_CODEC_SHORT               AttributeId = "OUTPUT_CODEC_SHORT"               // 41
	AI_OUTPUT_CONVERSION_TYPE           AttributeId = "OUTPUT_CONVERSION_TYPE"           // 42
	AI_OUTPUT_AUDIO_SAMPLE_RATE         AttributeId = "OUTPUT_AUDIO_SAMPLE_RATE"         // 43
	AI_OUTPUT_AUDIO_SAMPLE_SIZE         AttributeId = "OUTPUT_AUDIO_SAMPLE_SIZE"         // 44
	AI_OUTPUT_AUDIO_CHANNELS_COUNT      AttributeId = "OUTPUT_AUDIO_CHANNELS_COUNT"      // 45
	AI_OUTPUT_AUDIO_CHANNEL_LAYOUT_NAME AttributeId = "OUTPUT_AUDIO_CHANNEL_LAYOUT_NAME" // 46
	AI_OUTPUT_AUDIO_CHANNEL_LAYOUT      AttributeId = "OUTPUT_AUDIO_CHANNEL_LAYOUT"      // 47
	AI_OUTPUT_AUDIO_MIX_DESCRIPTION     AttributeId = "OUTPUT_AUDIO_MIX_DESCRIPTION"     // 48
	AI_COMMENT                          AttributeId = "COMMENT"                          // 49
	AI_OFFSET_SEQUENCE_ID               AttributeId = "OFFSET_SEQUENCE_ID"               // 50
)

// attributeTable maps the values reported by MakeMKV for attribute IDs to one
// of the AttributeId constants. The values for the AttributeId constants come
// from the MakeMKV v1.17.7 source code in header file apdefs.h
var attributeTable = map[int]AttributeId{
	0:  AI_UNKNOWN,
	1:  AI_TYPE,
	2:  AI_NAME,
	3:  AI_LANG_CODE,
	4:  AI_LANG_NAME,
	5:  AI_CODEC_ID,
	6:  AI_CODEC_SHORT,
	7:  AI_CODEC_LONG,
	8:  AI_CHAPTER_COUNT,
	9:  AI_DURATION,
	10: AI_DISK_SIZE,
	11: AI_DISK_SIZE_BYTES,
	12: AI_STREAM_TYPE_EXTENSION,
	13: AI_BITRATE,
	14: AI_AUDIO_CHANNELS_COUNT,
	15: AI_ANGLE_INFO,
	16: AI_SOURCE_FILE_NAME,
	17: AI_AUDIO_SAMPLE_RATE,
	18: AI_AUDIO_SAMPLE_SIZE,
	19: AI_VIDEO_SIZE,
	20: AI_VIDEO_ASPECT_RATIO,
	21: AI_VIDEO_FRAME_RATE,
	22: AI_STREAM_FLAGS,
	23: AI_DATE_TIME,
	24: AI_ORIGINAL_TITLE_ID,
	25: AI_SEGMENTS_COUNT,
	26: AI_SEGMENTS_MAP,
	27: AI_OUTPUT_FILE_NAME,
	28: AI_METADATA_LANGUAGE_CODE,
	29: AI_METADATA_LANGUAGE_NAME,
	30: AI_TREE_INFO,
	31: AI_PANEL_TITLE,
	32: AI_VOLUME_NAME,
	33: AI_ORDER_WEIGHT,
	34: AI_OUTPUT_FORMAT,
	35: AI_OUTPUT_FORMAT_DESCRIPTION,
	36: AI_SEAMLESS_INFO,
	37: AI_PANEL_TEXT,
	38: AI_MKV_FLAGS,
	39: AI_MKV_FLAGS_TEXT,
	40: AI_AUDIO_CHANNEL_LAYOUT_NAME,
	41: AI_OUTPUT_CODEC_SHORT,
	42: AI_OUTPUT_CONVERSION_TYPE,
	43: AI_OUTPUT_AUDIO_SAMPLE_RATE,
	44: AI_OUTPUT_AUDIO_SAMPLE_SIZE,
	45: AI_OUTPUT_AUDIO_CHANNELS_COUNT,
	46: AI_OUTPUT_AUDIO_CHANNEL_LAYOUT_NAME,
	47: AI_OUTPUT_AUDIO_CHANNEL_LAYOUT,
	48: AI_OUTPUT_AUDIO_MIX_DESCRIPTION,
	49: AI_COMMENT,
	50: AI_OFFSET_SEQUENCE_ID,
}

// GetAttributeId returns the AttributeId associated with value `v` outputted by
// MakeMKV.
func GetAttributeId(v int) (AttributeId, bool) {
	id, ok := attributeTable[v]
	return id, ok
}
