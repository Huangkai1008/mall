package main

import (
	"log"

	"mall/internal/pkg/application"
)

func main() {
	app, err := application.New()
	if err != nil {
		log.Printf("%+v\n", err)
		panic(err)
	}

	app.AwaitSignal()
}
