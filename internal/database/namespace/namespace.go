package namespace

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

func (namespace *Namespace[T]) CheckHealth() bool {
	if namespace.name != "" && namespace.metadata != nil {
		return true
	}
	return false
}
