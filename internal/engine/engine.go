package engine

import "raketa/internal/database/database"

type Engine struct {
	db map[string]*database.Database
}

func Constructor() *Engine {
	return &Engine{db: map[string]*database.Database{}}
}

func (engine *Engine) Run() {
	var name, defaultDb = database.InitializeDefaultDatabase()
	engine.db[name] = defaultDb
}

func (engine *Engine) GetDatabase(identifier string) *database.Database {
	value, ok := engine.db[identifier]
	if !ok {
		return nil
	}
	return value
}
