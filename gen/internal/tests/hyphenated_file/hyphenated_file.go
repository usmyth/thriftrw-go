// Code generated by thriftrw v1.27.0-dev. DO NOT EDIT.
// @generated

package hyphenated_file

import (
	errors "errors"
	fmt "fmt"
	multierr "go.uber.org/multierr"
	non_hyphenated "go.uber.org/thriftrw/gen/internal/tests/non_hyphenated"
	thriftreflect "go.uber.org/thriftrw/thriftreflect"
	wire "go.uber.org/thriftrw/wire"
	zapcore "go.uber.org/zap/zapcore"
	strings "strings"
)

type DocumentStructure struct {
	R2 *non_hyphenated.Second `json:"r2,required"`
}

// ToWire translates a DocumentStructure struct into a Thrift-level intermediate
// representation. This intermediate representation may be serialized
// into bytes using a ThriftRW protocol implementation.
//
// An error is returned if the struct or any of its fields failed to
// validate.
//
//   x, err := v.ToWire()
//   if err != nil {
//     return err
//   }
//
//   if err := binaryProtocol.Encode(x, writer); err != nil {
//     return err
//   }
func (v *DocumentStructure) ToWire() (wire.Value, error) {
	var (
		fields [1]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)

	if v.R2 == nil {
		return w, errors.New("field R2 of DocumentStructure is required")
	}
	w, err = v.R2.ToWire()
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 1, Value: w}
	i++

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

// FromWire deserializes a DocumentStructure struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a DocumentStructure struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v DocumentStructure
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *DocumentStructure) FromWire(w wire.Value) error {
	var ptrFields struct {
		R2 non_hyphenated.Second
	}
	_ = ptrFields

	var err error

	r2IsSet := false

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TStruct {
				err = ptrFields.R2.FromWire(field.Value)
				v.R2 = &ptrFields.R2
				if err != nil {
					return err
				}
				r2IsSet = true
			}
		}
	}

	if !r2IsSet {
		return errors.New("field R2 of DocumentStructure is required")
	}

	return nil
}

// String returns a readable string representation of a DocumentStructure
// struct.
func (v *DocumentStructure) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [1]string
	i := 0
	fields[i] = fmt.Sprintf("R2: %v", v.R2)
	i++

	return fmt.Sprintf("DocumentStructure{%v}", strings.Join(fields[:i], ", "))
}

// Equals returns true if all the fields of this DocumentStructure match the
// provided DocumentStructure.
//
// This function performs a deep comparison.
func (v *DocumentStructure) Equals(rhs *DocumentStructure) bool {
	if v == nil {
		return rhs == nil
	} else if rhs == nil {
		return false
	}
	if !v.R2.Equals(rhs.R2) {
		return false
	}

	return true
}

// MarshalLogObject implements zapcore.ObjectMarshaler, enabling
// fast logging of DocumentStructure.
func (v *DocumentStructure) MarshalLogObject(enc zapcore.ObjectEncoder) (err error) {
	if v == nil {
		return nil
	}
	err = multierr.Append(err, enc.AddObject("r2", v.R2))
	return err
}

// GetR2 returns the value of R2 if it is set or its
// zero value if it is unset.
func (v *DocumentStructure) GetR2() (o *non_hyphenated.Second) {
	if v != nil {
		o = v.R2
	}
	return
}

// IsSetR2 returns true if R2 is not nil.
func (v *DocumentStructure) IsSetR2() bool {
	return v != nil && v.R2 != nil
}

// ThriftModule represents the IDL file used to generate this package.
var ThriftModule = &thriftreflect.ThriftModule{
	Name:     "hyphenated_file",
	Package:  "go.uber.org/thriftrw/gen/internal/tests/hyphenated_file",
	FilePath: "hyphenated_file.thrift",
	SHA1:     "efdcd233efa65e3d451cdf36c518da9e2d0c40b1",
	Includes: []*thriftreflect.ThriftModule{
		non_hyphenated.ThriftModule,
	},
	Raw: rawIDL,
}

const rawIDL = "// This file is named hyphenated_file to possibly conflict with the code\n// generated from hyphenated-file.\n\ninclude \"./non_hyphenated.thrift\"\n\nstruct DocumentStructure {\n 1: required non_hyphenated.Second r2\n}\n"
