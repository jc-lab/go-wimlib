//go:build !windows
// +build !windows

package binding

/*
#include <wimlib.h>
#include <stdlib.h>
*/
import "C"
import (
	"unsafe"
)

func malloc(size int) unsafe.Pointer {
	return C.malloc(C.ulong(size))
}

// convertToTchar converts a Go string to a wimlib_tchar
func convertToTchar(s string) (*C.wimlib_tchar, func()) {
	cstr := C.CString(s)
	ptr := unsafe.Pointer(cstr)
	free := func() {
		C.free(ptr)
	}
	return (*C.wimlib_tchar)(ptr), free
}

func convertFromTchar(tstr *C.wimlib_tchar) string {
	return C.GoString((*C.char)(unsafe.Pointer(tstr)))
}
