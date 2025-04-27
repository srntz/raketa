package engine

import "raketa/internal/database"

type Engine struct {
	db []*database.Database
}

func Constructor() *Engine {
	return &Engine{db: []*database.Database{}}
}

func (engine *Engine) Run() {
	engine.db = append(engine.db, database.InitializeDefaultDatabase())
}

func (engine *Engine) GetDatabases() []*database.Database {
	return engine.db
}
