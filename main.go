package main

import (
	"go_web/route"
	"log"
)

func main() {
	app := route.Serve()
	err := app.Run(":8000")
	if err != nil {
		log.Panic(err)
	}
}
