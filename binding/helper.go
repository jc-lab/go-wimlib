package binding

import "C"
import (
	"github.com/jc-lab/go-wimlib/model"
	"runtime"
	"unsafe"
)

/*
#include <wimlib.h>
#include <stdlib.h>
*/
import "C"

type ProgressFunc func(msgType model.ProgressMsgType, msg *model.ProgressMsg, err error) model.ProgressStatus

type ProgressFuncHolder struct {
	Func ProgressFunc
}

type Wim struct {
	WIMStruct    *C.WIMStruct
	ProgressFunc ProgressFunc
	pinner       runtime.Pinner

	holder unsafe.Pointer
}

func (w *Wim) Close() error {
	w.pinner.Unpin()
	if w.WIMStruct != nil {
		FreeWim(w.WIMStruct)
	}
	if w.holder != nil {
		C.free(w.holder)
		w.holder = nil
	}
	return nil
}

type wimHolder struct {
	*Wim
}

var wimHolderSize = int(unsafe.Sizeof(wimHolder{}))

func GetWimInfoEx(wim *C.WIMStruct) (*model.WimInfo, error) {
	var cInfo C.struct_wimlib_wim_info
	err := GetWimInfo(wim, &cInfo)
	if err != nil {
		return nil, err
	}
	return NewWimInfoFromC(&cInfo), nil
}

func NewWimInfoFromC(cInfo *C.struct_wimlib_wim_info) *model.WimInfo {
	goInfo := &model.WimInfo{}

	// Copy GUID
	for i := 0; i < model.WIMLIB_GUID_LEN; i++ {
		goInfo.Guid[i] = byte(cInfo.guid[i])
	}

	goInfo.ImageCount = uint32(cInfo.image_count)
	goInfo.BootIndex = uint32(cInfo.boot_index)
	goInfo.WimVersion = uint32(cInfo.wim_version)
	goInfo.ChunkSize = uint32(cInfo.chunk_size)
	goInfo.PartNumber = uint16(cInfo.part_number)
	goInfo.TotalParts = uint16(cInfo.total_parts)
	goInfo.CompressionType = model.WimlibCompressionType(cInfo.compression_type)
	goInfo.TotalBytes = uint64(cInfo.total_bytes)

	// Extract boolean fields from flags
	flags := *(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(&cInfo.total_bytes)) + 8))
	goInfo.HasIntegrityTable = (flags & 0x1) != 0
	goInfo.OpenedFromFile = (flags & 0x2) != 0
	goInfo.IsReadonly = (flags & 0x4) != 0
	goInfo.HasRpfix = (flags & 0x8) != 0
	goInfo.IsMarkedReadonly = (flags & 0x10) != 0
	goInfo.Spanned = (flags & 0x20) != 0
	goInfo.WriteInProgress = (flags & 0x40) != 0
	goInfo.MetadataOnly = (flags & 0x80) != 0
	goInfo.ResourceOnly = (flags & 0x100) != 0
	goInfo.Pipable = (flags & 0x200) != 0

	return goInfo
}
