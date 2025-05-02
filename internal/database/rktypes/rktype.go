package rktypes

type RKType struct {
	datatype RKDatatype
	value    string
}

type IRKType interface {
	GetDatatype() RKDatatype
}
