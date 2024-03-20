package main

type UtilType int

const (
	Enum UtilType = iota
)

// Stringer interface for enums
// Similar to toString() in Java
func (ut UtilType) String() string {
	return [...]string{"enum"}[ut]
}

var (
	utilMap = map[string]UtilType{
		"enum": Enum,
	}
)
