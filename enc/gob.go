package enc

import (
	"encoding/gob"
	"github.com/andrebq/assimp"
	"io"
)

// Write the file using the gob format
func GobWrite(w io.Writer, s *assimp.Scene) error {
	enc := gob.NewEncoder(w)
	err := enc.Encode(*s)
	return err
}

// Read a file written using GobWrite function
func GobRead(r io.Reader) (*assimp.Scene, error) {
	dec := gob.NewDecoder(r)
	sc := &assimp.Scene{}
	err := dec.Decode(sc)
	return sc, err
}
