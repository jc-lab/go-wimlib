package binding

/*
#include <stdlib.h>
#include "progress_bridge.h"
*/
import "C"
import (
	"github.com/jc-lab/go-wimlib/model"
	"unsafe"
)

//export go_wimlib_progress_go
func go_wimlib_progress_go(progctx unsafe.Pointer, cMsgType C.int, mpackPtr unsafe.Pointer, mpackLength C.int) C.int {
	holder := (*wimHolder)(progctx)

	msgType := model.ProgressMsgType(cMsgType)
	mpackData := unsafe.Slice((*byte)(mpackPtr), mpackLength)
	var progressMsg model.ProgressMsg
	if _, err := progressMsg.UnmarshalMsg(mpackData); err != nil {
		holder.ProgressFunc(msgType, nil, err)
		return C.int(model.WIMLIB_PROGRESS_STATUS_CONTINUE)
	}

	return C.int(holder.ProgressFunc(msgType, &progressMsg, nil))
}

// OpenWimWithProgress opens a WIM file with progress
func OpenWimWithProgress(wimFile string, openFlags int, progressFunc ProgressFunc) (*Wim, error) {
	cWimFile, freeWimFile := convertToTchar(wimFile)
	defer freeWimFile()

	instance := &Wim{
		ProgressFunc: progressFunc,
	}
	instance.pinner.Pin(instance)

	holder := (*wimHolder)(malloc(wimHolderSize))
	holder.Wim = instance
	instance.holder = unsafe.Pointer(holder)

	ret := C.wimlib_open_wim_with_progress(cWimFile, C.int(openFlags), &instance.WIMStruct, (*[0]byte)(unsafe.Pointer(C.go_wimlib_progress_c)), unsafe.Pointer(holder))
	if ret != 0 {
		instance.Close()
		return nil, &WimlibError{Code: int(ret)}
	}

	return instance, nil
}
