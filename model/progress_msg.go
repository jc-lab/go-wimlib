package model

import (
	"fmt"
	"github.com/vmihailenco/msgpack/v5"
)

//go:generate msgp

type ProgressMsgType int

type ProgressMsg struct {
	MsgType ProgressMsgType        `msg:"msg_type" json:"msg_type"`
	Info    map[string]interface{} `msg:"info" json:"info"`
}

// ProgressWriteStreams structure
type ProgressWriteStreams struct {
	TotalBytes               uint64 `msg:"total_bytes" json:"total_bytes"`
	TotalStreams             uint64 `msg:"total_streams" json:"total_streams"`
	CompletedBytes           uint64 `msg:"completed_bytes" json:"completed_bytes"`
	CompletedStreams         uint64 `msg:"completed_streams" json:"completed_streams"`
	NumThreads               uint32 `msg:"num_threads" json:"num_threads"`
	CompressionType          int32  `msg:"compression_type" json:"compression_type"`
	TotalParts               uint32 `msg:"total_parts" json:"total_parts"`
	CompletedParts           uint32 `msg:"completed_parts" json:"completed_parts"`
	CompletedCompressedBytes uint64 `msg:"completed_compressed_bytes" json:"completed_compressed_bytes"`
}

// ProgressScan structure
type ProgressScan struct {
	Source            string `msg:"source" json:"source"`
	CurPath           string `msg:"cur_path" json:"cur_path"`
	Status            int    `msg:"status" json:"status"`
	WimTargetPath     string `msg:"wim_target_path" json:"wim_target_path"`
	NumDirsScanned    uint64 `msg:"num_dirs_scanned" json:"num_dirs_scanned"`
	NumNondirsScanned uint64 `msg:"num_nondirs_scanned" json:"num_nondirs_scanned"`
	NumBytesScanned   uint64 `msg:"num_bytes_scanned" json:"num_bytes_scanned"`
}

// ProgressExtract structure
type ProgressExtract struct {
	Image            uint32 `msg:"image" json:"image"`
	ExtractFlags     uint32 `msg:"extract_flags" json:"extract_flags"`
	WimfileName      string `msg:"wimfile_name" json:"wimfile_name"`
	ImageName        string `msg:"image_name" json:"image_name"`
	Target           string `msg:"target" json:"target"`
	TotalBytes       uint64 `msg:"total_bytes" json:"total_bytes"`
	CompletedBytes   uint64 `msg:"completed_bytes" json:"completed_bytes"`
	TotalStreams     uint64 `msg:"total_streams" json:"total_streams"`
	CompletedStreams uint64 `msg:"completed_streams" json:"completed_streams"`
	PartNumber       uint32 `msg:"part_number" json:"part_number"`
	TotalParts       uint32 `msg:"total_parts" json:"total_parts"`
	Guid             []byte `msg:"guid" json:"guid"`
	CurrentFileCount uint64 `msg:"current_file_count" json:"current_file_count"`
}

// ProgressRename structure
type ProgressRename struct {
	From string `msg:"from" json:"from"`
	To   string `msg:"to" json:"to"`
}

// ProgressUpdate structure
type ProgressUpdate struct {
	Command           uint64 `msg:"command" json:"command"`
	CompletedCommands uint64 `msg:"completed_commands" json:"completed_commands"`
	TotalCommands     uint64 `msg:"total_commands" json:"total_commands"`
}

// ProgressIntegrity structure
type ProgressIntegrity struct {
	TotalBytes      uint64 `msg:"total_bytes" json:"total_bytes"`
	CompletedBytes  uint64 `msg:"completed_bytes" json:"completed_bytes"`
	TotalChunks     uint32 `msg:"total_chunks" json:"total_chunks"`
	CompletedChunks uint32 `msg:"completed_chunks" json:"completed_chunks"`
	ChunkSize       uint32 `msg:"chunk_size" json:"chunk_size"`
	Filename        string `msg:"filename" json:"filename"`
}

