package rktypes

type RKAny struct {
	typedef RKType
}

func NewRKAny(value string) *RKAny {
	return &RKAny{
		typedef: RKType{
			Datatype: DatatypeAny,
			Value:    value,
		},
	}
}

func (rktype RKAny) ToEnum() RKDatatype {
	return DatatypeAny
}
