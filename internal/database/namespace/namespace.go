package namespace

import (
	"errors"
	"fmt"
	"raketa/internal/database/rktypes"
)

type Namespace struct {
	name     string
	storage  map[string]rktypes.IRKType
	metadata *metadata
}

func NewNamespace(name string, datatype rktypes.RKDatatype, metadata *MetadataOptions) (string, INamespace) {
	return name, &Namespace{
		name:     name,
		storage:  map[string]rktypes.IRKType{},
		metadata: newNamespaceMetadata(datatype, metadata),
	}
}

func (n *Namespace) Insert(key string, val rktypes.IRKType) (rktypes.IRKType, error) {
	if val.GetDatatype() != n.metadata.restrictValueTypeTo && n.metadata.restrictValueTypeTo != rktypes.DatatypeAny {
		return nil, errors.New(
			fmt.Sprintf("Type %s does not satisfy the constraint", n.metadata.restrictValueTypeTo.String()),
		)
	}

	if _, ok := n.storage[key]; ok {
		return nil, errors.New(fmt.Sprintf("Key %s already exists", key))
	}

	n.storage[key] = val
	return val, nil
}

func (n *Namespace) Upsert(key string, val rktypes.IRKType) (rktypes.IRKType, error) {
	if val.GetDatatype() != n.metadata.restrictValueTypeTo && n.metadata.restrictValueTypeTo != rktypes.DatatypeAny {
		return nil, errors.New(
			fmt.Sprintf("Value %s does not satisfy the constraint", n.metadata.restrictValueTypeTo.String()),
		)
	}

	n.storage[key] = val
	return val, nil
}

func (n *Namespace) GetName() string {
	return n.name
}

func (namespace *Namespace) CheckHealth() bool {
	if namespace.name != "" && namespace.metadata != nil {
		return true
	}
	return false
}
