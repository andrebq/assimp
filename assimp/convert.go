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
	"errors"
	"fmt"
	"github.com/andrebq/assimp"
	"unsafe"
)

// Convert a Scene from Assimp to a Go structure.
func convertAiScene(scenePtr unsafe.Pointer) (gScene *assimp.Scene) {
	gScene = &assimp.Scene{}
	cScene := (*C.struct_aiScene)(scenePtr)
	gScene.Mesh = make([]*assimp.Mesh, 0, cScene.mNumMeshes)
	convertAiMesh(gScene, scenePtr)
	return
}

// convert all the mesh objects from the scene
func convertAiMesh(gScene *assimp.Scene, scenePtr unsafe.Pointer) {
	cScene := (*C.struct_aiScene)(scenePtr)
	numMeshes := uint(cScene.mNumMeshes)
	for i := uint(0); i < numMeshes; i++ {
		gMesh := &assimp.Mesh{}
		gScene.AddMesh(gMesh)
		cMesh := (*C.struct_aiMesh)(C.aiw_read_mesh(cScene, C.uint(i)))

		// reading mesh vertices
		gMesh.Vertices = make([]assimp.Vector3, cMesh.mNumVertices)
		for i, _ := range gMesh.Vertices {
			cVector3d := (*C.struct_aiVector3D)(C.aiw_read_vec(cMesh, C.uint(i)))
			gMesh.Vertices[i][0] = float64(cVector3d.x)
			gMesh.Vertices[i][1] = float64(cVector3d.y)
			gMesh.Vertices[i][2] = float64(cVector3d.z)
		}

		// reading mesh normals
		if int(C.aiw_mesh_has_normals(cMesh)) == 1 {
			gMesh.Normals = make([]assimp.Vector3, int(cMesh.mNumVertices))
			for i, _ := range gMesh.Normals {
				cVector3d := (*C.struct_aiVector3D)(C.aiw_read_norm(cMesh, C.uint(i)))
				gMesh.Normals[i][0] = float64(cVector3d.x)
				gMesh.Normals[i][1] = float64(cVector3d.y)
				gMesh.Normals[i][2] = float64(cVector3d.z)
			}
		}

		// reading mesh faces
		gMesh.Faces = make([]*assimp.Face, int(cMesh.mNumFaces))
		for i, _ := range gMesh.Faces {
			cFace := (*C.struct_aiFace)(C.aiw_read_face(cMesh, C.uint(i)))
			gFace := &assimp.Face{Indices: make([]int, int(cFace.mNumIndices))}
			for j, _ := range gFace.Indices {
				gFace.Indices[j] = int(C.aiw_read_vec_index_from_face(cFace, C.uint(j)))
			}
			gMesh.Faces[i] = gFace
		}
	}
}

// Load the assent from the given path.
//
// All resources are dealocated before returning.
func loadAsset(path string) (*assimp.Scene, error) {
	var err error
	cs := C.CString(path)
	defer C.free(unsafe.Pointer(cs))
	cScene := C.aiw_import_file(cs, C.aiProcessPreset_TargetRealtime_MaxQuality)
	if uintptr(unsafe.Pointer(cScene)) == 0 {
		err = errors.New(fmt.Sprintf("Unable to load %v.\n", path))
		return nil, err
	}
	defer func() {
		C.aiw_release_scene(cScene)
	}()

	return convertAiScene(unsafe.Pointer(cScene)), nil
}
