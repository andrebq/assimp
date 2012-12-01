package main

/*Copyright (c) 2012 André Luiz Alves Moraes

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.*/

// Error type
type Error string

// Return the string representation
func (e Error) String() string {
	return string(e)
}

// error interface
func (e Error) Error() string {
	return string(e)
}

// most common error
const (
	UnexpectedEOF = Error("Unexpected end-of-file")
	ZeroOrInvalidSize = Error("Zero or invalid size of data")
	VertexSizeIsDifferent = Error("Size from file incompatible with vertex count")
)