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
	"github.com/banthar/gl"
	"github.com/banthar/glu"
	"github.com/jteeuwen/glfw"

	"math/rand"
	"os"
)

const (
	Title = "Go-wrapper for Open Asset Importer"
)

var (
	quadAngle float32
	running   bool
	Width     = int(640)
	Height    = int(480)
)

func main() {
	var err error
	if err = glfw.Init(); err != nil {
		log("[e] %v\n", err)
		return
	}

	defer glfw.Terminate()

	if err = glfw.OpenWindow(Width, Height, 8, 8, 8, 8, 0, 8, glfw.Windowed); err != nil {
		log("[e] %v\n", err)
		return
	}

	defer glfw.CloseWindow()

	glfw.SetSwapInterval(1)
	glfw.SetWindowTitle(Title)
	glfw.SetWindowSizeCallback(onResize)
	glfw.SetKeyCallback(onKey)

	if scene, err := loadAsset("cube.dae"); err == nil {
		for _, m := range scene.Mesh {
			randomColors(m)
		}

		initGL()

		running = true
		for running && glfw.WindowParam(glfw.Opened) == 1 {
			drawScene(scene)
		}
	} else {
		log("Unable to load scene. Cause: %v", err)
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
	gl.Init()
	gl.ShadeModel(gl.SMOOTH)
	gl.ClearColor(0, 0, 0, 0)
	gl.ClearDepth(1)
	gl.Enable(gl.DEPTH_TEST)
	gl.DepthFunc(gl.LEQUAL)
	gl.Hint(gl.PERSPECTIVE_CORRECTION_HINT, gl.NICEST)
}

func randomColors(m *Mesh) {
	m.Colors = make([]Vector3, len(m.Vertices))
	for i, _ := range m.Colors {
		m.Colors[i] = Vector3{rand.Float64(), rand.Float64(), rand.Float64()}
	}
}

// Render the mesh using OpenGL immediate mode
func renderMeshImmediate(m *Mesh) {
	gl.Begin(gl.TRIANGLES)
	for _, f := range m.Faces {
		for _, i := range f.Indices {
			v := m.Vertices[i]
			c := m.Colors[i]
			gl.Color3d(c[0], c[1], c[2])
			gl.Vertex3d(v[0], v[1], v[2])
		}
	}
	gl.End()
}

// Render the mesh using OpenGL buffers
func renderMeshWithBuff(m *Mesh) {
	buff := CreateBufferFor(m)
	if buff != nil {
		defer buff.Dispose()
	}
	// we are cheating, we are just checking if the buffer can be created and filled
	renderMeshImmediate(m)
}

func drawScene(scene *Scene) {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

	gl.LoadIdentity()
	gl.Translatef(0, 0, -10)
	gl.Rotatef(quadAngle, 1, 0, 1)

	for _, m := range scene.Mesh {
		//renderMeshWithBuff(m)
		// using vertex buffer isn't working right now
		renderMeshImmediate(m)
	}

	quadAngle -= 0.15

	glfw.SwapBuffers()
}

func log(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, msg, args...)
}
