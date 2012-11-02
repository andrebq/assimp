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
	// open gl related
	"github.com/jteeuwen/glfw"
	"github.com/banthar/gl"
	"github.com/banthar/glu"
	
	"math/rand"
	"os"
)

const (
	Title  = "Go-wrapper for Open Asset Importer"
)

var (
	quadAngle float32
	running   bool
	Width  = int(640)
	Height = int(480)
)

func main() {
	var err error
	if err = glfw.Init(); err != nil {
		fmt.Fprintf(os.Stderr, "[e] %v\n", err)
		return
	}

	defer glfw.Terminate()

	if err = glfw.OpenWindow(Width, Height, 8, 8, 8, 8, 0, 8, glfw.Windowed); err != nil {
		fmt.Fprintf(os.Stderr, "[e] %v\n", err)
		return
	}

	defer glfw.CloseWindow()

	glfw.SetSwapInterval(1)
	glfw.SetWindowTitle(Title)
	glfw.SetWindowSizeCallback(onResize)
	glfw.SetKeyCallback(onKey)
	
	if scene, err := loadAsset("cube.dae"); err == nil {
		initGL()

		running = true
		for running && glfw.WindowParam(glfw.Opened) == 1 {
			drawScene(scene)
		}
	} else {
		fmt.Fprintf(os.Stderr, "Unable to load scene. Cause: %v", err)
	}
}

func onResize(w, h int) {
	if h == 0 {
		h = 1
	}
	
	Height = h
	Width = w

	gl.Viewport(0, 0, Width, Height)
	gl.MatrixMode(gl.PROJECTION)
	gl.LoadIdentity()
	glu.Perspective(45.0, float64(Width)/float64(Height), 0.1, 100.0)
	gl.MatrixMode(gl.MODELVIEW)
	gl.LoadIdentity()
}

func onKey(key, state int) {
	switch key {
	case glfw.KeyEsc:
		running = false
	}
}

func initGL() {
	gl.ShadeModel(gl.SMOOTH)
	gl.ClearColor(0, 0, 0, 0)
	gl.ClearDepth(1)
	gl.Enable(gl.DEPTH_TEST)
	gl.DepthFunc(gl.LEQUAL)
	gl.Hint(gl.PERSPECTIVE_CORRECTION_HINT, gl.NICEST)
}

func drawScene(scene *Scene) {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

	gl.LoadIdentity()
	gl.Translatef(0, 0, -10)
	gl.Rotatef(quadAngle, 1, 0, 1)

	gl.Begin(gl.TRIANGLES)
	for _, m := range scene.Mesh {
		for _, f := range m.Faces {
			for _, i := range f.Indices {
				if i % 2 == 0 {
					gl.Color3f(rand.Float32(), rand.Float32(), rand.Float32())
				}
				v := m.Vertices[i]
				gl.Vertex3d(v[0], v[1], v[2])
			}
		}
	}
	gl.End()
	
	quadAngle -= 0.15

	glfw.SwapBuffers()
}