// ProgressSplit structure
type ProgressSplit struct {
	TotalBytes     uint64 `msg:"total_bytes" json:"total_bytes"`
	CompletedBytes uint64 `msg:"completed_bytes" json:"completed_bytes"`
	CurPartNumber  uint   `msg:"cur_part_number" json:"cur_part_number"`
	TotalParts     uint   `msg:"total_parts" json:"total_parts"`
	PartName       string `msg:"part_name" json:"part_name"`
}

// ProgressReplaceFileInWim structure
type ProgressReplaceFileInWim struct {
	PathInWim string `msg:"path_in_wim" json:"path_in_wim"`
}

// ProgressWimbootExclude structure
type ProgressWimbootExclude struct {
	PathInWim      string `msg:"path_in_wim" json:"path_in_wim"`
	ExtractionPath string `msg:"extraction_path" json:"extraction_path"`
}

// ProgressUnmount structure
type ProgressUnmount struct {
	Mountpoint   string `msg:"mountpoint" json:"mountpoint"`
	MountedWim   string `msg:"mounted_wim" json:"mounted_wim"`
	MountedImage uint32 `msg:"mounted_image" json:"mounted_image"`
	MountFlags   uint32 `msg:"mount_flags" json:"mount_flags"`
	UnmountFlags uint32 `msg:"unmount_flags" json:"unmount_flags"`
}

// ProgressDoneWithFile structure
type ProgressDoneWithFile struct {
	PathToFile string `msg:"path_to_file" json:"path_to_file"`
}

// ProgressVerifyImage structure
type ProgressVerifyImage struct {
	Wimfile      string `msg:"wimfile" json:"wimfile"`
	TotalImages  uint32 `msg:"total_images" json:"total_images"`
	CurrentImage uint32 `msg:"current_image" json:"current_image"`
}

// ProgressVerifyStreams structure
type ProgressVerifyStreams struct {
	Wimfile          string `msg:"wimfile" json:"wimfile"`
	TotalStreams     uint64 `msg:"total_streams" json:"total_streams"`
	TotalBytes       uint64 `msg:"total_bytes" json:"total_bytes"`
	CompletedStreams uint64 `msg:"completed_streams" json:"completed_streams"`
	CompletedBytes   uint64 `msg:"completed_bytes" json:"completed_bytes"`
}

// ProgressTestFileExclusion structure
type ProgressTestFileExclusion struct {
	Path        string `msg:"path" json:"path"`
	WillExclude bool   `msg:"will_exclude" json:"will_exclude"`
}

// ProgressHandleError structure
type ProgressHandleError struct {
	Path       string `msg:"path" json:"path"`
	ErrorCode  int    `msg:"error_code" json:"error_code"`
	WillIgnore bool   `msg:"will_ignore" json:"will_ignore"`
}

