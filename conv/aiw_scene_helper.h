#ifndef _AIW_SCENE_HELPER_H
#define _AIW_SCENE_HELPER_H

#include <assimp/scene.h>

// Resource loading related functions

// return the mesh on the given index
struct aiMesh* aiw_read_mesh(struct aiScene* s, unsigned int index);

// Release all resources used by the scene import
//
// This function isn't strictly necessary, but keeping all interaction in one single file make easier to maintain the code in the longrun
void aiw_release_scene(struct aiScene* s);

// Load the scene from the given file
const struct aiScene* aiw_import_file(const char* file);

#endif