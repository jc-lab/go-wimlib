package binding

/*
#include <wimlib.h>
#include <stdlib.h>
*/
import "C"
import "unsafe"

// AddEmptyImage adds an empty image to a WIM file
func AddEmptyImage(wim *C.WIMStruct, name string) (int, error) {
	cName, freeName := convertToTchar(name)
	defer freeName()

	var newIdx C.int
	ret := C.wimlib_add_empty_image(wim, cName, &newIdx)
	if ret != 0 {
		return 0, &WimlibError{Code: int(ret)}
	}
	return int(newIdx), nil
}

// AddImage adds an image to a WIM file
func AddImage(wim *C.WIMStruct, source, name, configFile string, addFlags int) error {
	cSource, freeSource := convertToTchar(source)
	defer freeSource()
	cName, freeName := convertToTchar(name)
	defer freeName()
	cConfigFile, freeConfigFile := convertToTchar(configFile)
	defer freeConfigFile()

	ret := C.wimlib_add_image(wim, cSource, cName, cConfigFile, C.int(addFlags))
	if ret != 0 {
		return &WimlibError{Code: int(ret)}
	}
	return nil
}

// AddImageMultisource adds an image to a WIM file from multiple sources
func AddImageMultisource(wim *C.WIMStruct, sources []C.struct_wimlib_capture_source, name, configFile string, addFlags int) error {
	cName, freeName := convertToTchar(name)
	defer freeName()
	cConfigFile, freeConfigFile := convertToTchar(configFile)
	defer freeConfigFile()

	ret := C.wimlib_add_image_multisource(wim, &sources[0], C.size_t(len(sources)), cName, cConfigFile, C.int(addFlags))
	if ret != 0 {
		return &WimlibError{Code: int(ret)}
	}
	return nil
}

// AddTree adds a directory tree to a WIM image
func AddTree(wim *C.WIMStruct, image int, fsSourcePath, wimTargetPath string, addFlags int) error {
	cFsSourcePath, freeFsSourcePath := convertToTchar(fsSourcePath)
	defer freeFsSourcePath()
	cWimTargetPath, freeWimTargetPath := convertToTchar(wimTargetPath)
	defer freeWimTargetPath()

	ret := C.wimlib_add_tree(wim, C.int(image), cFsSourcePath, cWimTargetPath, C.int(addFlags))
	if ret != 0 {
		return &WimlibError{Code: int(ret)}
	}
	return nil
}

// CreateNewWim creates a new WIM file
func CreateNewWim(ctype int) (*C.WIMStruct, error) {
	var wim *C.WIMStruct
	ret := C.wimlib_create_new_wim(C.enum_wimlib_compression_type(ctype), &wim)
	if ret != 0 {
		return nil, &WimlibError{Code: int(ret)}
	}
	return wim, nil
}

// DeleteImage deletes an image from a WIM file
func DeleteImage(wim *C.WIMStruct, image int) error {
	ret := C.wimlib_delete_image(wim, C.int(image))
	if ret != 0 {
		return &WimlibError{Code: int(ret)}
	}
	return nil
}

// DeletePath deletes a path from a WIM image
func DeletePath(wim *C.WIMStruct, image int, path string, deleteFlags int) error {
	cPath, freePath := convertToTchar(path)
	defer freePath()

	ret := C.wimlib_delete_path(wim, C.int(image), cPath, C.int(deleteFlags))
	if ret != 0 {
		return &WimlibError{Code: int(ret)}
	}
	return nil
}

// ExportImage exports an image from one WIM file to another
func ExportImage(srcWim *C.WIMStruct, srcImage int, destWim *C.WIMStruct, destName, destDescription string, exportFlags int) error {
	cDestName, freeDestName := convertToTchar(destName)
	defer freeDestName()
	cDestDescription, freeDestDescription := convertToTchar(destDescription)
	defer freeDestDescription()

	ret := C.wimlib_export_image(srcWim, C.int(srcImage), destWim, cDestName, cDestDescription, C.int(exportFlags))
	if ret != 0 {
		return &WimlibError{Code: int(ret)}
	}
	return nil
}

