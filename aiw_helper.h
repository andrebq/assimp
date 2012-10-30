#ifndef _HELPER_H_
#define _HELPER_H_

#include <assimp/scene.h>
#include <assimp/mesh.h>
#include <assimp/vector3.h>

struct aiVector3D* aiw_read_vec(struct aiMesh* m, unsigned int idx);
struct aiMesh* aiw_read_mesh(struct aiScene* s, unsigned int idx);

#endif