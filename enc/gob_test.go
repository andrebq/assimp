package enc

/*Copyright (c) 2012 André Luiz Alves Moraes

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.*/

import (
	"bytes"
	"github.com/andrebq/assimp"
	"reflect"
	"testing"
)

// Test the encoding/decoding of gob files
func randomScene() *assimp.Scene {
	sc := &assimp.Scene{}
	mesh := &assimp.Mesh{}
	v := assimp.Vector3{1, 1, 1}
	mesh.Vertices = make([]assimp.Vector3, 0)
	mesh.Vertices = append(mesh.Vertices, v)
	sc.AddMesh(mesh)

	return sc
}

// Test the gob encoding/decoding
func TestGob(t *testing.T) {
	sc := randomScene()

	buff := new(bytes.Buffer)
	err := GobWrite(buff, sc)
	if err != nil {
		t.Fatalf("Error writing scene. %v", err)
	}

	buff = bytes.NewBuffer(buff.Bytes())

	fromDisk, err := GobRead(buff)
	if err != nil {
		t.Fatalf("Error reading scene. %v", err)
	}

	if !reflect.DeepEqual(sc.Mesh[0].Vertices, fromDisk.Mesh[0].Vertices) {
		t.Errorf("Vertex array is different")
	}
}
