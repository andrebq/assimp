package enc

import (
	"testing"
	"bytes"
	"github.com/andrebq/assimp"
)

// Test the encoding/decoding of gob files
func randomScene() *assimp.Scene {
	sc := &assimp.Scene{}
	mesh := &assimpMesh{}
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

	sc, err = GobRead(buff)
	if err != nil {
		t.Fatalf("Error reading scene. %v", err)
	}
	if sc.Mesh[0].Vertices[0] != v {
		t.Errorf("Expecting %v but got %v", v, sc.Mesh[0].Vertices[0])
	}
}