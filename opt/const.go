package opt

const (
	Clean    = "clean"
	Find     = "find"
	Replace  = "replace"
	Suffix   = "suffix"
	Prefix   = "prefix"
	CWD      = "cwd"
	Casing   = "casing"
	Start    = "start"
	Zeroes   = "zeroes"
	Position = "position"
	Sep      = "sep"
	Asciiify = "asciiify"
)

const (
	Camel = iota
	Kebab
	LowerCamel
	Snake
	Lower
	Upper
)

//go:generate stringer -type PadPosition -linecomment
type PadPosition int

const (
	Beginning  PadPosition = iota
	BeforeName             // Before Name
	AfterName              // After Name
	End
)
