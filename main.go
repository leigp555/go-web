package main

import (
	"go_web/model"
	"go_web/route"
	"log"
)

func main() {
	model.InitDB()
	app := route.Serve()
	err := app.Run(":8000")
	if err != nil {
		log.Panic(err)
	}
}
