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
)

const (
	Camel = iota
	Kebab
	LowerCamel
	Snake
	Lower
	Upper
)

const (
	PosBeginning = iota
	//PosStart PadPosition = iota
	PosBeforeName
	PosAfterName
	PosEnd
)
