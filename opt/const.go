package opt

const (
	Clean    = "clean"
	Find     = "find"
	Replace  = "replace"
	Suffix   = "suffix"
	Prefix   = "prefix"
	Dir      = "dir"
	Casing   = "casing"
	Start    = "min"
	Zeroes   = "zeroes"
	Position = "position"
	Sep      = "sep"
	Asciiify = "asciiify"
	Name     = "name"
)

//go:generate stringer -type Case -linecomment
type Case int

const (
	Camel      Case = iota // CamelCase
	Kebab                  // kebab-case
	LowerCamel             // lowerCamel
	Snake                  // snake_case (default)
	Lower                  // lower case
	Upper                  // UPPER CASE
)

//go:generate stringer -type PadPosition -linecomment
type PadPosition int

const (
	Beginning  PadPosition = iota
	BeforeName             // Before Name
	AfterName              // After Name
	End                    // End (default)
)
