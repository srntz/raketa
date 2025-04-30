package rktypes

type RKType struct {
	datatype RKTypeconst
	value    string
}

type IRKType interface {
	ToEnum() RKTypeconst
}
