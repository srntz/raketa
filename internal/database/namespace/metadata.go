package namespace

import (
	"raketa/internal/database/rktypes"
)

type namespaceMetadata struct {
	restrictValueTypeTo rktypes.RKDatatype
}

func newNamespaceMetadata(restrictValueTypeTo rktypes.RKDatatype) *namespaceMetadata {
	return &namespaceMetadata{
		restrictValueTypeTo: restrictValueTypeTo,
	}
}

func InitializeDefaultNamespace() (string, INamespace) {
	return NewNamespace("main", rktypes.DatatypeAny)
}

// TODO implement conversion of storage data at restriction change.
func (n *Namespace) SetRestrictValueTypeTo(restrictTo rktypes.RKDatatype) (rktypes.RKDatatype, error) {
	n.metadata.restrictValueTypeTo = restrictTo
	return n.metadata.restrictValueTypeTo, nil
}

func (n *Namespace) GetRestrictValueTypeTo() rktypes.RKDatatype {
	return n.metadata.restrictValueTypeTo
}
