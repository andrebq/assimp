package assimp

import (
	"bytes"
	"math/rand"
	"reflect"
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
	// just to make sure that everything is multiple of
	// 3 (x,y,z)
	vecInfo := make([]float32, int(ByteSize/3)*3)
	normInfo := make([]float32, len(vecInfo))
	// for each triplet in vecInfo (x,y,z), include an extra value
	// to hold the alpha value
	colorInfo := make([]float32, len(vecInfo)/3*4)

	for i, ci := 0, 0; i < len(vecInfo); i, ci = i+3, ci+4 {
		for j := 0; j < 3; j++ {
			vecInfo[i+j] = rand.Float32()
			normInfo[i+j] = rand.Float32()
			colorInfo[ci+j] = rand.Float32()
			if j == 2 {
				colorInfo[ci+j+1] = 1 // alpha
			}
		}
	}

	m.Vertices = make([]Vector3, 0)
	m.Normals = make([]Vector3, 0)
	m.Colors = make([]Vector4, 0)

	for i, ci := 0, 0; i < len(vecInfo); i, ci = i+3, ci+4 {
		vec := Vector3{
			float64(vecInfo[i]),
			float64(vecInfo[i+1]),
			float64(vecInfo[i+2])}

		nor := Vector3{
			float64(normInfo[i]),
			float64(normInfo[i+1]),
			float64(normInfo[i+2])}

		col := Vector4{
			float64(colorInfo[ci]),
			float64(colorInfo[ci+1]),
			float64(colorInfo[ci+2]),
			float64(colorInfo[ci+3])}

		m.Vertices = append(m.Vertices, vec)
		m.Normals = append(m.Normals, nor)
		m.Colors = append(m.Colors, col)
	}
	m.Faces = make([]*Face, 0)

	expIdx := make([]byte, 0)
	for i := 0; i < 100; i++ {
		f := &Face{make([]int, 3)}
		f.Indices[0] = rand.Int() % len(m.Vertices)
		f.Indices[1] = rand.Int() % len(m.Vertices)
		f.Indices[2] = rand.Int() % len(m.Vertices)

		expIdx = append(expIdx, byte(f.Indices[0]))
		expIdx = append(expIdx, byte(f.Indices[1]))
		expIdx = append(expIdx, byte(f.Indices[2]))

		m.Faces = append(m.Faces, f)
	}

	fm := NewFlatMesh(m)
	if !reflect.DeepEqual(fm.Vertex, vecInfo) {
		t.Errorf("Vertex array is different")
	}
	if !reflect.DeepEqual(fm.Normal, normInfo) {
		t.Errorf("Normal array is different")
	}
	if !reflect.DeepEqual(fm.Color, colorInfo) {
		t.Errorf("Color array is different.")
		if len(fm.Color) != len(colorInfo) {
			t.Errorf("Color array have a different size")
		} else {
			for i, v := range colorInfo {
				if v != fm.Color[i] {
					t.Errorf("Difference started at index: %v", i)
					t.Errorf("Expected: %v, got %v", colorInfo[i:i+4], fm.Color[i:i+4])
					break
				}
			}
		}
	}

	if !reflect.DeepEqual(fm.ByteIndex, expIdx) {
		t.Errorf("Index array is different")
	}
}
