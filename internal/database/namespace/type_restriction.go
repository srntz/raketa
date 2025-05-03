package namespace

import (
	"raketa/internal/database/rktypes"
	rkstr "raketa/internal/database/rktypes/str"
)

type NamespaceStorageValueTypes interface {
	*rkstr.RKStr | *rktypes.RKAny
	rktypes.IRKType
}
