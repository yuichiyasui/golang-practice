package main

import "app/router"

func main() {
	engine := router.Load()

	engine.Run()
}
