// This file contains the strucutres used in the program
package assimp

/*Copyright (c) 2012 André Luiz Alves Moraes

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.*/

import (
	"math/rand"
)

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
	Colors []Vector4

	// List of faces
	Faces []*Face

	// List of UV Coordinates (at this point, only one texture for each mesh)
	UVCoords []Vector2

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

// A 3D Vertex
type Vector3 [3]float64

// A 4D Vertex
type Vector4 [4]float64

// A 2D Vertex
type Vector2 [2]float64

// This structure is optimized to be used with
// OpenGL. The vertex information is flat and can be passed
// directly to OpenGL API.
//
// Vertex and Normals are 3 components (x,y,z)
//
// Colors are 4 components (r,g,b,a)
//
// Face index can be a list of: int8 or int16 or int (usually it's int16),
// this is used to reduce the amount of data sent to the VRAM, the user
// must check the value of index size to discover that property should be used
//
// 32 bit floats are used instead of 64 since most of the time 32 bit's have
// enough space to hold most geometries
type FlatMesh struct {
	Vertex  []float32
	Normal  []float32
	Color   []float32
	Texture []float32
	Index   []uint32
}

// Return a flat representation of the given mesh
func NewFlatMesh(m *Mesh) *FlatMesh {
	fm := &FlatMesh{}
	colorInfo := m.Colors != nil && len(m.Colors) > 0
	texInfo := m.UVCoords != nil && len(m.UVCoords) > 0

	fm.Vertex = make([]float32, len(m.Vertices)*3)
	fm.Normal = make([]float32, len(fm.Vertex))
	if colorInfo {
		fm.Color = make([]float32, len(m.Colors)*4)
	}
	if texInfo {
		fm.Texture = make([]float32, len(m.UVCoords)*2)
	}

	for i, v := range m.Vertices {
		fm.Vertex[i*3] = float32(v[0])
		fm.Vertex[i*3+1] = float32(v[1])
		fm.Vertex[i*3+2] = float32(v[2])

		n := m.Normals[i]
		fm.Normal[i*3] = float32(n[0])
		fm.Normal[i*3+1] = float32(n[1])
		fm.Normal[i*3+2] = float32(n[2])

		if colorInfo {
			c := m.Colors[i]
			fm.Color[i*4] = float32(c[0])
			fm.Color[i*4+1] = float32(c[1])
			fm.Color[i*4+2] = float32(c[2])
			fm.Color[i*4+3] = float32(c[3])
		}

		if texInfo {
			t := m.UVCoords[i]
			fm.Texture[i*2] = float32(t[0])
			fm.Texture[i*2+1] = float32(t[1])
		}
	}

	fm.Index = make([]uint32, 0)
	for _, f := range m.Faces {
		for _, i := range f.Indices {
			fm.Index = append(fm.Index, uint32(i))
		}
	}

	return fm
}

// Return a copy of the mesh index array with elements of the size informed.
//
// No overflow is checked, so if you request the ByteSize and there are indexes with values greater than
// 255, the value will be truncated and the mesh will be deformed.
func (fm *FlatMesh) FillIndexArray(size IndexSize) interface{} {
	switch size {
	case ByteSize:
		out := make([]byte, len(fm.Index))
		for i, _ := range out {
			out[i] = byte(fm.Index[i])
		}
		return out
	case ShortSize:
		out := make([]uint16, len(fm.Index))
		for i, _ := range out {
			out[i] = uint16(fm.Index[i])
		}
		return out
	case IntSize:
		out := make([]uint32, len(fm.Index))
		// better copy to keep consistent with the other options
		copy(out, fm.Index)
		return out
	}
	return nil
}

// Just fill the mesh color array with some random colors
func RandomColor(m *Mesh) {
	if len(m.Colors) == 0 {
		m.Colors = make([]Vector4, len(m.Vertices))
	}

	for i, _ := range m.Colors {
		m.Colors[i] = Vector4{
			rand.Float64(),
			rand.Float64(),
			rand.Float64(),
			1, // 100% opacity
		}
	}
}

// Size of the vertex index
type IndexSize int

const (
	// All indexes are under 255
	ByteSize IndexSize = 255
	// All indexes are under 65536
	ShortSize = 32767
	// Fall back
	IntSize = 2147483647
)
