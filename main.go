package main

import (
	"blog-api/bootstrap"
	"log"
)

func main() {
	app := bootstrap.InitApp()

	err := app.ListenAndServe()
	if err != nil {
		log.Panicln("ERRORS: ", err.Error())
	}
}