// GetInfo method for ProgressMsg
func (m *ProgressMsg) GetInfo() (interface{}, error) {
	// Marshal the `Info` map to msgpack byte format
	data, err := msgpack.Marshal(m.Info)
	if err != nil {
		return nil, err
	}

	// Declare a variable to hold the unmarshaled structure
	var result interface{}

	// Use switch to determine which structure to unmarshal into based on MsgType
	switch m.MsgType {
	case WIMLIB_PROGRESS_MSG_WRITE_STREAMS:
		var writeStreams ProgressWriteStreams
		if _, err := writeStreams.UnmarshalMsg(data); err != nil {
			return nil, err
		}
		result = &writeStreams

	case WIMLIB_PROGRESS_MSG_SCAN_BEGIN, WIMLIB_PROGRESS_MSG_SCAN_DENTRY, WIMLIB_PROGRESS_MSG_SCAN_END:
		var scan ProgressScan
		if _, err := scan.UnmarshalMsg(data); err != nil {
			return nil, err
		}
		result = &scan

	case WIMLIB_PROGRESS_MSG_EXTRACT_SPWM_PART_BEGIN, WIMLIB_PROGRESS_MSG_EXTRACT_IMAGE_BEGIN,
		WIMLIB_PROGRESS_MSG_EXTRACT_TREE_BEGIN, WIMLIB_PROGRESS_MSG_EXTRACT_FILE_STRUCTURE,
		WIMLIB_PROGRESS_MSG_EXTRACT_STREAMS, WIMLIB_PROGRESS_MSG_EXTRACT_METADATA,
		WIMLIB_PROGRESS_MSG_EXTRACT_TREE_END, WIMLIB_PROGRESS_MSG_EXTRACT_IMAGE_END:
		var extract ProgressExtract
		if _, err := extract.UnmarshalMsg(data); err != nil {
			return nil, err
		}
		result = &extract

	case WIMLIB_PROGRESS_MSG_RENAME:
		var rename ProgressRename
		if _, err := rename.UnmarshalMsg(data); err != nil {
			return nil, err
		}
		result = &rename

	case WIMLIB_PROGRESS_MSG_UPDATE_BEGIN_COMMAND, WIMLIB_PROGRESS_MSG_UPDATE_END_COMMAND:
		var update ProgressUpdate
		if _, err := update.UnmarshalMsg(data); err != nil {
			return nil, err
		}
		result = &update

	case WIMLIB_PROGRESS_MSG_VERIFY_INTEGRITY, WIMLIB_PROGRESS_MSG_CALC_INTEGRITY:
		var integrity ProgressIntegrity
		if _, err := integrity.UnmarshalMsg(data); err != nil {
			return nil, err
		}
		result = &integrity

	case WIMLIB_PROGRESS_MSG_SPLIT_BEGIN_PART, WIMLIB_PROGRESS_MSG_SPLIT_END_PART:
		var split ProgressSplit
		if _, err := split.UnmarshalMsg(data); err != nil {
			return nil, err
		}
		result = &split

	case WIMLIB_PROGRESS_MSG_REPLACE_FILE_IN_WIM:
		var replaceFile ProgressReplaceFileInWim
		if _, err := replaceFile.UnmarshalMsg(data); err != nil {
			return nil, err
		}
		result = &replaceFile

	case WIMLIB_PROGRESS_MSG_WIMBOOT_EXCLUDE:
		var wimbootExclude ProgressWimbootExclude
		if _, err := wimbootExclude.UnmarshalMsg(data); err != nil {
			return nil, err
		}
		result = &wimbootExclude

	case WIMLIB_PROGRESS_MSG_UNMOUNT_BEGIN:
		var unmount ProgressUnmount
		if _, err := unmount.UnmarshalMsg(data); err != nil {
			return nil, err
		}
		result = &unmount

	case WIMLIB_PROGRESS_MSG_DONE_WITH_FILE:
		var doneWithFile ProgressDoneWithFile
		if _, err := doneWithFile.UnmarshalMsg(data); err != nil {
			return nil, err
		}
		result = &doneWithFile

	case WIMLIB_PROGRESS_MSG_BEGIN_VERIFY_IMAGE, WIMLIB_PROGRESS_MSG_END_VERIFY_IMAGE:
		var verifyImage ProgressVerifyImage
		if _, err := verifyImage.UnmarshalMsg(data); err != nil {
			return nil, err
		}
		result = &verifyImage

	case WIMLIB_PROGRESS_MSG_VERIFY_STREAMS:
		var verifyStreams ProgressVerifyStreams
		if _, err := verifyStreams.UnmarshalMsg(data); err != nil {
			return nil, err
		}
		result = &verifyStreams

	case WIMLIB_PROGRESS_MSG_TEST_FILE_EXCLUSION:
		var testFileExclusion ProgressTestFileExclusion
		if _, err := testFileExclusion.UnmarshalMsg(data); err != nil {
			return nil, err
		}
		result = &testFileExclusion

	case WIMLIB_PROGRESS_MSG_HANDLE_ERROR:
		var handleError ProgressHandleError
		if _, err := handleError.UnmarshalMsg(data); err != nil {
			return nil, err
		}
		result = &handleError

	default:
		return nil, fmt.Errorf("unknown message type: %v", m.MsgType)
	}

	return result, nil
}
