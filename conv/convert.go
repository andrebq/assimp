// Hold the routines used to convert from C structures to the Go ones
package conv

/*Copyright (c) 2012 André Luiz Alves Moraes

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.*/

// #cgo pkg-config: assimp
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

		if C.aiw_has_colors(cMesh, C.uint(0)) > 0 {
			gMesh.Colors = make([]assimp.Vector4, int(cMesh.mNumVertices))
			for i, _ := range gMesh.Colors {
				cColor := C.aiw_read_color(cMesh, C.uint(0), C.uint(i))
				gMesh.Colors[i] = assimp.Vector4{
					float64(cColor.r),
					float64(cColor.g),
					float64(cColor.b),
					float64(cColor.a),
				}
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
func LoadAsset(path string) (*assimp.Scene, error) {
	var err error
	cs := C.CString(path)
	defer C.free(unsafe.Pointer(cs))
	cScene := C.aiw_import_file(cs)
	if uintptr(unsafe.Pointer(cScene)) == 0 {
		err = errors.New(fmt.Sprintf("Unable to load %v.\n", path))
		return nil, err
	}
	defer func() {
		C.aiw_release_scene(cScene)
	}()

	return convertAiScene(unsafe.Pointer(cScene)), nil
}