// ExtractImage extracts an image from a WIM file
func ExtractImage(wim *C.WIMStruct, image int, target string, extractFlags int) error {
	cTarget, freeTarget := convertToTchar(target)
	defer freeTarget()

	ret := C.wimlib_extract_image(wim, C.int(image), cTarget, C.int(extractFlags))
	if ret != 0 {
		return &WimlibError{Code: int(ret)}
	}
	return nil
}

// ExtractImageFromPipe extracts an image from a pipeable WIM
func ExtractImageFromPipe(pipeFd int, imageNumOrName string, target string, extractFlags int) error {
	cImageNumOrName, freeImageNumOrName := convertToTchar(imageNumOrName)
	defer freeImageNumOrName()
	cTarget, freeTarget := convertToTchar(target)
	defer freeTarget()

	ret := C.wimlib_extract_image_from_pipe(C.int(pipeFd), cImageNumOrName, cTarget, C.int(extractFlags))
	if ret != 0 {
		return &WimlibError{Code: int(ret)}
	}
	return nil
}

// ExtractPathlist extracts a list of paths from a WIM image
func ExtractPathlist(wim *C.WIMStruct, image int, target, pathListFile string, extractFlags int) error {
	cTarget, freeTarget := convertToTchar(target)
	defer freeTarget()
	cPathListFile, freePathListFile := convertToTchar(pathListFile)
	defer freePathListFile()

	ret := C.wimlib_extract_pathlist(wim, C.int(image), cTarget, cPathListFile, C.int(extractFlags))
	if ret != 0 {
		return &WimlibError{Code: int(ret)}
	}
	return nil
}

// ExtractPaths extracts a list of paths from a WIM image
func ExtractPaths(wim *C.WIMStruct, image int, target string, paths []string, extractFlags int) error {
	cTarget, freeTarget := convertToTchar(target)
	defer freeTarget()

	cPaths := make([]*C.wimlib_tchar, len(paths))
	freeFuncs := make([]func(), len(paths))
	for i, path := range paths {
		cPaths[i], freeFuncs[i] = convertToTchar(path)
	}
	defer func() {
		for _, free := range freeFuncs {
			free()
		}
	}()

	ret := C.wimlib_extract_paths(wim, C.int(image), cTarget, (**C.wimlib_tchar)(unsafe.Pointer(&cPaths[0])), C.size_t(len(paths)), C.int(extractFlags))
	if ret != 0 {
		return &WimlibError{Code: int(ret)}
	}
	return nil
}

// ExtractXMLData extracts the XML data from a WIM file
func ExtractXMLData(wim *C.WIMStruct, fp *C.FILE) error {
	ret := C.wimlib_extract_xml_data(wim, fp)
	if ret != 0 {
		return &WimlibError{Code: int(ret)}
	}
	return nil
}

// FreeWim frees a WIM structure
func FreeWim(wim *C.WIMStruct) {
	C.wimlib_free(wim)
}

// GetXMLData gets the XML data from a WIM file
func GetXMLData(wim *C.WIMStruct) ([]byte, error) {
	var buf unsafe.Pointer
	var bufsize C.size_t
	ret := C.wimlib_get_xml_data(wim, &buf, &bufsize)
	if ret != 0 {
		return nil, &WimlibError{Code: int(ret)}
	}
	defer C.free(buf)
	return C.GoBytes(buf, C.int(bufsize)), nil
}

// GlobalInit initializes the wimlib library
func GlobalInit(initFlags int) error {
	ret := C.wimlib_global_init(C.int(initFlags))
	if ret != 0 {
		return &WimlibError{Code: int(ret)}
	}
	return nil
}

// GlobalCleanup cleans up the wimlib library
func GlobalCleanup() {
	C.wimlib_global_cleanup()
}

// ImageNameInUse checks if an image name is already in use in a WIM file
func ImageNameInUse(wim *C.WIMStruct, name string) bool {
	cName, freeName := convertToTchar(name)
	defer freeName()

	return bool(C.wimlib_image_name_in_use(wim, cName))
}

// IterateDirTree iterates over the directory tree of a WIM image
func IterateDirTree(wim *C.WIMStruct, image int, path string, flags int, cb C.wimlib_iterate_dir_tree_callback_t, userCtx unsafe.Pointer) error {
	cPath, freePath := convertToTchar(path)
	defer freePath()

	ret := C.wimlib_iterate_dir_tree(wim, C.int(image), cPath, C.int(flags), cb, userCtx)
	if ret != 0 {
		return &WimlibError{Code: int(ret)}
	}
	return nil
}

