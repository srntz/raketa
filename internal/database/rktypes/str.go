package rktypes

import (
	"errors"
)

const (
	RKSTR_CHARLIMIT_DEFAULT = -1
)

type rkStrMetadata struct {
	charLimit int
}

type RKStr struct {
	typedef  RKType
	metadata rkStrMetadata
}

func NewRKStr(value string) *RKStr {
	return &RKStr{
		typedef: RKType{
			datatype: DatatypeStr,
			value:    value,
		},
		metadata: rkStrMetadata{
			charLimit: RKSTR_CHARLIMIT_DEFAULT,
		},
	}
}

func (str RKStr) ToEnum() RKDatatype {
	return DatatypeStr
}

func (str *RKStr) SetCharLimit(limit int) error {
	if limit <= 0 {
		return errors.New("CharLimit on type STR can not be smaller than 1")
	}
	str.metadata.charLimit = limit
	return nil
}
