package rktypes

type RKAny struct {
	typedef RKType
}

func NewRKAny(value string) *RKAny {
	return &RKAny{
		typedef: RKType{
			datatype: TYPECONST_ANY,
			value:    value,
		},
	}
}

func (rktype RKAny) ToEnum() RKTypeconst {
	return TYPECONST_ANY
}
