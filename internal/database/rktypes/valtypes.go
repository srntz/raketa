package rktypes

type RKDatatype int

const (
	DatatypeAny RKDatatype = iota
	DatatypeStr
	DatatypeNum
	DatatypeBool
)
