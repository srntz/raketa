package namespace

import (
	"raketa/internal/database/rktypes"
)

type INamespace interface {
	SetRestrictValueTypeTo(restrictTo rktypes.RKDatatype)
}
