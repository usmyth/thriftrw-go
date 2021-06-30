// Copyright (c) 2021 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package stream

import (
	"io"

	"go.uber.org/thriftrw/wire"
)

// EnvelopeHeader represents the envelope of a response or a request which includes
// metadata about the method, the type of data in the envelope, and the value.
// It is equivalent of `wire.Envelope`, but for streaming purposes.
type EnvelopeHeader struct {
	Name  string
	Type  wire.EnvelopeType
	SeqID int32
}

// EnvelopeAgnosticProtocol is the streaming equivalent of protocol.EnvelopeAgnosticProtocol
type EnvelopeAgnosticProtocol interface {
	Protocol

	// ReadRequestEnvelope reads off the request envelope (if present) from a Reader
	// and returns a stream.Reader to read the remaining un-enveloped request struct.
	// This allows a Thrift request handler to transparently read requests
	// regardless of whether the caller is configured to submit envelopes.
	// The caller specifies the expected EnvelopeType, either OneWay or Unary,
	// on which the read asserts the specified envelope is present.
	ReadRequestEnvelope(et wire.EnvelopeType, r io.Reader) (Reader, Responder, error)
}

// Responder captures how to respond to a request, concerning whether and what
// kind of envelope style to use
type Responder interface {
	// WriteResponseEnvelope writes a response envelope to the Writer with the envelope
	// style of the corresponding request, and returns a stream.Writer to write
	// remaining un-enveloped response bytes. Once writing of the response is complete,
	// whether successful or not (error), users must call Close() on the stream.Writer.
	//
	// The EnvelopeType should be either wire.Reply or wire.Exception.
	WriteResponseEnvelope(et wire.EnvelopeType, w io.Writer) (Writer, error)
}
