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

type ProgressStatus int

const (
	WIMLIB_PROGRESS_STATUS_CONTINUE ProgressStatus = 0
	WIMLIB_PROGRESS_STATUS_ABORT    ProgressStatus = 1
)

const (
	WIMLIB_PROGRESS_MSG_EXTRACT_IMAGE_BEGIN     ProgressMsgType = 0
	WIMLIB_PROGRESS_MSG_EXTRACT_TREE_BEGIN      ProgressMsgType = 1
	WIMLIB_PROGRESS_MSG_EXTRACT_FILE_STRUCTURE  ProgressMsgType = 3
	WIMLIB_PROGRESS_MSG_EXTRACT_STREAMS         ProgressMsgType = 4
	WIMLIB_PROGRESS_MSG_EXTRACT_SPWM_PART_BEGIN ProgressMsgType = 5
	WIMLIB_PROGRESS_MSG_EXTRACT_METADATA        ProgressMsgType = 6
	WIMLIB_PROGRESS_MSG_EXTRACT_IMAGE_END       ProgressMsgType = 7
	WIMLIB_PROGRESS_MSG_EXTRACT_TREE_END        ProgressMsgType = 8
	WIMLIB_PROGRESS_MSG_SCAN_BEGIN              ProgressMsgType = 9
	WIMLIB_PROGRESS_MSG_SCAN_DENTRY             ProgressMsgType = 10
	WIMLIB_PROGRESS_MSG_SCAN_END                ProgressMsgType = 11
	WIMLIB_PROGRESS_MSG_WRITE_STREAMS           ProgressMsgType = 12
	WIMLIB_PROGRESS_MSG_WRITE_METADATA_BEGIN    ProgressMsgType = 13
	WIMLIB_PROGRESS_MSG_WRITE_METADATA_END      ProgressMsgType = 14
	WIMLIB_PROGRESS_MSG_RENAME                  ProgressMsgType = 15
	WIMLIB_PROGRESS_MSG_VERIFY_INTEGRITY        ProgressMsgType = 16
	WIMLIB_PROGRESS_MSG_CALC_INTEGRITY          ProgressMsgType = 17
	WIMLIB_PROGRESS_MSG_SPLIT_BEGIN_PART        ProgressMsgType = 19
	WIMLIB_PROGRESS_MSG_SPLIT_END_PART          ProgressMsgType = 20
	WIMLIB_PROGRESS_MSG_UPDATE_BEGIN_COMMAND    ProgressMsgType = 21
	WIMLIB_PROGRESS_MSG_UPDATE_END_COMMAND      ProgressMsgType = 22
	WIMLIB_PROGRESS_MSG_REPLACE_FILE_IN_WIM     ProgressMsgType = 23
	WIMLIB_PROGRESS_MSG_WIMBOOT_EXCLUDE         ProgressMsgType = 24
	WIMLIB_PROGRESS_MSG_UNMOUNT_BEGIN           ProgressMsgType = 25
	WIMLIB_PROGRESS_MSG_DONE_WITH_FILE          ProgressMsgType = 26
	WIMLIB_PROGRESS_MSG_BEGIN_VERIFY_IMAGE      ProgressMsgType = 27
	WIMLIB_PROGRESS_MSG_END_VERIFY_IMAGE        ProgressMsgType = 28
	WIMLIB_PROGRESS_MSG_VERIFY_STREAMS          ProgressMsgType = 29
	WIMLIB_PROGRESS_MSG_TEST_FILE_EXCLUSION     ProgressMsgType = 30
	WIMLIB_PROGRESS_MSG_HANDLE_ERROR            ProgressMsgType = 31
)
