package enc

/*Copyright (c) 2012 André Luiz Alves Moraes

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.*/

// Loading/writing FlatMesh objects directly from disk.
//
// Unlike the gob format, this one is designed to load all vertex/normal/color information from a file/block.

import (
	"encoding/binary"
	"github.com/andrebq/assimp"
	"io"
)

// Read at most len(out) float32 values from the given reader
func readFloat32Array(out []float32, r io.Reader) error {
	err := binary.Read(r, binary.BigEndian, &out)
	return err
}

// Read the size of the next value
func readSize(r io.Reader) (int32, error) {
	sz := int32(0)
	err := binary.Read(r, binary.BigEndian, &sz)
	return sz, err
}

// Read a list of uint 32 values
func readUInt32Array(out []uint32, r io.Reader) error {
	err := binary.Read(r, binary.BigEndian, &out)
	return err
}

// Custom memory allocation
type MemAlloc interface {
	// Allocate an array of sz floats32 values
	AllocFloat32(sz int32) []float32

	// Allocate an array of sz uint32 values
	AllocUInt32(sz int32) []uint32

	// Allocate an array of sz bytes
	AllocByte(sz int32) []byte

	// Recycle the given array
	Recycle(data interface{})
}

// Default memory allocator
type defaultMemAlloc struct{}

// Allocate using make
func (d *defaultMemAlloc) AllocFloat32(sz int32) []float32 {
	return make([]float32, sz)
}

// Allocate using make
func (d *defaultMemAlloc) AllocUInt32(sz int32) []uint32 {
	return make([]uint32, sz)
}

// Allocate using make
func (d *defaultMemAlloc) AllocByte(sz int32) []byte {
	return make([]byte, sz)
}

// NO-OP, let the GC make good use of the data
func (d *defaultMemAlloc) Recycle(data interface{}) {
	return
}

// Flat mesh codec
type FlatMeshCodec struct {
	memAlloc MemAlloc
	mesh     *assimp.FlatMesh
}

// Return a new instance of the FlatMeshCodec
//
// If the user don't provide an MemAlloc the default one is used (aka, just call make), the user must pass a valid mesh object. Passing nil will result in error
func NewFlatMeshCodec(alloc MemAlloc, mesh *assimp.FlatMesh) (*FlatMeshCodec, error) {
	if alloc == nil {
		alloc = &defaultMemAlloc{}
	}
	if mesh == nil {
		return nil, Error("Argument mesh cannot be nil")
	}
	return &FlatMeshCodec{memAlloc: alloc, mesh: mesh}, nil
}

// read the vertex data from the given mesh
func (f *FlatMeshCodec) readVertexData(r io.Reader) error {
	sz, err := readSize(r)
	if err != nil {
		return err
	}
	if len(f.mesh.Vertex) <= int(sz) {
		f.mesh.Vertex = f.memAlloc.AllocFloat32(sz)
	} else if len(f.mesh.Vertex) >= int(sz) {
		f.mesh.Vertex = f.mesh.Vertex[:int(sz)]
	}

	err = readFloat32Array(f.mesh.Vertex, r)
	return err
}

// read normal data from the given mesh
func (f *FlatMeshCodec) readNormalData(r io.Reader) error {
	sz, err := readSize(r)
	if err != nil {
		return err
	}
	if sz == 0 {
		return nil
	}
	if int(sz) != len(f.mesh.Vertex) {
		return VertexSizeIsDifferent
	}
	f.mesh.Normal = f.memAlloc.AllocFloat32(sz)
	err = readFloat32Array(f.mesh.Normal, r)
	return err
}

// read the color information (rgba8888)
//
// note that the color information is stored in on block of 32 bits and then it's converted to 4 float32 values
func (f *FlatMeshCodec) readColorData(r io.Reader) error {
	sz, err := readSize(r)
	if err != nil {
		return err
	}

	// after normalization, each color use 4 floats
	// and during the normalization, each color use 4 bytes
	normSz := sz * 4

	// each color take 4 bytes (rgba)
	colors := f.memAlloc.AllocByte(normSz)
	defer f.memAlloc.Recycle(colors)

	if len(f.mesh.Color) < int(normSz) {
		f.mesh.Color = f.memAlloc.AllocFloat32(normSz)
	} else {
		f.mesh.Color = f.mesh.Color[:int(normSz)]
	}
	_, err = io.ReadFull(r, colors)
	if err != nil {
		return err
	}

	// calculating normal values
	for i := 0; i < int(normSz); i += 4 {
		f.mesh.Color[i] = float32(colors[i]) / 255
		f.mesh.Color[i+1] = float32(colors[i+1]) / 255
		f.mesh.Color[i+2] = float32(colors[i+2]) / 255
		f.mesh.Color[i+3] = float32(colors[i+3]) / 255
	}

	return err
}

// Read a flat mesh from the given reader and store it under m
func ReadFlatMesh(m *assimp.FlatMesh, memAlloc MemAlloc, r io.Reader) error {
	fm, err := NewFlatMeshCodec(memAlloc, m)
	if err != nil {
		return err
	}

	err = fm.readVertexData(r)
	if err != nil {
		return err
	}

	err = fm.readNormalData(r)
	if err != nil {
		return err
	}

	err = fm.readColorData(r)
	if err != nil {
		return err
	}

	return err
}
