#include "aiw_helper.h"
#include <assimp/scene.h>
#include <assimp/mesh.h>
#include <assimp/vector3.h>
#include <assimp/cimport.h>
#include <assimp/postprocess.h>
#include <stdlib.h>

// return the vector information on the given index
struct aiVector3D* aiw_read_vec(struct aiMesh* m, unsigned int index)
{
	return &(m->mVertices[index]);
}

// return the normal information on the given index
struct aiVector3D* aiw_read_norm(struct aiMesh* m, unsigned int index)
{
	return &(m->mNormals[index]);
}

// return the mesh on the given index
struct aiMesh* aiw_read_mesh(struct aiScene* s, unsigned int index)
{
	return s->mMeshes[index];
}

// return the face on the given index
struct aiFace* aiw_read_face(struct aiMesh* m, unsigned int index)
{
	return &(m->mFaces[index]);
}

// Return the vertex index (from the mesh list) that correspond's to the given face index.
//
// I.e: given the face 0 of the mesh A, give the index in the mesh vector array that correspond's to the face's first element.
unsigned int aiw_read_vec_index_from_face(struct aiFace* f, unsigned int index)
{
	return f->mIndices[index];
}

// Release all resources used by the scene import
void aiw_release_scene(struct aiScene* s)
{
	// no need to call free(s) after calling aiReleaseImport since inside that code the scene object is already deleted.
	//
	// calling free on s after calling aiReleaseImport generate a double-free error and the program is aborted.
	aiReleaseImport(s);
}

// Load the scene from the given file
const struct aiScene* aiw_import_file(const char* file, unsigned int flags)
{
		return aiImportFile(file, aiProcessPreset_TargetRealtime_MaxQuality);
}