#include "aiw_helper.h"
#include <assimp/scene.h>
#include <assimp/mesh.h>
#include <assimp/vector3.h>
#include <assimp/cimport.h>
#include <assimp/postprocess.h>
#include <stdlib.h>

// return the vector information on the given position
struct aiVector3D* aiw_read_vec(struct aiMesh* m, unsigned int index)
{
	unsigned int offset = index * sizeof(struct aiVector3D);
	return &(m->mVertices[offset]);
}

// return the mesh on the given position
struct aiMesh* aiw_read_mesh(struct aiScene* s, unsigned int idx) {
	return s->mMeshes[idx];
}
