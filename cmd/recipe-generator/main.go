package main

import (
	"reciper/recipe-generator/internal/env"
	"reciper/recipe-generator/internal/server"
)

func main() {
	env.GetEnv()
	server.CreateServer()
}
