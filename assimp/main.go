package main

/*Copyright (c) 2012 André Luiz Alves Moraes

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.*/

import (
	"flag"
	"fmt"
	"github.com/andrebq/assimp"
	"github.com/andrebq/assimp/conv"
	"io"
	"os"
	"strings"
)

var (
	_if  = flag.String("if", "", "Input file")
	_of  = flag.String("of", "-", "Output file")
	help = flag.Bool("h", false, "Help")
)

func main() {
	flag.Parse()

	if *help {
		printUsage("")
	}

	if *_if == "" {
		printUsage("The input file is required")
	}
	if scene, err := conv.LoadAsset(*_if); err != nil {
		log("Unable to load scene.\nCause: %v", err)
	} else {
		dumpScene(scene, *_of)
	}
}

// Dump a scene loaded from assimp to a gob file
// this file can later be used to load resources into the game
// or manipulated to a faster format.
func dumpScene(s *assimp.Scene, outpath string) {
	w, err := openWriterFor(outpath)
	if err != nil {
		fatal("Error opening %v for write. Cause: %v", outpath, err)
	}
	if w, ok := w.(io.Closer); ok {
		defer w.Close()
	}
}

func openWriterFor(file string) (io.Writer, error) {
	if file == "-" {
		return os.Stdout, nil
	} else {
		f, err := os.Create(file)
		if err != nil {
			return nil, err
		}
		return f, err
	}
	panic("Not reached")
	return nil, nil
}

// Just log some information
func log(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, msg, args...)
	if !strings.HasSuffix(msg, "\n") || !strings.HasSuffix(msg, "\r\n") {
		fmt.Fprintf(os.Stderr, "\n")
	}
}

// just like log, but call's os.Exit(1) after
func fatal(msg string, args ...interface{}) {
	log(msg, args...)
	os.Exit(1)
}

// print usage
func printUsage(msg string) {
	if msg != "" {
		log(msg)
	}
	flag.Usage()
	os.Exit(1)
}
