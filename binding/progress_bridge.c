#include <msgpack.h>
#include <wimlib.h>

#if defined _WIN32 || defined __CYGWIN__
#include <windows.h>

static void my_msgpack_pack_str_body_key(msgpack_packer *pk,
                                           const char *val);

static void my_msgpack_pack_str_body_value(msgpack_packer *pk,
                                           const wchar_t *val) {
  int size_needed =
      WideCharToMultiByte(CP_UTF8, 0, val, -1, NULL, 0, NULL, NULL);
  if (size_needed == 0) {
  	msgpack_pack_nil(pk);
  	return;
  }

  char* u8buf = (char*)malloc(size_needed);
  size_needed = WideCharToMultiByte(CP_UTF8, 0, val, -1, u8buf, size_needed, NULL, NULL);

  msgpack_pack_str(pk, size_needed);
  msgpack_pack_str_body(pk, u8buf, size_needed);

  free(u8buf);
}
#else
#define my_msgpack_pack_str_body_value my_msgpack_pack_str_body_key
#endif

extern int go_wimlib_progress_go(void *progctx, int msg_type, void *mpack_ptr,
                                 int mpack_length);

static void my_msgpack_pack_str_body_key(msgpack_packer *pk,
                                           const char *val) {
  size_t len = strlen(val);
  msgpack_pack_str(pk, len);
  msgpack_pack_str_body(pk, val, len);
}

