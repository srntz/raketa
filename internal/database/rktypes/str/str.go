package rkstr

import "raketa/internal/database/rktypes"

const (
	RKSTR_CHARLIMIT_DEFAULT = -1
)

type RKStr struct {
	rktypes.RKType
}

func NewRKStr(value string) *RKStr {
	return &RKStr{
		rktypes.RKType{
			Datatype: rktypes.DatatypeStr,
			Value:    value,
		},
	}
}

func (str RKStr) ToEnum() rktypes.RKDatatype {
	return rktypes.DatatypeStr
}
