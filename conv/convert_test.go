package conv

/*Copyright (c) 2012 André Luiz Alves Moraes

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.*/

import (
	"math"
	"path/filepath"
	"testing"
)

// Do a simple round on a float value
// if f is more than or equal to 0.5 return the Ceil
// if f is less than 0.5 return the Floor
func simpleRound(f float64) int64 {
	if f >= 0.5 {
		return int64(math.Ceil(f))
	}

	return int64(math.Floor(f))
}

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

	// blender export a lot of vertex data to collada format
	if len(cube.Vertices) != 24 {
		t.Errorf("Expecting %v vertex but got %v", 24, len(cube.Vertices))
	}

	if len(cube.Colors) != 24 {
		t.Errorf("Expecting %v colors but got %v", 24, len(cube.Colors))
	}

	countRed, countBlue := 12, 12
	for _, v := range cube.Colors {
		val := simpleRound(v[0])*1 + simpleRound(v[1])*2 + simpleRound(v[2])*4
		// ignore the alpha channel for now
		if val == 1 {
			countRed--
		} else if val == 4 {
			countBlue--
		}
	}
	if countBlue != 0 {
		t.Errorf("Expecting %v blue's but got %v", 12, 12-countBlue)
	}
	if countRed != 0 {
		t.Errorf("Expecting %v red's but got %v", 12, 12-countRed)
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