// IterateLookupTable iterates over the lookup table of a WIM file
func IterateLookupTable(wim *C.WIMStruct, flags int, cb C.wimlib_iterate_lookup_table_callback_t, userCtx unsafe.Pointer) error {
	ret := C.wimlib_iterate_lookup_table(wim, C.int(flags), cb, userCtx)
	if ret != 0 {
		return &WimlibError{Code: int(ret)}
	}
	return nil
}

// Join joins split WIM files into a single WIM file
func Join(swms []string, outputPath string, swmOpenFlags, wimWriteFlags int) error {
	cSwms := make([]*C.wimlib_tchar, len(swms))
	freeFuncs := make([]func(), len(swms))
	for i, swm := range swms {
		cSwms[i], freeFuncs[i] = convertToTchar(swm)
	}
	defer func() {
		for _, free := range freeFuncs {
			free()
		}
	}()

	cOutputPath, freeOutputPath := convertToTchar(outputPath)
	defer freeOutputPath()

	ret := C.wimlib_join((**C.wimlib_tchar)(unsafe.Pointer(&cSwms[0])), C.unsigned(len(swms)), cOutputPath, C.int(swmOpenFlags), C.int(wimWriteFlags))
	if ret != 0 {
		return &WimlibError{Code: int(ret)}
	}
	return nil
}

// LoadTextFile loads a text file into memory
func LoadTextFile(path string) (string, error) {
	cPath, freePath := convertToTchar(path)
	defer freePath()

	var tstr *C.wimlib_tchar
	var tstrNchars C.size_t
	ret := C.wimlib_load_text_file(cPath, &tstr, &tstrNchars)
	if ret != 0 {
		return "", &WimlibError{Code: int(ret)}
	}
	defer C.free(unsafe.Pointer(tstr))

	return convertFromTchar(tstr), nil
}

// MountImage mounts a WIM image
func MountImage(wim *C.WIMStruct, image int, dir string, mountFlags int, stagingDir string) error {
	cDir, freeDir := convertToTchar(dir)
	defer freeDir()
	cStagingDir, freeStagingDir := convertToTchar(stagingDir)
	defer freeStagingDir()

	ret := C.wimlib_mount_image(wim, C.int(image), cDir, C.int(mountFlags), cStagingDir)
	if ret != 0 {
		return &WimlibError{Code: int(ret)}
	}
	return nil
}

// OpenWim opens a WIM file
func OpenWim(wimFile string, openFlags int) (*C.WIMStruct, error) {
	cWimFile, freeWimFile := convertToTchar(wimFile)
	defer freeWimFile()

	var wim *C.WIMStruct
	ret := C.wimlib_open_wim(cWimFile, C.int(openFlags), &wim)
	if ret != 0 {
		return nil, &WimlibError{Code: int(ret)}
	}
	return wim, nil
}

// Overwrite overwrites a WIM file
func Overwrite(wim *C.WIMStruct, writeFlags int, numThreads uint) error {
	ret := C.wimlib_overwrite(wim, C.int(writeFlags), C.unsigned(numThreads))
	if ret != 0 {
		return &WimlibError{Code: int(ret)}
	}
	return nil
}

// PrintAvailableImages prints information about available images in a WIM file
func PrintAvailableImages(wim *C.WIMStruct, image int) {
	C.wimlib_print_available_images(wim, C.int(image))
}

// PrintHeader prints the header of a WIM file
func PrintHeader(wim *C.WIMStruct) {
	C.wimlib_print_header(wim)
}

// ReferenceResourceFiles references resource files
func ReferenceResourceFiles(wim *C.WIMStruct, resourceWimfilesOrGlobs []string, refFlags, openFlags int) error {
	cResourceWimfilesOrGlobs := make([]*C.wimlib_tchar, len(resourceWimfilesOrGlobs))
	freeFuncs := make([]func(), len(resourceWimfilesOrGlobs))
	for i, path := range resourceWimfilesOrGlobs {
		cResourceWimfilesOrGlobs[i], freeFuncs[i] = convertToTchar(path)
	}
	defer func() {
		for _, free := range freeFuncs {
			free()
		}
	}()

	ret := C.wimlib_reference_resource_files(wim, (**C.wimlib_tchar)(unsafe.Pointer(&cResourceWimfilesOrGlobs[0])), C.unsigned(len(resourceWimfilesOrGlobs)), C.int(refFlags), C.int(openFlags))
	if ret != 0 {
		return &WimlibError{Code: int(ret)}
	}
	return nil
}

