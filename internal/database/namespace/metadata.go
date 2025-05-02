package namespace

import (
	"raketa/internal/database/rktypes"
)

type namespaceMetadata struct {
	restrictValueTypeTo rktypes.RKDatatype
	preventDeletion     bool
	strict              bool
}

type MetadataOptions struct {
	preventDeletion *bool
}

func newNamespaceMetadata(restrictValueTypeTo rktypes.RKDatatype, options *MetadataOptions) *namespaceMetadata {
	m := &namespaceMetadata{
		restrictValueTypeTo: restrictValueTypeTo,
	}

	if options == nil || options.preventDeletion == nil {
		m.preventDeletion = false
	} else {
		m.preventDeletion = *options.preventDeletion
	}

	return m
}

func InitializeDefaultNamespace() (string, INamespace) {
	return "main", &Namespace{
		name:    "main",
		storage: map[string]rktypes.IRKType{},
		metadata: &namespaceMetadata{
			restrictValueTypeTo: rktypes.DatatypeAny,
			strict:              true,
			preventDeletion:     true,
		},
	}
}

// TODO implement conversion of storage data at restriction change.
func (n *Namespace) SetRestrictValueTypeTo(restrictTo rktypes.RKDatatype) (rktypes.RKDatatype, error) {
	n.metadata.restrictValueTypeTo = restrictTo
	return n.metadata.restrictValueTypeTo, nil
}

func (n *Namespace) GetRestrictValueTypeTo() rktypes.RKDatatype {
	return n.metadata.restrictValueTypeTo
}
