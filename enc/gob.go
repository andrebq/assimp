package enc

import (
	"encoding/gob"
	"io"
)

// Write the file using the gob format
func GobWrite(w io.Writer, s *assimp.Scene) {
	enc := gob.NewEncoder(w)
	err := enc.Encode(*s)
	return err
}

// Read a file written using GobWrite function
func GobRead(r io.Reader) *assimp.Scene {
	dec := gob.NewDecoder(r)
	sc := &Scene{}
	err := dec.Decode(sc)
	return sc, err
}