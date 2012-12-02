package enc

/*Copyright (c) 2012 André Luiz Alves Moraes

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.*/

// Write/Read data using the gob format and the default structure

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
