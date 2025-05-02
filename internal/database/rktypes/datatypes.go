package rktypes

type RKDatatype int

const (
	DatatypeAny RKDatatype = iota
	DatatypeStr
	DatatypeNum
	DatatypeBool
)

func (dt RKDatatype) String() string {
	switch dt {
	case DatatypeAny:
		return "any"
	case DatatypeStr:
		return "str"
	case DatatypeNum:
		return "num"
	case DatatypeBool:
		return "bool"
	default:
		return "unknown"
	}
}
