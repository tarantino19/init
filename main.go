package main

import (
	"github.com/tarantino19/init/internal/app"
)

func main(){
	app, err := app.NewApplication()

	if err != nil {
		panic(err)
	}

	app.Logger.Println("The app is running...")

	
}
