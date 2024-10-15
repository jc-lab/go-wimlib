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

type PayloadType string

const (
	PayloadProgress PayloadType = "progress"
	PayloadFinish   PayloadType = "finish"
)

type Payload struct {
	Type PayloadType `json:"type"`

	ProgressType *ProgressMsgType `json:"progress_type,omitempty"`
	ProgressInfo interface{}      `json:"progress_info,omitempty"`

	Finish *FinishData `json:"finish,omitempty"`
}

type FinishData struct {
	Success      bool   `json:"success"`
	ErrorMessage string `json:"error_message"`

	// Data
	// - WimInfoData
	Data interface{} `json:"data"`
}

type WimInfoData struct {
	Header *WimInfo `json:"header"`
	// Available Images
}
