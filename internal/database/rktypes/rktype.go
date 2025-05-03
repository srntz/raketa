package rktypes

type RKType struct {
	Datatype RKDatatype
	Value    string
}

type IRKType interface {
	GetDatatype() RKDatatype
}
