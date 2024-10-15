#include <wimlib.h>

extern int go_wimlib_progress_go(void *progctx, int msg_type, void *mpack_ptr,
                                 int mpack_length);

extern enum wimlib_progress_status
go_wimlib_progress_c(enum wimlib_progress_msg msg_type,
                     union wimlib_progress_info *info, void *progctx);
