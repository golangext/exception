package exception

import "io"

// ErrShortWrite means that a write accepted fewer bytes than requested
// but failed to return an explicit error.
var ErrShortWrite = IO.extendErrObj("ErrShortWrite", io.ErrShortWrite)

// ErrShortBuffer means that a read required a longer buffer than was provided.
var ErrShortBuffer = IO.extendErrObj("ErrShortBuffer", io.ErrShortBuffer)

// EOF is the error returned by Read when no more input is available.
// Functions should return EOF only to signal a graceful end of input.
// If the EOF occurs unexpectedly in a structured data stream,
// the appropriate error is either ErrUnexpectedEOF or some other error
// giving more detail.
var EOF = IO.extendErrObj("EOF", io.EOF)

// ErrUnexpectedEOF means that EOF was encountered in the
// middle of reading a fixed-size block or data structure.
var ErrUnexpectedEOF = EOF.extendErrObj("ErrUnexpectedEOF", io.ErrUnexpectedEOF)

// ErrNoProgress is returned by some clients of an io.Reader when
// many calls to Read have failed to return any data or error,
// usually the sign of a broken io.Reader implementation.
var ErrNoProgress = IO.extendErrObj("ErrNoProgress", io.ErrNoProgress)
