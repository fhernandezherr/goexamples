/*

A type determines a set of values together with operations and methods specific to those values. A type may be denoted by a type name, if it has one, or specified using a type literal, which composes a type from existing types.

TypeDecl   = "type" ( TypeSpec | "(" { TypeSpec ";" } ")" ) .
TypeSpec   = AliasDecl | TypeDef .
AliasDecl  = identifier "=" Type .
TypeDef    = identifier Type .
identifier = letter { letter | unicode_digit } .
letter     = unicode_letter | "_" .
Type       = TypeName | TypeLit | "(" Type ")" .
TypeName   = identifier | QualifiedIdent .
TypeLit    = ArrayType | StructType | PointerType | FunctionType |
             InterfaceType | SliceType | MapType | ChannelType .
*/

package Types

import "fmt"

// An alias declaration binds an identifier to the given type.
type a = int

// A type definition creates a new, distinct type with the same underlying type and operations as the given type, and binds an identifier to it
type b int

// Numeric types

type (
	c = string // alias
	d string   // type definition
)

func main() {

	fmt.Printf("%T", a)
}
