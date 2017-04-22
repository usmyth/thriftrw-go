// Code generated by thriftrw v1.3.0
// @generated

package collision

import "go.uber.org/thriftrw/thriftreflect"

var ThriftModule = &thriftreflect.ThriftModule{Name: "collision", Package: "go.uber.org/thriftrw/gen/testdata/collision", FilePath: "collision.thrift", SHA1: "4492031685a099efe37d2e1377e5152e26a6e7c5", Raw: rawIDL}

const rawIDL = "\nstruct StructCollision {\n\t1: required bool collisionField\n\t2: required string collision_field (go.name = \"CollisionField2\")\n}\n\nstruct struct_collision {\n\t1: required bool collisionField\n\t2: required string collision_field (go.name = \"CollisionField2\")\n} (go.name=\"StructCollision2\")\n\nstruct PrimitiveContainers {\n    1: optional list<string> ListOrSetOrMap (go.name = \"A\")\n    3: optional set<string>  List_Or_SetOrMap (go.name = \"B\")\n    5: optional map<string, string> ListOrSet_Or_Map (go.name = \"C\")\n}\n\nenum MyEnum {\n    X = 123,\n    Y = 456,\n    Z = 789,\n    FooBar,\n    foo_bar (go.name=\"FooBar2\"),\n}\n\nenum my_enum {\n    X = 12,\n    Y = 34,\n    Z = 56,\n} (go.name=\"MyEnum2\")\n\ntypedef i64 LittlePotatoe\ntypedef double little_potatoe (go.name=\"LittlePotatoe2\")\n\nconst struct_collision struct_constant = {\n\t\"collisionField\": false,\n\t\"collision_field\": \"false indeed\",\n}\n\nunion UnionCollision {\n\t1: bool collisionField\n\t2: string collision_field (go.name = \"CollisionField2\")\n}\n\nunion union_collision {\n\t1: bool collisionField\n\t2: string collision_field (go.name = \"CollisionField2\")\n} (go.name=\"UnionCollision2\")\n\nstruct WithDefault {\n\t1: required struct_collision pouet = struct_constant\n}\n"
