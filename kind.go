package json_to_go

type Kind uint

const (
	Invalid Kind = iota
	Bool
	Int64
	Float64
	String
	Struct
	Interface
)

// String returns a string representation of the type.
func (t Kind) String() string {
	switch t {
	default:
		return ""
	case Invalid:
		return "Invalid"
	case Bool:
		return "bool"
	case Int64:
		return "int64"
	case Float64:
		return "float64"
	case String:
		return "string"
	case Struct:
		return "struct"
	case Interface:
		return "any"
	}
}
