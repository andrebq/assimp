#include "aiw_vertex_helper.h"

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

// return the color at the given colorset/index
//
// Note that each colorset have N colors (where N is the number of vertex in the mesh)
struct aiw_color4d aiw_read_color(struct aiMesh* m, unsigned int colorset, unsigned int idx)
{
	C_STRUCT aiColor4D color = m->mColors[colorset][idx];
	struct aiw_color4d ret = { color.r, color.g, color.b, color.a };
	return ret;
}

// Check if the mesh have colors in the given colorset
int aiw_has_colors(struct aiMesh* m, unsigned int colorset)
{
	if (m->mColors[colorset] != NULL) {
		return 1;
	} else {
		return 0;
	}
}

// return 1 if the mesh has normals and 0 if not
int aiw_mesh_has_normals(struct aiMesh* m)
{
	if (m->mNormals != NULL) { return 1; }
	else { return 0; }
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