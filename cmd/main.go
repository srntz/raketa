package main

import (
	"fmt"
	"raketa/internal/engine"
)

func main() {
	var engine = engine.Constructor()
	engine.Run()
	fmt.Println(*engine.GetDatabases()[0])
}
