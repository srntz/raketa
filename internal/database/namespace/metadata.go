package namespace

import (
	"raketa/internal/database/rktypes"
	rkstr "raketa/internal/database/rktypes/str"
)

type metadata struct {
	restrictValueTypeTo rktypes.RKDatatype
	preventDeletion     bool
	strict              bool
	strMetadata         *rkstr.Metadata
}

type MetadataOptions struct {
	preventDeletion *bool
}

func newNamespaceMetadata(restrictValueTypeTo rktypes.RKDatatype, options *MetadataOptions) *metadata {
	m := &metadata{
		restrictValueTypeTo: restrictValueTypeTo,
	}

	if restrictValueTypeTo == rktypes.DatatypeStr {
		m.strMetadata = rkstr.NewMetadata()
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
		metadata: &metadata{
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
