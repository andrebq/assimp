// This file contains the strucutres used in the program

package main

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
}

// Index of the given mesh
func (s *Scene) IndexOfMesh(m *Mesh) (idx int, has bool) {
	for i, v := range s.Mesh {
		if v == m {
			idx = i
			has = true
			return
		}
	}
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
func (n *Node) UseMesh(m *Mesh, s *Scene) {
}

// Hold the Mesh information
// Vertex, Textures and Normals
type Mesh struct {
}