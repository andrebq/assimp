// Open Asset Importer -> Sample OpenGL Viewer
//
// This is a sample application to check if cgo is capable of linking to Open Asset Importer (http://assimp.sourceforge.net/)
//
// The code here is basically a translation from the Sample_SimpleOpenGL.c
//
// You will need GLFW, GL and GLU installed.
package main

import (
	"os"
	"fmt"
	"time"
	"flag"
	
	// opengl related imports
	"github.com/banthar/gl"
	"github.com/banthar/glu"
	"github.com/jteeuwen/glfw"
)

// start the glfw library
func initGlfw() {
	if err := glfw.Init(); err != nil {
		fmt.Fprintf(os.Stderr, "[e] %v\n", err)
		os.Exit(1)
	}
	
	glfw.SetSwapInterval(1)
	glfw.SetWindowTitle("Simple Assim Go binding test")
	glfw.SetWindowSizeCallback(onResize)
}

// handle the resize of the window
func onResize(w, h int) {
	if h == 0 {
		h = 1
	}

	gl.Viewport(0, 0, w, h)
	gl.MatrixMode(gl.PROJECTION)
	gl.LoadIdentity()
	glu.Perspective(45.0, float64(w)/float64(h), 0.1, 100.0)
	gl.MatrixMode(gl.MODELVIEW)
	gl.LoadIdentity()
}

// initialize the opengl context
func initGl() {

	gl.ClearColor(0.1,0.1,0.1,1.0)

	gl.Enable(gl.LIGHTING)
	gl.Enable(gl.LIGHT0)    // Uses default lighting parameters

	gl.Enable(gl.DEPTH_TEST)

	gl.LightModeli(gl.LIGHT_MODEL_TWO_SIDE, gl.TRUE)
	gl.Enable(gl.NORMALIZE)

	gl.ColorMaterial(gl.FRONT_AND_BACK, gl.DIFFUSE)
}

// draw the mesh on the window
func drawScene(sc *Scene) {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	gl.LoadIdentity()
	glu.LookAt(0.0, 0.0, 3.0, 0.0, 0.0, -5.0, 0.0, 1.0, 0.0);
	
	gl.Translatef(0,0,0)

	/*gl.Begin(gl.TRIANGLES)
	gl.Color3f(0.5, 0.5, 1.0)
	gl.Vertex3f(1.0, 1.0, 0)
	gl.Vertex3f(0.0, 1.0, 0)
	gl.Vertex3f(1.0, 0.0, 0)
	gl.End()

	glfw.SwapBuffers()
	*/
	for _, m := range sc.Mesh {
		for _, f := range m.Faces {
			if f == nil || f.Indices == nil { continue }
			mode := gl.POLYGON
			switch len(f.Indices) {
				case 1: mode = gl.POINT
				case 2: mode = gl.LINE
				case 3: mode = gl.TRIANGLES
			}
			gl.Begin(gl.GLenum(mode))
			
			if m.HasNormals() {
				gl.Enable(gl.LIGHTING)
			} else {
				gl.Disable(gl.LIGHTING)
			}
			
			for _, vI := range f.Indices {				
				gl.Color3b(100, 100, 100)
				
				if m.HasNormals() {
					gl.Normal3dv(m.Normals[vI][:])
				}
				gl.Vertex3dv(m.Vertices[vI][:])
			}
			
			
			gl.End()
		}
	}
	glfw.SwapBuffers()
}

// run the program until the user close's the window or 
// the given timeout is reached.
func loop(timeout <-chan time.Time, scene *Scene) {
	for {
		select {
			case <-timeout:
				return
			default:
				if glfw.WindowParam(glfw.Opened) != 1 {
					fmt.Printf("out now...\n")
					return
				}
				drawScene(scene)
		}
	}
}

// open the glfw window
func openWindow() {
	if err := glfw.OpenWindow(800, 600, 8, 8, 8, 8, 0, 8, glfw.Windowed); err != nil {
		fmt.Fprintf(os.Stderr, "[e] %v\n", err)
		os.Exit(1)
	}
}

// Load the scene into the system.
func loadScene(file string) (*Scene, error) {
	return loadAsset(file)
}

func main() {
	flag.Parse()
	
	if len(flag.Args()) == 0 {
		fmt.Fprintf(os.Stderr, "You must provide the name of the file to import.\n")
		os.Exit(1)
	}
	
	initGlfw()
	
	defer glfw.Terminate()
	
	initGl()
	
	c := time.Tick(60 * time.Second)
	
	openWindow()
	
	s, err := loadScene(flag.Args()[0])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading scene: %v\n", err)
		os.Exit(1)
	}
	loop(c, s)
}
