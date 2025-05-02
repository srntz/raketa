package namespace

import (
	"raketa/internal/database/rktypes"
)

type INamespace interface {
	Insert(key string, val rktypes.IRKType) (rktypes.IRKType, error)
	Upsert(key string, val rktypes.IRKType) (rktypes.IRKType, error)
	GetName() string
	SetRestrictValueTypeTo(restrictTo rktypes.RKDatatype) (rktypes.RKDatatype, error)
	GetRestrictValueTypeTo() rktypes.RKDatatype
	CheckHealth() bool
}
