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

package binary

import (
	"io"

	"go.uber.org/thriftrw/wire"
)

// offsetReader provides a type that satisfies an io.Reader with only an
// io.ReaderAt.
type offsetReader struct {
	offset int64
	reader io.ReaderAt
}

// Read reads len(p) bytes into p.
func (or *offsetReader) Read(p []byte) (int, error) {
	n, err := or.reader.ReadAt(p, or.offset)
	or.offset += int64(n)

	return n, err
}

func newOffsetReader(r io.ReaderAt) offsetReader {
	return offsetReader{reader: r}
}

// reader functions as the actual reader behind the exported `Reader` type.
// This is necessary to avoid new calls to a `Reader.ReadValue` from changing
// the offset in already running 'ReadValue' calls.
type reader struct {
	reader io.ReaderAt
	or     *offsetReader
	sr     StreamReader
}

func newReader(r io.ReaderAt) reader {
	or := offsetReader{reader: r}

	return reader{
		reader: r,
		or:     &or,
		sr:     NewStreamReader(&or),
	}
}

func (br *reader) readStructStream() (wire.Struct, error) {
	var fields []wire.Field

	br.sr.ReadStructBegin()
	fh, hasVal, err := br.sr.ReadFieldBegin()
	if err != nil {
		return wire.Struct{}, err
	}

	var val wire.Value
	for hasVal {
		val, _, err = br.ReadValue(fh.Type, br.or.offset)
		if err != nil {
			return wire.Struct{}, err
		}

		fields = append(fields, wire.Field{ID: fh.ID, Value: val})
		if err = br.sr.ReadFieldEnd(); err != nil {
			return wire.Struct{}, err
		}

		if fh, hasVal, err = br.sr.ReadFieldBegin(); err != nil {
			return wire.Struct{}, err
		}
	}

	if err = br.sr.ReadStructEnd(); err != nil {
		return wire.Struct{}, err
	}

	return wire.Struct{Fields: fields}, nil
}

func (br *reader) readMapStream() (wire.MapItemList, error) {
	mh, err := br.sr.ReadMapBegin()
	if err != nil {
		return nil, err
	}

	start := br.or.offset
	for i := 0; i < mh.Length; i++ {
		if err := br.sr.Skip(mh.KeyType); err != nil {
			return nil, err
		}

		if err := br.sr.Skip(mh.ValueType); err != nil {
			return nil, err
		}
	}

	if err := br.sr.ReadMapEnd(); err != nil {
		return nil, err
	}

	items := borrowLazyMapItemList()
	items.ktype = mh.KeyType
	items.vtype = mh.ValueType
	items.count = int32(mh.Length)
	items.reader = br
	items.startOffset = start

	return items, nil
}

func (br *reader) readListStream() (wire.ValueList, error) {
	lh, err := br.sr.ReadListBegin()
	if err != nil {
		return nil, err
	}

	start := br.or.offset
	for i := 0; i < lh.Length; i++ {
		if err := br.sr.Skip(lh.Type); err != nil {
			return nil, err
		}
	}

	if err = br.sr.ReadListEnd(); err != nil {
		return nil, err
	}

	items := borrowLazyValueList()
	items.count = int32(lh.Length)
	items.typ = lh.Type
	items.reader = br
	items.startOffset = start

	return items, err
}

func (br *reader) readSetStream() (wire.ValueList, error) {
	sh, err := br.sr.ReadSetBegin()
	if err != nil {
		return nil, err
	}

	start := br.or.offset
	for i := 0; i < sh.Length; i++ {
		if err := br.sr.Skip(sh.Type); err != nil {
			return nil, err
		}
	}

	if err = br.sr.ReadSetEnd(); err != nil {
		return nil, err
	}

	items := borrowLazyValueList()
	items.count = int32(sh.Length)
	items.typ = sh.Type
	items.reader = br
	items.startOffset = start

	return items, err
}

// ReadValue is the underlying call made from the exported `Reader.ReadValue`
// that's meant to be safe for concurrent calls.
func (br *reader) ReadValue(t wire.Type, off int64) (wire.Value, int64, error) {
	br.or.offset = off

	switch t {
	case wire.TBool:
		b, err := br.sr.ReadBool()
		return wire.NewValueBool(b), br.or.offset, err

	case wire.TI8:
		b, err := br.sr.ReadInt8()
		return wire.NewValueI8(int8(b)), br.or.offset, err

	case wire.TDouble:
		value, err := br.sr.ReadDouble()
		return wire.NewValueDouble(value), br.or.offset, err

	case wire.TI16:
		n, err := br.sr.ReadInt16()
		return wire.NewValueI16(n), br.or.offset, err

	case wire.TI32:
		n, err := br.sr.ReadInt32()
		return wire.NewValueI32(n), br.or.offset, err

	case wire.TI64:
		n, err := br.sr.ReadInt64()
		return wire.NewValueI64(n), br.or.offset, err

	case wire.TBinary:
		v, err := br.sr.ReadBinary()
		return wire.NewValueBinary(v), br.or.offset, err

	case wire.TStruct:
		s, err := br.readStructStream()
		return wire.NewValueStruct(s), br.or.offset, err

	case wire.TMap:
		m, err := br.readMapStream()
		return wire.NewValueMap(m), br.or.offset, err

	case wire.TSet:
		s, err := br.readSetStream()
		return wire.NewValueSet(s), br.or.offset, err

	case wire.TList:
		l, err := br.readListStream()
		return wire.NewValueList(l), br.or.offset, err

	default:
		return wire.Value{}, br.or.offset, decodeErrorf("unknown ttype %v", t)
	}
}

// Reader implements a parser for the Thrift Binary Protocol based on an
// io.ReaderAt.
type Reader struct {
	reader io.ReaderAt
}

// NewReader builds a new Reader based on the given io.ReaderAt.
func NewReader(r io.ReaderAt) Reader {
	return Reader{reader: r}
}

// ReadValue reads a value off the given type off the wire starting at the
// given offset.
//
// Returns the Value, the new offset, and an error if there was a decode error.
func (br *Reader) ReadValue(t wire.Type, off int64) (wire.Value, int64, error) {
	realReader := newReader(br.reader)
	return realReader.ReadValue(t, off)
}
