// This file contains the strucutres used in the program
package main

// Represent an error in the structure of a scene/node.
type Error string

// Error interface
func (e Error) Error() string {
	return string(e)
}

const (
	ErrMeshNotFound = Error("Unable to find the given mesh in the Scene")
)

// The Scene object, hold the root node of the scene and
// the list of meshes
type Scene struct {
	// Root node of the scene
	Root *Node

	// List of mesh
	Mesh []*Mesh
}

// Push a new mesh into the Scene
func (s *Scene) AddMesh(m *Mesh) {
	if s.Mesh == nil {
		s.Mesh = []*Mesh{m}
	} else {
		s.Mesh = append(s.Mesh, m)
	}
	m.mid = len(s.Mesh)
}

// Index of the given mesh
func (s *Scene) IndexOfMesh(m *Mesh) (idx int, err error) {
	for i, v := range s.Mesh {
		if v == m {
			idx = i
			return
		}
	}
	err = ErrMeshNotFound
	return
}

// One node in the scene.
type Node struct {
	// The list of index used by this node
	Mesh []int

	// Child nodes
	Childs []*Node
}

// Add a mesh into this node
func (n *Node) AddMeshIndex(i int) {
	if n.Mesh == nil {
		n.Mesh = []int{i}
	}
	n.Mesh = append(n.Mesh, i)
}

// Use the given mesh from the given scene
func (n *Node) UseMesh(m *Mesh, s *Scene) (err error) {
	var idx int
	if idx, err = s.IndexOfMesh(m); err == nil {
		n.AddMeshIndex(idx)
	}
	return
}

// Hold the Mesh information
// Vertices, Textures and Normals
type Mesh struct {
	// List of vertices
	Vertices []Vector3

	// List of normals
	Normals []Vector3

	// List of colors
	Colors []Vector3

	// List of faces
	Faces []*Face

	// Mesh id
	mid int
}

// Represent the id of the given mesh in the given scene.
//
// This Id is valid only for Go and isn't loaded from the asset file
func (m *Mesh) Id() int {
	return m.mid
}

// Return true if the mesh has normal information
func (m *Mesh) HasNormals() bool {
	return m.Normals != nil && len(m.Normals) > 0
}

// Hold the information of a single face.
// Hold only the pointers to the vector stored in the Mesh
type Face struct {
	// List of vector indices
	Indices []int
}

// A 3d Vertex
type Vector3 [3]float64
