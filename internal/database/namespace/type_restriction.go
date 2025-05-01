package namespace

import (
	"raketa/internal/database/rktypes"
)

type NamespaceStorageValueTypes interface {
	*rktypes.RKStr | *rktypes.RKAny
	rktypes.IRKType
}
