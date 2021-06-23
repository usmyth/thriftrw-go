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

// EnvelopeHeader represents the beginning of a response or a request, equivalent
// to a `wire.Envelope`, but for streaming purposes.
//
// A special case of all fields set to their zero-value means that there is no
// enveloping.
type EnvelopeHeader struct {
	Name  string
	Type  wire.EnvelopeType
	SeqID int32
}

// EnvelopeAgnosticProtocol is the streaming equivalent of
// protocol.EnvelopeAgnosticProtocol
type EnvelopeAgnosticProtocol interface {
	Protocol

	ReadRequest(et wire.EnvelopeType, r io.Reader) (Reader, Responder, error)
}

type Responder interface {
	WriteResponse(et wire.EnvelopeType, w io.Writer) (Writer, error)
}
