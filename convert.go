// Hold the routines used to convert from C structures to the Go ones
package main

// #cgo pkg-config: assimp
// #include <assimp/scene.h>
// #include <assimp/cimport.h>
// #include <assimp/postprocess.h>
// #include <stdlib.h>
import "C"

import (
	"unsafe"
	"errors"
	"fmt"
)

// Convert a Scene from Assimp to a Go structure.
func convertAiScene(scenePtr unsafe.Pointer) (gScene *Scene) {
	return nil
}

// load the assent in the given path
func loadAsset(path string) (scene unsafe.Pointer, err error) {
	cs := C.CString(path)
	defer C.free(unsafe.Pointer(cs))
	csScene := C.aiImportFile(cs, C.aiProcessPreset_TargetRealtime_MaxQuality)
	if (uintptr(unsafe.Pointer(csScene)) == 0) {
		err = errors.New(fmt.Sprintf("Unable to load %v.\n", path))
	} else {
		scene = unsafe.Pointer(csScene)
	}
	return
}