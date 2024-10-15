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

type WimlibFileAttribute = int
type WimlibAddFlag = int
type WimlibDeleteFlag = int
type WimlibExportFlag = int
type WimlibExtractFlag = int
type WimlibMountFlag = int
type WimlibUnmountFlag = int
type WimlibUpdateFlag = int
type WimlibOpenFlag = int
type WimlibWriteFlag = int
type WimlibInitFlag = int
type WimlibRefFlag = int

const (
	WIMLIB_GUID_LEN                                                                  = 16
	WIMLIB_CHANGE_READONLY_FLAG                                                      = 0x00000001
	WIMLIB_CHANGE_GUID                                                               = 0x00000002
	WIMLIB_CHANGE_BOOT_INDEX                                                         = 0x00000004
	WIMLIB_CHANGE_RPFIX_FLAG                                                         = 0x00000008
	WIMLIB_FILE_ATTRIBUTE_READONLY                               WimlibFileAttribute = 0x00000001
	WIMLIB_FILE_ATTRIBUTE_HIDDEN                                 WimlibFileAttribute = 0x00000002
	WIMLIB_FILE_ATTRIBUTE_SYSTEM                                 WimlibFileAttribute = 0x00000004
	WIMLIB_FILE_ATTRIBUTE_DIRECTORY                              WimlibFileAttribute = 0x00000010
	WIMLIB_FILE_ATTRIBUTE_ARCHIVE                                WimlibFileAttribute = 0x00000020
	WIMLIB_FILE_ATTRIBUTE_DEVICE                                 WimlibFileAttribute = 0x00000040
	WIMLIB_FILE_ATTRIBUTE_NORMAL                                 WimlibFileAttribute = 0x00000080
	WIMLIB_FILE_ATTRIBUTE_TEMPORARY                              WimlibFileAttribute = 0x00000100
	WIMLIB_FILE_ATTRIBUTE_SPARSE_FILE                            WimlibFileAttribute = 0x00000200
	WIMLIB_FILE_ATTRIBUTE_REPARSE_POINT                          WimlibFileAttribute = 0x00000400
	WIMLIB_FILE_ATTRIBUTE_COMPRESSED                             WimlibFileAttribute = 0x00000800
	WIMLIB_FILE_ATTRIBUTE_OFFLINE                                WimlibFileAttribute = 0x00001000
	WIMLIB_FILE_ATTRIBUTE_NOT_CONTENT_INDEXEDWimlibFileAttribute WimlibFileAttribute = 0x00002000
	WIMLIB_FILE_ATTRIBUTE_ENCRYPTED                              WimlibFileAttribute = 0x00004000
	WIMLIB_FILE_ATTRIBUTE_VIRTUAL                                WimlibFileAttribute = 0x00010000
	WIMLIB_REPARSE_TAG_RESERVED_ZERO                                                 = 0x00000000
	WIMLIB_REPARSE_TAG_RESERVED_ONE                                                  = 0x00000001
	WIMLIB_REPARSE_TAG_MOUNT_POINT                                                   = 0xA0000003
	WIMLIB_REPARSE_TAG_HSM                                                           = 0xC0000004
	WIMLIB_REPARSE_TAG_HSM2                                                          = 0x80000006
	WIMLIB_REPARSE_TAG_DRIVER_EXTENDER                                               = 0x80000005
	WIMLIB_REPARSE_TAG_SIS                                                           = 0x80000007
	WIMLIB_REPARSE_TAG_DFS                                                           = 0x8000000A
	WIMLIB_REPARSE_TAG_DFSR                                                          = 0x80000012
	WIMLIB_REPARSE_TAG_FILTER_MANAGER                                                = 0x8000000B
	WIMLIB_REPARSE_TAG_WOF                                                           = 0x80000017
	WIMLIB_REPARSE_TAG_SYMLINK                                                       = 0xA000000C
	WIMLIB_ITERATE_DIR_TREE_FLAG_RECURSIVE                                           = 0x00000001
	WIMLIB_ITERATE_DIR_TREE_FLAG_CHILDREN                                            = 0x00000002
	WIMLIB_ITERATE_DIR_TREE_FLAG_RESOURCES_NEEDED                                    = 0x00000004
	WIMLIB_ADD_FLAG_NTFS                                         WimlibAddFlag       = 0x00000001
	WIMLIB_ADD_FLAG_DEREFERENCE                                  WimlibAddFlag       = 0x00000002
	WIMLIB_ADD_FLAG_VERBOSE                                      WimlibAddFlag       = 0x00000004
	WIMLIB_ADD_FLAG_BOOT                                         WimlibAddFlag       = 0x00000008
	WIMLIB_ADD_FLAG_UNIX_DATA                                    WimlibAddFlag       = 0x00000010
	WIMLIB_ADD_FLAG_NO_ACLSWimlibAddFlag                         WimlibAddFlag       = 0x00000020
	WIMLIB_ADD_FLAG_STRICT_ACLS                                  WimlibAddFlag       = 0x00000040
	WIMLIB_ADD_FLAG_EXCLUDE_VERBOSE                              WimlibAddFlag       = 0x00000080
	WIMLIB_ADD_FLAG_RPFIX                                        WimlibAddFlag       = 0x00000100
	WIMLIB_ADD_FLAG_NORPFIX                                      WimlibAddFlag       = 0x00000200
	WIMLIB_ADD_FLAG_NO_UNSUPPORTED_EXCLUDEWimlibAddFlag          WimlibAddFlag       = 0x00000400
	WIMLIB_ADD_FLAG_WINCONFIG                                    WimlibAddFlag       = 0x00000800
	WIMLIB_ADD_FLAG_WIMBOOT                                      WimlibAddFlag       = 0x00001000
	WIMLIB_ADD_FLAG_NO_REPLACE                                   WimlibAddFlag       = 0x00002000
	WIMLIB_ADD_FLAG_TEST_FILE_EXCLUSION                          WimlibAddFlag       = 0x00004000
	WIMLIB_ADD_FLAG_SNAPSHOT                                     WimlibAddFlag       = 0x00008000
	WIMLIB_ADD_FLAG_FILE_PATHS_UNNEEDED                          WimlibAddFlag       = 0x00010000
	WIMLIB_DELETE_FLAG_FORCE                                     WimlibDeleteFlag    = 0x00000001
	WIMLIB_DELETE_FLAG_RECURSIVE                                 WimlibDeleteFlag    = 0x00000002
	WIMLIB_EXPORT_FLAG_BOOT                                      WimlibExportFlag    = 0x00000001
	WIMLIB_EXPORT_FLAG_NO_NAMES                                  WimlibExportFlag    = 0x00000002
	WIMLIB_EXPORT_FLAG_NO_DESCRIPTIONS                           WimlibExportFlag    = 0x00000004
	WIMLIB_EXPORT_FLAG_GIFT                                      WimlibExportFlag    = 0x00000008
	WIMLIB_EXPORT_FLAG_WIMBOOT                                   WimlibExportFlag    = 0x00000010
	WIMLIB_EXTRACT_FLAG_NTFS                                     WimlibExtractFlag   = 0x00000001
	WIMLIB_EXTRACT_FLAG_RECOVER_DATA                             WimlibExtractFlag   = 0x00000002
	WIMLIB_EXTRACT_FLAG_UNIX_DATA                                WimlibExtractFlag   = 0x00000020
	WIMLIB_EXTRACT_FLAG_NO_ACLS                                  WimlibExtractFlag   = 0x00000040
	WIMLIB_EXTRACT_FLAG_STRICT_ACLS                              WimlibExtractFlag   = 0x00000080
	WIMLIB_EXTRACT_FLAG_RPFIX                                    WimlibExtractFlag   = 0x00000100
	WIMLIB_EXTRACT_FLAG_NORPFIX                                  WimlibExtractFlag   = 0x00000200
	WIMLIB_EXTRACT_FLAG_TO_STDOUT                                WimlibExtractFlag   = 0x00000400
	WIMLIB_EXTRACT_FLAG_REPLACE_INVALID_FILENAMES                WimlibExtractFlag   = 0x00000800
	WIMLIB_EXTRACT_FLAG_ALL_CASE_CONFLICTS                       WimlibExtractFlag   = 0x00001000
	WIMLIB_EXTRACT_FLAG_STRICT_TIMESTAMPS                        WimlibExtractFlag   = 0x00002000
	WIMLIB_EXTRACT_FLAG_STRICT_SHORT_NAMES                       WimlibExtractFlag   = 0x00004000
	WIMLIB_EXTRACT_FLAG_STRICT_SYMLINKS                          WimlibExtractFlag   = 0x00008000
	WIMLIB_EXTRACT_FLAG_GLOB_PATHS                               WimlibExtractFlag   = 0x00040000
	WIMLIB_EXTRACT_FLAG_STRICT_GLOB                              WimlibExtractFlag   = 0x00080000
	WIMLIB_EXTRACT_FLAG_NO_ATTRIBUTES                            WimlibExtractFlag   = 0x00100000
	WIMLIB_EXTRACT_FLAG_NO_PRESERVE_DIR_STRUCTURE                WimlibExtractFlag   = 0x00200000
	WIMLIB_EXTRACT_FLAG_WIMBOOT                                  WimlibExtractFlag   = 0x00400000
	WIMLIB_EXTRACT_FLAG_COMPACT_XPRESS4K                         WimlibExtractFlag   = 0x01000000
	WIMLIB_EXTRACT_FLAG_COMPACT_XPRESS8K                         WimlibExtractFlag   = 0x02000000
	WIMLIB_EXTRACT_FLAG_COMPACT_XPRESS16K                        WimlibExtractFlag   = 0x04000000
	WIMLIB_EXTRACT_FLAG_COMPACT_LZX                              WimlibExtractFlag   = 0x08000000
	WIMLIB_MOUNT_FLAG_READWRITE                                  WimlibMountFlag     = 0x00000001
	WIMLIB_MOUNT_FLAG_DEBUG                                      WimlibMountFlag     = 0x00000002
	WIMLIB_MOUNT_FLAG_STREAM_INTERFACE_NONE                      WimlibMountFlag     = 0x00000004
	WIMLIB_MOUNT_FLAG_STREAM_INTERFACE_XATTR                     WimlibMountFlag     = 0x00000008
	WIMLIB_MOUNT_FLAG_STREAM_INTERFACE_WINDOWS                   WimlibMountFlag     = 0x00000010
	WIMLIB_MOUNT_FLAG_UNIX_DATA                                  WimlibMountFlag     = 0x00000020
	WIMLIB_MOUNT_FLAG_ALLOW_OTHER                                WimlibMountFlag     = 0x00000040
	WIMLIB_OPEN_FLAG_CHECK_INTEGRITY                             WimlibOpenFlag      = 0x00000001
	WIMLIB_OPEN_FLAG_ERROR_IF_SPLIT                              WimlibOpenFlag      = 0x00000002
	WIMLIB_OPEN_FLAG_WRITE_ACCESS                                WimlibOpenFlag      = 0x00000004
	WIMLIB_UNMOUNT_FLAG_CHECK_INTEGRITY                          WimlibUnmountFlag   = 0x00000001
	WIMLIB_UNMOUNT_FLAG_COMMIT                                   WimlibUnmountFlag   = 0x00000002
	WIMLIB_UNMOUNT_FLAG_REBUILD                                  WimlibUnmountFlag   = 0x00000004
	WIMLIB_UNMOUNT_FLAG_RECOMPRESS                               WimlibUnmountFlag   = 0x00000008
	WIMLIB_UNMOUNT_FLAG_FORCE                                    WimlibUnmountFlag   = 0x00000010
	WIMLIB_UNMOUNT_FLAG_NEW_IMAGE                                WimlibUnmountFlag   = 0x00000020
	WIMLIB_UPDATE_FLAG_SEND_PROGRESS                             WimlibUpdateFlag    = 0x00000001
	WIMLIB_WRITE_FLAG_CHECK_INTEGRITY                            WimlibWriteFlag     = 0x00000001
	WIMLIB_WRITE_FLAG_NO_CHECK_INTEGRITY                         WimlibWriteFlag     = 0x00000002
	WIMLIB_WRITE_FLAG_PIPABLE                                    WimlibWriteFlag     = 0x00000004
	WIMLIB_WRITE_FLAG_NOT_PIPABLE                                WimlibWriteFlag     = 0x00000008
	WIMLIB_WRITE_FLAG_RECOMPRESS                                 WimlibWriteFlag     = 0x00000010
	WIMLIB_WRITE_FLAG_FSYNC                                      WimlibWriteFlag     = 0x00000020
	WIMLIB_WRITE_FLAG_REBUILD                                    WimlibWriteFlag     = 0x00000040
	WIMLIB_WRITE_FLAG_SOFT_DELETE                                WimlibWriteFlag     = 0x00000080
	WIMLIB_WRITE_FLAG_IGNORE_READONLY_FLAG                       WimlibWriteFlag     = 0x00000100
	WIMLIB_WRITE_FLAG_SKIP_EXTERNAL_WIMS                         WimlibWriteFlag     = 0x00000200
	WIMLIB_WRITE_FLAG_STREAMS_OK                                 WimlibWriteFlag     = 0x00000400
	WIMLIB_WRITE_FLAG_RETAIN_GUID                                WimlibWriteFlag     = 0x00000800
	WIMLIB_WRITE_FLAG_SOLID                                      WimlibWriteFlag     = 0x00001000
	WIMLIB_WRITE_FLAG_SEND_DONE_WITH_FILE_MESSAGES               WimlibWriteFlag     = 0x00002000
	WIMLIB_WRITE_FLAG_NO_SOLID_SORT                              WimlibWriteFlag     = 0x00004000
	WIMLIB_WRITE_FLAG_UNSAFE_COMPACT                             WimlibWriteFlag     = 0x00008000
	WIMLIB_INIT_FLAG_ASSUME_UTF8                                 WimlibInitFlag      = 0x00000001
	WIMLIB_INIT_FLAG_DONT_ACQUIRE_PRIVILEGES                     WimlibInitFlag      = 0x00000002
	WIMLIB_INIT_FLAG_STRICT_CAPTURE_PRIVILEGES                   WimlibInitFlag      = 0x00000004
	WIMLIB_INIT_FLAG_STRICT_APPLY_PRIVILEGES                     WimlibInitFlag      = 0x00000008
	WIMLIB_INIT_FLAG_DEFAULT_CASE_SENSITIVE                      WimlibInitFlag      = 0x00000010
	WIMLIB_INIT_FLAG_DEFAULT_CASE_INSENSITIVE                    WimlibInitFlag      = 0x00000020
	WIMLIB_REF_FLAG_GLOB_ENABLE                                  WimlibRefFlag       = 0x00000001
	WIMLIB_REF_FLAG_GLOB_ERR_ON_NOMATCH                          WimlibRefFlag       = 0x00000002
	WIMLIB_NO_IMAGE                                                                  = 0
	WIMLIB_ALL_IMAGES                                                                = -1
	WIMLIB_COMPRESSOR_FLAG_DESTRUCTIVE                                               = 0x80000000
)

type WimlibCompressionType int

const (
	WIMLIB_COMPRESSION_TYPE_NONE   WimlibCompressionType = 0
	WIMLIB_COMPRESSION_TYPE_XPRESS WimlibCompressionType = 1
	WIMLIB_COMPRESSION_TYPE_LZX    WimlibCompressionType = 2
	WIMLIB_COMPRESSION_TYPE_LZMS   WimlibCompressionType = 3
)

func (t WimlibCompressionType) String() string {
	switch t {
	case WIMLIB_COMPRESSION_TYPE_NONE:
		return "NONE"
	case WIMLIB_COMPRESSION_TYPE_XPRESS:
		return "XPRESS"
	case WIMLIB_COMPRESSION_TYPE_LZX:
		return "LZX"
	case WIMLIB_COMPRESSION_TYPE_LZMS:
		return "LZMS"
	default:
		return "INVALID"
	}
}
