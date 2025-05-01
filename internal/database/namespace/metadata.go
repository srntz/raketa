package namespace

import (
	"raketa/internal/database/rktypes"
)

type namespaceMetadata struct {
	restrictValueTypeTo rktypes.RKDatatype
}

func newNamespaceMetadata(typeConstraint rktypes.IRKType) *namespaceMetadata {
	return &namespaceMetadata{
		restrictValueTypeTo: typeConstraint.ToEnum(),
	}
}

func InitializeDefaultNamespace() (string, INamespace) {
	return NewNamespace("main", &rktypes.RKAny{})
}

// TODO re-initialize namespace storage. create a way to check the validity of type-typeconst pair
func (namespace *Namespace[T]) SetRestrictValueTypeTo(restrictTo rktypes.RKDatatype) {
	namespace.metadata.restrictValueTypeTo = restrictTo
	switch restrictTo {
	case rktypes.DatatypeAny:
		namespace.storage = map[string]*rktypes.RKAny{}
	}
}
