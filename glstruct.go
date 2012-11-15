package main

import (
	"github.com/banthar/gl"
)

type VertexBuf struct {
	buf gl.Buffer
	mid int
}

// Create a new VertexBuffer for the given mesh and fill's it with
// the vertex information from the mesh
func CreateBufferFor(m *Mesh) *VertexBuf {
	vbuf := &VertexBuf{}
	vbuf.mid = m.Id()
	vbuf.buf = gl.GenBuffer()

	vbuf.buf.Bind(gl.ARRAY_BUFFER)
	err := gl.GetError()
	if err != gl.NO_ERROR {
		log("Error binding vertex buffer. GL_CODE: %v", err)
		return nil
	}

	gl.BufferData(gl.VERTEX_ARRAY_BUFFER_BINDING, len(m.Vertices), &m.Vertices, gl.STATIC_DRAW)
	if gl.GetError() != 0 {
		log("Error writing buffer data. GL_CODE: %v", gl.GetError())
	}

	return vbuf
}

// Release the VRAM allocated for the buffer
func (v *VertexBuf) Dispose() {
	v.buf.Delete()
}
