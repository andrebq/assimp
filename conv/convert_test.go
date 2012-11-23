package conv

import (
	"path/filepath"
	"testing"
)

// Try to load a simple cube from a collada file
func TestCubeDae(t *testing.T) {
	assetName := filepath.Join(filepath.FromSlash("../data/"), "cube.dae")
	gScene, err := LoadAsset(assetName)
	if err != nil {
		t.Fatalf("Unable to load %v. Cause %v", assetName, err)
	}

	if len(gScene.Mesh) != 1 {
		t.Fatalf("Expecting %v mesh but got %v", 1, len(gScene.Mesh))
	}

	cube := gScene.Mesh[0]
	// 12 faces since each square uses 2 triangles
	if len(cube.Faces) != 12 {
		t.Errorf("Expecting %v faces but got %v", 12, len(cube.Faces))
	}

	for i, f := range cube.Faces {
		// each face must have only 3 indexes
		// and all indexes must point to valid mesh
		// vertexes
		if len(f.Indices) != 3 {
			t.Errorf("Face %v should have %v indices but got %v indices", i, 3, len(f.Indices))
		}
		for j, vIdx := range f.Indices {
			if vIdx < 0 || vIdx > len(cube.Vertices)-1 {
				t.Errorf("Face index %v value is %v but the mesh don't have that vertex", j, vIdx)
			}
		}
	}
}
