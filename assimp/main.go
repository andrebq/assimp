// Open Asset Importer -> Sample OpenGL Viewer
//
// This is a sample application to check if cgo is capable of linking to Open Asset Importer (http://assimp.sourceforge.net/)
//
// The code here is basically a translation from the Sample_SimpleOpenGL.c
//
// You will need GLFW, GL and GLU installed.
package main

import (
	"fmt"
	"os"
	"strings"
	"github.com/andrebq/assimp"
)

func main() {
	if scene, err := loadAsset("cube.dae"); err != nil {
		log("Unable to load scene. Cause: %v", err)
	} else {
		dumpScene(scene)
	}
}

// Dump a scene loaded from assimp to a gob file
// this file can later be used to load resources into the game
// or manipulated to a faster format.
func dumpScene(s *assimp.Scene) {
	// dummy method here
}

func log(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, msg, args...)
	if !strings.HasSuffix(msg, "\n") || !strings.HasSuffix(msg, "\r\n") {
		fmt.Fprintf(os.Stderr, "\n")
	}
}
