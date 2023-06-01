package opt

const (
	Tidy     = "tidy"
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

const (
	Camel      = "camel"
	Kebab      = "kebab"
	LowerCamel = "lowerCamel"
	Snake      = "snake"
	Lower      = "lower"
	Upper      = "upper"
)

//go:generate stringer -type PadPosition -linecomment
type PadPosition int

const (
	Beginning  PadPosition = iota
	BeforeName             // Before Name
	AfterName              // After Name
	End                    // End (default)
)
