package database

type databaseMetadata struct {
	deletionAllowed bool
}

type Database struct {
	name     string
	metadata databaseMetadata
}

func Constructor(name string) *Database {
	return &Database{
		name: name,
		metadata: databaseMetadata{
			deletionAllowed: false,
		},
	}
}

func InitializeDefaultDatabase() *Database {
	return &Database{
		name: "default",
		metadata: databaseMetadata{
			deletionAllowed: false,
		},
	}
}
