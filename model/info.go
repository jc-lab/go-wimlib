// Copyright 2024 JC-Lab
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package model

import "C"
import (
	"encoding/base64"
	"encoding/hex"
)

type WimGuid [WIMLIB_GUID_LEN]byte

type WimInfo struct {
	Guid              WimGuid               `json:"guid"`
	ImageCount        uint32                `json:"image_count"`
	BootIndex         uint32                `json:"boot_index"`
	WimVersion        uint32                `json:"wim_version"`
	ChunkSize         uint32                `json:"chunk_size"`
	PartNumber        uint16                `json:"part_number"`
	TotalParts        uint16                `json:"total_parts"`
	CompressionType   WimlibCompressionType `json:"compression_type"`
	TotalBytes        uint64                `json:"total_bytes"`
	HasIntegrityTable bool                  `json:"has_integrity_table"`
	OpenedFromFile    bool                  `json:"opened_from_file"`
	IsReadonly        bool                  `json:"is_readonly"`
	HasRpfix          bool                  `json:"has_rpfix"`
	IsMarkedReadonly  bool                  `json:"is_marked_readonly"`
	Spanned           bool                  `json:"spanned"`
	WriteInProgress   bool                  `json:"write_in_progress"`
	MetadataOnly      bool                  `json:"metadata_only"`
	ResourceOnly      bool                  `json:"resource_only"`
	Pipable           bool                  `json:"pipable"`
}

func (g WimGuid) String() string {
	return "0x" + hex.EncodeToString(g[:])
}

func (g WimGuid) MarshalJSON() ([]byte, error) {
	return []byte("\"" + base64.StdEncoding.EncodeToString(g[:]) + "\""), nil
}
