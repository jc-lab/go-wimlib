package binding

/*
#include <wimlib.h>
#include <stdlib.h>
*/
import "C"

// WimlibError represents an error returned from wimlib C functions.
type WimlibError struct {
	Code int
}

func (e *WimlibError) Error() string {
	return convertFromTchar(C.wimlib_get_error_string(C.enum_wimlib_error_code(e.Code)))
}
