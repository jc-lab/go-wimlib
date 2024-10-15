package binding

/*
#include <wimlib.h>
*/
import "C"

import (
	"golang.org/x/sys/windows"
	"unsafe"
)

func malloc(size int) unsafe.Pointer {
	return C.malloc(C.ulonglong(size))
}

func convertToTchar(s string) (*C.wimlib_tchar, func()) {
	buf := windows.StringToUTF16(s)
	ptr := unsafe.Pointer(&buf[0])
	retainer.Keep(ptr, buf)
	return (*C.wimlib_tchar)(ptr), func() {
		retainer.Remove(ptr)
	}
}

func convertFromTchar(tstr *C.wimlib_tchar) string {
	return windows.UTF16PtrToString((*uint16)(tstr))
}
