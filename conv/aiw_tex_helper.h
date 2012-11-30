#ifndef _AIW_TEX_HELPER_H
#define _AIW_TEX_HELPER_H

#include <assimp/mesh.h>
#include <assimp/vector3.h>

// Texture related functions

// Check if the mesh have a texture at the given texture set
int aiw_has_texture_at(struct aiMesh* m, unsigned int tex_set);

// Read the texture information for the mesh at the given position
struct aiVector3D* aiw_read_texture_at(struct aiMesh* m, unsigned int tex_set, unsigned int vertex_idx);

#endif