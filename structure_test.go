package assimp

import (
	"bytes"
	"math/rand"
	"testing"
)

// Test the encoding/decoding of gob files
func TestCodec(t *testing.T) {
	sc := &Scene{}
	mesh := &Mesh{}
	v := Vector3{1, 1, 1}
	mesh.Vertices = make([]Vector3, 0)
	mesh.Vertices = append(mesh.Vertices, v)
	sc.AddMesh(mesh)

	buff := new(bytes.Buffer)
	err := WriteScene(buff, sc)
	if err != nil {
		t.Fatalf("Error writing scene. %v", err)
	}

	buff = bytes.NewBuffer(buff.Bytes())

	sc, err = ReadScene(buff)
	if err != nil {
		t.Fatalf("Error reading scene. %v", err)
	}
	if sc.Mesh[0].Vertices[0] != v {
		t.Errorf("Expecting %v but got %v", v, sc.Mesh[0].Vertices[0])
	}
}

// Test the conversion between Mesh and FlatMesh
func TestFlatMesh(t *testing.T) {
	m := &Mesh{}
	vecInfo := make([]float32, ByteSize*3)
	normInfo := make([]float32, len(vecInfo))
	colorInfo := make([]float32, ByteSize*4)
	for i, _ := range vecInfo {
		vecInfo[i] = rand.Float32()
		normInfo[i] = rand.Float32()
		colorInfo[i] = rand.Float32()
		// the value of colors aren't important
		colorInfo[i+1] = rand.Float32()
	}

	m.Vertices = make([]Vector3, 0)
	m.Normals = make([]Vector3, 0)
	m.Colors = make([]Vector4, 0)

	for i, ci := 0, 0; i < len(vecInfo); i, ci = i+3, ci+4 {
		vec := Vector3{float64(vecInfo[i]), float64(vecInfo[i+1]), float64(vecInfo[i+2])}
		nor := Vector3{float64(normInfo[i]), float64(normInfo[i+1]), float64(normInfo[i+2])}
		col := Vector4{float64(colorInfo[i]), float64(colorInfo[i+1]), float64(colorInfo[i+2]), float64(colorInfo[i+3])}

		m.Vertices = append(m.Vertices, vec)
		m.Normals = append(m.Normals, nor)
		m.Colors = append(m.Colors, col)
	}
}
