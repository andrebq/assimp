#ifndef _HELPER_H_
#define _HELPER_H_

#include <assimp/scene.h>
#include <assimp/mesh.h>
#include <assimp/vector3.h>
#include <assimp/color4.h>

// Hold the color information at the given index
struct aiw_color4d {
	float r,g,b,a;
};

// return the vector information on the given index
struct aiVector3D* aiw_read_vec(struct aiMesh* m, unsigned int index);

// return the normal information on the given index
struct aiVector3D* aiw_read_norm(struct aiMesh* m, unsigned int index);

// return the color at the given colorset/index
//
// Note that each colorset have N colors (where N is the number of vertex in the mesh)
struct aiw_color4d aiw_read_color(struct aiMesh* m, unsigned int colorset, unsigned int idx);

// Check if the mesh have colors in the given colorset
int aiw_has_colors(struct aiMesh* m, unsigned int colorset);

// return 1 if the mesh has normals and 0 if not
int aiw_mesh_has_normals(struct aiMesh* m);

// return the mesh on the given index
struct aiMesh* aiw_read_mesh(struct aiScene* s, unsigned int index);

// return the face on the given index
struct aiFace* aiw_read_face(struct aiMesh* m, unsigned int index);

// return the vertex index (from the mesh list) that correspond's to the given face index
//
// I.e: given the face 0 of the mesh A, give the index in the mesh vector array that correspond's to the face's first element
unsigned int aiw_read_vec_index_from_face(struct aiFace* f, unsigned int index);

// Release all resources used by the scene import
//
// This function isn't strictly necessary, but keeping all interaction in one single file make easier to maintain the code in the longrun
void aiw_release_scene(struct aiScene* s);

// Load the scene from the given file
const struct aiScene* aiw_import_file(const char* file, unsigned int flags);

#endif