// ReferenceResources references resources from other WIM files
func ReferenceResources(wim *C.WIMStruct, resourceWims []*C.WIMStruct, refFlags int) error {
	ret := C.wimlib_reference_resources(wim, (**C.WIMStruct)(unsafe.Pointer(&resourceWims[0])), C.unsigned(len(resourceWims)), C.int(refFlags))
	if ret != 0 {
		return &WimlibError{Code: int(ret)}
	}
	return nil
}

// ReferenceTemplateImage references a template image
func ReferenceTemplateImage(wim *C.WIMStruct, newImage int, templateWim *C.WIMStruct, templateImage int, flags int) error {
	ret := C.wimlib_reference_template_image(wim, C.int(newImage), templateWim, C.int(templateImage), C.int(flags))
	if ret != 0 {
		return &WimlibError{Code: int(ret)}
	}
	return nil
}

// RenamePath renames a path in a WIM image
func RenamePath(wim *C.WIMStruct, image int, sourcePath, destPath string) error {
	cSourcePath, freeSourcePath := convertToTchar(sourcePath)
	defer freeSourcePath()
	cDestPath, freeDestPath := convertToTchar(destPath)
	defer freeDestPath()

	ret := C.wimlib_rename_path(wim, C.int(image), cSourcePath, cDestPath)
	if ret != 0 {
		return &WimlibError{Code: int(ret)}
	}
	return nil
}

// ResolveImage resolves an image name or number to an image index
func ResolveImage(wim *C.WIMStruct, imageNameOrNum string) (int, error) {
	cImageNameOrNum, freeImageNameOrNum := convertToTchar(imageNameOrNum)
	defer freeImageNameOrNum()

	ret := C.wimlib_resolve_image(wim, cImageNameOrNum)
	if ret < 0 {
		return 0, &WimlibError{Code: int(ret)}
	}
	return int(ret), nil
}

// GetWimInfo sets information about a WIM file
func GetWimInfo(wim *C.WIMStruct, info *C.struct_wimlib_wim_info) error {
	ret := C.wimlib_get_wim_info(wim, info)
	if ret != 0 {
		return &WimlibError{Code: int(ret)}
	}
	return nil
}

// SetWimInfo sets information about a WIM file
func SetWimInfo(wim *C.WIMStruct, info *C.struct_wimlib_wim_info, which int) error {
	ret := C.wimlib_set_wim_info(wim, info, C.int(which))
	if ret != 0 {
		return &WimlibError{Code: int(ret)}
	}
	return nil
}

// Split splits a WIM file into multiple parts
func Split(wim *C.WIMStruct, swmName string, partSize uint64, writeFlags int) error {
	cSwmName, freeSwmName := convertToTchar(swmName)
	defer freeSwmName()

	ret := C.wimlib_split(wim, cSwmName, C.uint64_t(partSize), C.int(writeFlags))
	if ret != 0 {
		return &WimlibError{Code: int(ret)}
	}
	return nil
}

// VerifyWim verifies the integrity of a WIM file
func VerifyWim(wim *C.WIMStruct, verifyFlags int) error {
	ret := C.wimlib_verify_wim(wim, C.int(verifyFlags))
	if ret != 0 {
		return &WimlibError{Code: int(ret)}
	}
	return nil
}

// UnmountImage unmounts a WIM image
func UnmountImage(dir string, unmountFlags int) error {
	cDir, freeDir := convertToTchar(dir)
	defer freeDir()

	ret := C.wimlib_unmount_image(cDir, C.int(unmountFlags))
	if ret != 0 {
		return &WimlibError{Code: int(ret)}
	}
	return nil
}

// UpdateImage updates a WIM image
func UpdateImage(wim *C.WIMStruct, image int, cmds []C.struct_wimlib_update_command, updateFlags int) error {
	ret := C.wimlib_update_image(wim, C.int(image), &cmds[0], C.size_t(len(cmds)), C.int(updateFlags))
	if ret != 0 {
		return &WimlibError{Code: int(ret)}
	}
	return nil
}

