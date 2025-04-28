package namespace

import "raketa/internal/database/rktypes"

type namespaceMetadata struct {
	restrictValueTypeTo rktypes.RKTypeconst
}

func newNamespaceMetadata() *namespaceMetadata {
	return &namespaceMetadata{
		restrictValueTypeTo: rktypes.TYPECONST_ANY,
	}
}

type Namespace struct {
	name     string
	metadata *namespaceMetadata
}

func NewNamespace(name string) (string, *Namespace) {
	return name, &Namespace{name: name, metadata: newNamespaceMetadata()}
}

func InitializeDefaultNamespace() (string, *Namespace) {
	return "main", &Namespace{name: "main", metadata: newNamespaceMetadata()}
}

func (namespace *Namespace) SetRestrictValueTypeTo(restrictTo rktypes.RKTypeconst) {
	namespace.metadata.restrictValueTypeTo = restrictTo
}

func (namespace *Namespace) CheckHealth() bool {
	if namespace.name != "" && namespace.metadata != nil {
		return true
	}
	return false
}
