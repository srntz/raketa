package rktypes

import "errors"

const (
	RKSTR_CHARLIMIT_DEFAULT = -1
)

type rkStrMetadata struct {
	charLimit int
}

type RKStr struct {
	name     RKTypeconst
	metadata rkStrMetadata
}

func NewRKStr() *RKStr {
	return &RKStr{
		name: TYPECONST_STR,
		metadata: rkStrMetadata{
			charLimit: RKSTR_CHARLIMIT_DEFAULT,
		},
	}
}

func (str *RKStr) SetCharLimit(limit int) error {
	if limit <= 0 {
		return errors.New("CharLimit on type STR can not be smaller than 1")
	}
	str.metadata.charLimit = limit
	return nil
}