enum wimlib_progress_status
go_wimlib_progress_c(enum wimlib_progress_msg msg_type,
                     union wimlib_progress_info *info, void *progctx) {
  enum wimlib_progress_status result;
  msgpack_sbuffer sbuf;
  msgpack_packer pk;

  msgpack_sbuffer_init(&sbuf);
  msgpack_packer_init(&pk, &sbuf, msgpack_sbuffer_write);

  msgpack_pack_map(&pk, 2);
  my_msgpack_pack_str_body_key(&pk, "msg_type");
  msgpack_pack_int(&pk, msg_type);
  my_msgpack_pack_str_body_key(&pk, "info");

  switch (msg_type) {
  case WIMLIB_PROGRESS_MSG_WRITE_STREAMS:
    msgpack_pack_map(&pk, 10);
    my_msgpack_pack_str_body_key(&pk, "total_bytes");
    msgpack_pack_uint64(&pk, info->write_streams.total_bytes);
    my_msgpack_pack_str_body_key(&pk, "total_streams");
    msgpack_pack_uint64(&pk, info->write_streams.total_streams);
    my_msgpack_pack_str_body_key(&pk, "completed_bytes");
    msgpack_pack_uint64(&pk, info->write_streams.completed_bytes);
    my_msgpack_pack_str_body_key(&pk, "completed_streams");
    msgpack_pack_uint64(&pk, info->write_streams.completed_streams);
    my_msgpack_pack_str_body_key(&pk, "num_threads");
    msgpack_pack_uint32(&pk, info->write_streams.num_threads);
    my_msgpack_pack_str_body_key(&pk, "compression_type");
    msgpack_pack_int32(&pk, info->write_streams.compression_type);
    my_msgpack_pack_str_body_key(&pk, "total_parts");
    msgpack_pack_uint32(&pk, info->write_streams.total_parts);
    my_msgpack_pack_str_body_key(&pk, "completed_parts");
    msgpack_pack_uint32(&pk, info->write_streams.completed_parts);
    my_msgpack_pack_str_body_key(&pk, "completed_compressed_bytes");
    msgpack_pack_uint64(&pk, info->write_streams.completed_compressed_bytes);
    break;

  case WIMLIB_PROGRESS_MSG_SCAN_BEGIN:
  case WIMLIB_PROGRESS_MSG_SCAN_DENTRY:
  case WIMLIB_PROGRESS_MSG_SCAN_END:
    msgpack_pack_map(&pk, 7);
    my_msgpack_pack_str_body_key(&pk, "source");
    my_msgpack_pack_str_body_value(&pk, info->scan.source);
    my_msgpack_pack_str_body_key(&pk, "cur_path");
    my_msgpack_pack_str_body_value(&pk, info->scan.cur_path);
    my_msgpack_pack_str_body_key(&pk, "status");
    msgpack_pack_int(&pk, info->scan.status);
    my_msgpack_pack_str_body_key(&pk, "wim_target_path");
    my_msgpack_pack_str_body_value(&pk, info->scan.wim_target_path);
    my_msgpack_pack_str_body_key(&pk, "num_dirs_scanned");
    msgpack_pack_uint64(&pk, info->scan.num_dirs_scanned);
    my_msgpack_pack_str_body_key(&pk, "num_nondirs_scanned");
    msgpack_pack_uint64(&pk, info->scan.num_nondirs_scanned);
    my_msgpack_pack_str_body_key(&pk, "num_bytes_scanned");
    msgpack_pack_uint64(&pk, info->scan.num_bytes_scanned);
    break;

  case WIMLIB_PROGRESS_MSG_EXTRACT_SPWM_PART_BEGIN:
  case WIMLIB_PROGRESS_MSG_EXTRACT_IMAGE_BEGIN:
  case WIMLIB_PROGRESS_MSG_EXTRACT_TREE_BEGIN:
  case WIMLIB_PROGRESS_MSG_EXTRACT_FILE_STRUCTURE:
  case WIMLIB_PROGRESS_MSG_EXTRACT_STREAMS:
  case WIMLIB_PROGRESS_MSG_EXTRACT_METADATA:
  case WIMLIB_PROGRESS_MSG_EXTRACT_TREE_END:
  case WIMLIB_PROGRESS_MSG_EXTRACT_IMAGE_END:
    msgpack_pack_map(&pk, 13);
    my_msgpack_pack_str_body_key(&pk, "image");
    msgpack_pack_uint32(&pk, info->extract.image);
    my_msgpack_pack_str_body_key(&pk, "extract_flags");
    msgpack_pack_uint32(&pk, info->extract.extract_flags);
    my_msgpack_pack_str_body_key(&pk, "wimfile_name");
    my_msgpack_pack_str_body_value(&pk, info->extract.wimfile_name);
    my_msgpack_pack_str_body_key(&pk, "image_name");
    my_msgpack_pack_str_body_value(&pk, info->extract.image_name);
    my_msgpack_pack_str_body_key(&pk, "target");
    my_msgpack_pack_str_body_value(&pk, info->extract.target);
    my_msgpack_pack_str_body_key(&pk, "total_bytes");
    msgpack_pack_uint64(&pk, info->extract.total_bytes);
    my_msgpack_pack_str_body_key(&pk, "completed_bytes");
    msgpack_pack_uint64(&pk, info->extract.completed_bytes);
    my_msgpack_pack_str_body_key(&pk, "total_streams");
    msgpack_pack_uint64(&pk, info->extract.total_streams);
    my_msgpack_pack_str_body_key(&pk, "completed_streams");
    msgpack_pack_uint64(&pk, info->extract.completed_streams);
    my_msgpack_pack_str_body_key(&pk, "part_number");
    msgpack_pack_uint32(&pk, info->extract.part_number);
    my_msgpack_pack_str_body_key(&pk, "total_parts");
    msgpack_pack_uint32(&pk, info->extract.total_parts);
    my_msgpack_pack_str_body_key(&pk, "guid");
    msgpack_pack_bin(&pk, WIMLIB_GUID_LEN);
    msgpack_pack_bin_body(&pk, info->extract.guid, WIMLIB_GUID_LEN);
    my_msgpack_pack_str_body_key(&pk, "current_file_count");
    msgpack_pack_uint64(&pk, info->extract.current_file_count);
    break;

  case WIMLIB_PROGRESS_MSG_RENAME:
    msgpack_pack_map(&pk, 2);
    my_msgpack_pack_str_body_key(&pk, "from");
    my_msgpack_pack_str_body_value(&pk, info->rename.from);
    my_msgpack_pack_str_body_key(&pk, "to");
    my_msgpack_pack_str_body_value(&pk, info->rename.to);
    break;

  case WIMLIB_PROGRESS_MSG_UPDATE_BEGIN_COMMAND:
  case WIMLIB_PROGRESS_MSG_UPDATE_END_COMMAND:
    msgpack_pack_map(&pk, 3);
    my_msgpack_pack_str_body_key(&pk, "command");
    msgpack_pack_uint64(&pk, (uint64_t)info->update.command);
    my_msgpack_pack_str_body_key(&pk, "completed_commands");
    msgpack_pack_uint64(&pk, info->update.completed_commands);
    my_msgpack_pack_str_body_key(&pk, "total_commands");
    msgpack_pack_uint64(&pk, info->update.total_commands);
    break;

  case WIMLIB_PROGRESS_MSG_VERIFY_INTEGRITY:
  case WIMLIB_PROGRESS_MSG_CALC_INTEGRITY:
    msgpack_pack_map(&pk, 6);
    my_msgpack_pack_str_body_key(&pk, "total_bytes");
    msgpack_pack_uint64(&pk, info->integrity.total_bytes);
    my_msgpack_pack_str_body_key(&pk, "completed_bytes");
    msgpack_pack_uint64(&pk, info->integrity.completed_bytes);
    my_msgpack_pack_str_body_key(&pk, "total_chunks");
    msgpack_pack_uint32(&pk, info->integrity.total_chunks);
    my_msgpack_pack_str_body_key(&pk, "completed_chunks");
    msgpack_pack_uint32(&pk, info->integrity.completed_chunks);
    my_msgpack_pack_str_body_key(&pk, "chunk_size");
    msgpack_pack_uint32(&pk, info->integrity.chunk_size);
    my_msgpack_pack_str_body_key(&pk, "filename");
    my_msgpack_pack_str_body_value(&pk, info->integrity.filename);
    break;

  case WIMLIB_PROGRESS_MSG_SPLIT_BEGIN_PART:
  case WIMLIB_PROGRESS_MSG_SPLIT_END_PART:
    msgpack_pack_map(&pk, 6);
    my_msgpack_pack_str_body_key(&pk, "total_bytes");
    msgpack_pack_uint64(&pk, info->split.total_bytes);
    my_msgpack_pack_str_body_key(&pk, "completed_bytes");
    msgpack_pack_uint64(&pk, info->split.completed_bytes);
    my_msgpack_pack_str_body_key(&pk, "cur_part_number");
    msgpack_pack_unsigned_int(&pk, info->split.cur_part_number);
    my_msgpack_pack_str_body_key(&pk, "total_parts");
    msgpack_pack_unsigned_int(&pk, info->split.total_parts);
    my_msgpack_pack_str_body_key(&pk, "part_name");
    my_msgpack_pack_str_body_value(&pk, info->split.part_name);
    break;

  case WIMLIB_PROGRESS_MSG_REPLACE_FILE_IN_WIM:
    msgpack_pack_map(&pk, 1);
    my_msgpack_pack_str_body_key(&pk, "path_in_wim");
    my_msgpack_pack_str_body_value(&pk, info->replace.path_in_wim);
    break;

  case WIMLIB_PROGRESS_MSG_WIMBOOT_EXCLUDE:
    msgpack_pack_map(&pk, 2);
    my_msgpack_pack_str_body_key(&pk, "path_in_wim");
    my_msgpack_pack_str_body_value(&pk, info->wimboot_exclude.path_in_wim);
    my_msgpack_pack_str_body_key(&pk, "extraction_path");
    my_msgpack_pack_str_body_value(&pk, info->wimboot_exclude.extraction_path);
    break;
  case WIMLIB_PROGRESS_MSG_UNMOUNT_BEGIN:
    msgpack_pack_map(&pk, 5);
    my_msgpack_pack_str_body_key(&pk, "mountpoint");
    my_msgpack_pack_str_body_value(&pk, info->unmount.mountpoint);
    my_msgpack_pack_str_body_key(&pk, "mounted_wim");
    my_msgpack_pack_str_body_value(&pk, info->unmount.mounted_wim);
    my_msgpack_pack_str_body_key(&pk, "mounted_image");
    msgpack_pack_uint32(&pk, info->unmount.mounted_image);
    my_msgpack_pack_str_body_key(&pk, "mount_flags");
    msgpack_pack_uint32(&pk, info->unmount.mount_flags);
    my_msgpack_pack_str_body_key(&pk, "unmount_flags");
    msgpack_pack_uint32(&pk, info->unmount.unmount_flags);
    break;

  case WIMLIB_PROGRESS_MSG_DONE_WITH_FILE:
    msgpack_pack_map(&pk, 1);
    my_msgpack_pack_str_body_key(&pk, "path_to_file");
    my_msgpack_pack_str_body_value(&pk, info->done_with_file.path_to_file);
    break;

  case WIMLIB_PROGRESS_MSG_BEGIN_VERIFY_IMAGE:
  case WIMLIB_PROGRESS_MSG_END_VERIFY_IMAGE:
    msgpack_pack_map(&pk, 3);
    my_msgpack_pack_str_body_key(&pk, "wimfile");
    my_msgpack_pack_str_body_value(&pk, info->verify_image.wimfile);
    my_msgpack_pack_str_body_key(&pk, "total_images");
    msgpack_pack_uint32(&pk, info->verify_image.total_images);
    my_msgpack_pack_str_body_key(&pk, "current_image");
    msgpack_pack_uint32(&pk, info->verify_image.current_image);
    break;

  case WIMLIB_PROGRESS_MSG_VERIFY_STREAMS:
    msgpack_pack_map(&pk, 5);
    my_msgpack_pack_str_body_key(&pk, "wimfile");
    my_msgpack_pack_str_body_value(&pk, info->verify_streams.wimfile);
    my_msgpack_pack_str_body_key(&pk, "total_streams");
    msgpack_pack_uint64(&pk, info->verify_streams.total_streams);
    my_msgpack_pack_str_body_key(&pk, "total_bytes");
    msgpack_pack_uint64(&pk, info->verify_streams.total_bytes);
    my_msgpack_pack_str_body_key(&pk, "completed_streams");
    msgpack_pack_uint64(&pk, info->verify_streams.completed_streams);
    my_msgpack_pack_str_body_key(&pk, "completed_bytes");
    msgpack_pack_uint64(&pk, info->verify_streams.completed_bytes);
    break;

  case WIMLIB_PROGRESS_MSG_TEST_FILE_EXCLUSION:
    msgpack_pack_map(&pk, 2);
    my_msgpack_pack_str_body_key(&pk, "path");
    my_msgpack_pack_str_body_value(&pk, info->test_file_exclusion.path);
    my_msgpack_pack_str_body_key(&pk, "will_exclude");
    if (info->test_file_exclusion.will_exclude) {
      msgpack_pack_true(&pk);
    } else {
      msgpack_pack_false(&pk);
    }
    break;

  case WIMLIB_PROGRESS_MSG_HANDLE_ERROR:
    msgpack_pack_map(&pk, 3);
    my_msgpack_pack_str_body_key(&pk, "path");
    if (info->handle_error.path) {
      my_msgpack_pack_str_body_value(&pk, info->handle_error.path);
    } else {
      msgpack_pack_nil(&pk);
    }
    my_msgpack_pack_str_body_key(&pk, "error_code");
    msgpack_pack_int(&pk, info->handle_error.error_code);
    my_msgpack_pack_str_body_key(&pk, "will_ignore");
    if (info->handle_error.will_ignore) {
      msgpack_pack_true(&pk);
    } else {
      msgpack_pack_false(&pk);
    }
    break;

  default:
    msgpack_pack_nil(&pk);
    break;
  }

  result = (enum wimlib_progress_status)go_wimlib_progress_go(
      progctx, msg_type, sbuf.data, (int)(sbuf.size));

  msgpack_sbuffer_destroy(&sbuf);

  return result;
}
