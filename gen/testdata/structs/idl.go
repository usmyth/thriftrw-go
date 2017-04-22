// Code generated by thriftrw v1.3.0
// @generated

package structs

import (
	"go.uber.org/thriftrw/gen/testdata/enums"
	"go.uber.org/thriftrw/thriftreflect"
)

var ThriftModule = &thriftreflect.ThriftModule{Name: "structs", Package: "go.uber.org/thriftrw/gen/testdata/structs", FilePath: "structs.thrift", SHA1: "d019aed07c659d5f052a1a1992e658e0c6851597", Includes: []*thriftreflect.ThriftModule{enums.ThriftModule}, Raw: rawIDL}

const rawIDL = "include \"./enums.thrift\"\n\nstruct EmptyStruct {}\n\n//////////////////////////////////////////////////////////////////////////////\n// Structs with primitives\n\nstruct PrimitiveRequiredStruct {\n    1: required bool boolField\n    2: required byte byteField\n    3: required i16 int16Field\n    4: required i32 int32Field\n    5: required i64 int64Field\n    6: required double doubleField\n    7: required string stringField\n    8: required binary binaryField\n}\n\nstruct PrimitiveOptionalStruct {\n    1: optional bool boolField\n    2: optional byte byteField\n    3: optional i16 int16Field\n    4: optional i32 int32Field\n    5: optional i64 int64Field\n    6: optional double doubleField\n    7: optional string stringField\n    8: optional binary binaryField\n}\n\n//////////////////////////////////////////////////////////////////////////////\n// Nested structs (Required)\n\nstruct Point {\n    1: required double x\n    2: required double y\n}\n\nstruct Size {\n    1: required double width\n    2: required double height\n}\n\nstruct Frame {\n    1: required Point topLeft\n    2: required Size size\n}\n\nstruct Edge {\n    1: required Point startPoint\n    2: required Point endPoint\n}\n\nstruct Graph {\n    1: required list<Edge> edges\n}\n\n//////////////////////////////////////////////////////////////////////////////\n// Nested structs (Optional)\n\nstruct ContactInfo {\n    1: required string emailAddress\n}\n\nstruct User {\n    1: required string name\n    2: optional ContactInfo contact\n}\n\n//////////////////////////////////////////////////////////////////////////////\n// self-referential struct\n\ntypedef Node List\n\nstruct Node {\n    1: required i32 value\n    2: optional List tail\n}\n\n\n//////////////////////////////////////////////////////////////////////////////\n// Default values\n\nstruct DefaultsStruct {\n    1: required i32 requiredPrimitive = 100\n    2: optional i32 optionalPrimitive = 200\n\n    3: required enums.EnumDefault requiredEnum = enums.EnumDefault.Bar\n    4: optional enums.EnumDefault optionalEnum = 2\n\n    5: required list<string> requiredList = [\"hello\", \"world\"]\n    6: optional list<double> optionalList = [1, 2.0, 3]\n\n    7: required Frame requiredStruct = {\n        \"topLeft\": {\"x\": 1, \"y\": 2},\n        \"size\": {\"width\": 100, \"height\": 200},\n    }\n    8: optional Edge optionalStruct = {\n        \"startPoint\": {\"x\": 1, \"y\": 2},\n        \"endPoint\":   {\"x\": 3, \"y\": 4},\n    }\n}\n"
