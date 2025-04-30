package database

import (
	"raketa/internal/database/namespace"
)

type databaseMetadata struct {
	deletionAllowed bool
}

type Database struct {
	name       string
	namespaces map[string]namespace.INamespace
	metadata   databaseMetadata
}

func NewDatabase(name string) *Database {
	return &Database{
		name:       name,
		namespaces: map[string]namespace.INamespace{},
		metadata: databaseMetadata{
			deletionAllowed: false,
		},
	}
}

func InitializeDefaultDatabase() (string, *Database) {
	defaultDb := Database{
		name:       "default",
		namespaces: map[string]namespace.INamespace{},
		metadata: databaseMetadata{
			deletionAllowed: false,
		},
	}

	name, namespace := namespace.InitializeDefaultNamespace()

	defaultDb.namespaces[name] = namespace

	return defaultDb.name, &defaultDb
}

func (database *Database) GetNamespace(identifier string) namespace.INamespace {
	value, ok := database.namespaces[identifier]
	if !ok {
		return nil
	}
	return value
}