// WriteWim writes a WIM file
func WriteWim(wim *C.WIMStruct, path string, image int, writeFlags int, numThreads uint) error {
	cPath, freePath := convertToTchar(path)
	defer freePath()

	ret := C.wimlib_write(wim, cPath, C.int(image), C.int(writeFlags), C.unsigned(numThreads))
	if ret != 0 {
		return &WimlibError{Code: int(ret)}
	}
	return nil
}

// WriteToFd writes a WIM file to a file descriptor
func WriteToFd(wim *C.WIMStruct, fd int, image int, writeFlags int, numThreads uint) error {
	ret := C.wimlib_write_to_fd(wim, C.int(fd), C.int(image), C.int(writeFlags), C.unsigned(numThreads))
	if ret != 0 {
		return &WimlibError{Code: int(ret)}
	}
	return nil
}

// SetDefaultCompressionLevel sets the default compression level for a compression type
func SetDefaultCompressionLevel(ctype int, compressionLevel uint) error {
	ret := C.wimlib_set_default_compression_level(C.int(ctype), C.unsigned(compressionLevel))
	if ret != 0 {
		return &WimlibError{Code: int(ret)}
	}
	return nil
}

// GetCompressorNeededMemory gets the amount of memory needed for a compressor
func GetCompressorNeededMemory(ctype int, maxBlockSize uint, compressionLevel uint) uint64 {
	return uint64(C.wimlib_get_compressor_needed_memory(C.enum_wimlib_compression_type(ctype), C.size_t(maxBlockSize), C.unsigned(compressionLevel)))
}

// CreateCompressor creates a compressor
func CreateCompressor(ctype int, maxBlockSize uint, compressionLevel uint) (*C.struct_wimlib_compressor, error) {
	var compressor *C.struct_wimlib_compressor
	ret := C.wimlib_create_compressor(C.enum_wimlib_compression_type(ctype), C.size_t(maxBlockSize), C.unsigned(compressionLevel), &compressor)
	if ret != 0 {
		return nil, &WimlibError{Code: int(ret)}
	}
	return compressor, nil
}

// Compress compresses data using a compressor
func Compress(uncompressedData []byte, compressor *C.struct_wimlib_compressor) ([]byte, error) {
	compressedSize := C.size_t(len(uncompressedData) + 1000) // Add some extra space for compression overhead
	compressedData := make([]byte, compressedSize)

	actualSize := C.wimlib_compress(
		unsafe.Pointer(&uncompressedData[0]),
		C.size_t(len(uncompressedData)),
		unsafe.Pointer(&compressedData[0]),
		C.size_t(compressedSize),
		compressor,
	)

	if actualSize == 0 {
		return nil, &WimlibError{Code: -1} // Compression failed
	}

	return compressedData[:actualSize], nil
}

// FreeCompressor frees a compressor
func FreeCompressor(compressor *C.struct_wimlib_compressor) {
	C.wimlib_free_compressor(compressor)
}

// CreateDecompressor creates a decompressor
func CreateDecompressor(ctype int, maxBlockSize uint) (*C.struct_wimlib_decompressor, error) {
	var decompressor *C.struct_wimlib_decompressor
	ret := C.wimlib_create_decompressor(C.enum_wimlib_compression_type(ctype), C.size_t(maxBlockSize), &decompressor)
	if ret != 0 {
		return nil, &WimlibError{Code: int(ret)}
	}
	return decompressor, nil
}

// Decompress decompresses data using a decompressor
func Decompress(compressedData []byte, uncompressedSize uint, decompressor *C.struct_wimlib_decompressor) ([]byte, error) {
	uncompressedData := make([]byte, uncompressedSize)

	ret := C.wimlib_decompress(
		unsafe.Pointer(&compressedData[0]),
		C.size_t(len(compressedData)),
		unsafe.Pointer(&uncompressedData[0]),
		C.size_t(uncompressedSize),
		decompressor,
	)

	if ret != 0 {
		return nil, &WimlibError{Code: int(ret)}
	}

	return uncompressedData, nil
}

// FreeDecompressor frees a decompressor
func FreeDecompressor(decompressor *C.struct_wimlib_decompressor) {
	C.wimlib_free_decompressor(decompressor)
}
