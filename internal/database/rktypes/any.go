package rktypes

type RKAny struct {
	typedef RKType
}

func NewRKAny(value string) *RKAny {
	return &RKAny{
		typedef: RKType{
			datatype: DatatypeAny,
			value:    value,
		},
	}
}

func (rktype RKAny) ToEnum() RKDatatype {
	return DatatypeAny
}
