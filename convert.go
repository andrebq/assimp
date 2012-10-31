// Hold the routines used to convert from C structures to the Go ones
package main

// #cgo pkg-config: assimp
// #include <assimp/scene.h>
// #include <assimp/cimport.h>
// #include <assimp/postprocess.h>
// #include <assimp/vector3.h>
// #include <stdlib.h>
// #include "aiw_helper.h"
import "C"

import (
	"unsafe"
	"errors"
	"fmt"
)

// Convert a Scene from Assimp to a Go structure.
func convertAiScene(scenePtr unsafe.Pointer) (gScene *Scene) {
	gScene = &Scene{}
	cScene := (*C.struct_aiScene)(scenePtr)
	gScene.Mesh = make([]*Mesh, 0, cScene.mNumMeshes)
	convertAiMesh(gScene, scenePtr)
	return
}

// convert all the mesh objects from the scene
func convertAiMesh(gScene *Scene, scenePtr unsafe.Pointer) {
	cScene := (*C.struct_aiScene)(scenePtr)
	numMeshes := uint(cScene.mNumMeshes)
	for i := uint(0); i < numMeshes; i++ {
		gMesh := &Mesh{}
		cMesh := (*C.struct_aiMesh)(C.aiw_read_mesh(cScene, C.uint(i)))
		
		// reading mesh vertices
		gMesh.Vertices = make([]Vector3, cMesh.mNumVertices)
		for i, _ := range gMesh.Vertices {
			cVector3d := (*C.struct_aiVector3D)(C.aiw_read_vec(cMesh, C.uint(i)))
			gMesh.Vertices[i][0] = float64(cVector3d.x)
			gMesh.Vertices[i][1] = float64(cVector3d.y)
			gMesh.Vertices[i][2] = float64(cVector3d.z)
		}
		
		// reading mesh faces
		gMesh.Faces = make([]*Face, int(cMesh.mNumFaces))
		for i, _ := range gMesh.Faces {
			cFace := (*C.struct_aiFace)(C.aiw_read_face(cMesh, C.uint(i)))
			gFace := &Face{Indices:make([]int, int(cFace.mNumIndices))}
			for j, _ := range gFace.Indices {
				gFace.Indices[i] = int(C.aiw_read_vec_index_from_face(cFace, C.uint(j)))
				println("face")
			}
		}
	}
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