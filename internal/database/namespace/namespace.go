package namespace

import (
	"raketa/internal/database/rktypes"
)

type namespaceMetadata struct {
	restrictValueTypeTo rktypes.RKTypeconst
}

type NamespaceStorageValueTypes interface {
	*rktypes.RKStr | *rktypes.RKAny
	rktypes.IRKType
}

type INamespace interface {
	SetRestrictValueTypeTo(restrictTo rktypes.RKTypeconst)
}

type Namespace[T NamespaceStorageValueTypes] struct {
	name     string
	storage  map[string]T
	metadata *namespaceMetadata
}

func NewNamespace[T NamespaceStorageValueTypes](name string, typeConstraint T) (string, INamespace) {
	return name, &Namespace[T]{
		name:     name,
		storage:  map[string]T{},
		metadata: newNamespaceMetadata(typeConstraint),
	}
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
func (namespace *Namespace[T]) SetRestrictValueTypeTo(restrictTo rktypes.RKTypeconst) {
	namespace.metadata.restrictValueTypeTo = restrictTo
}

func (namespace *Namespace[T]) CheckHealth() bool {
	if namespace.name != "" && namespace.metadata != nil {
		return true
	}
	return false
}
