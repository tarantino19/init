package app

import (
	"log"
	"os"
)

//we'll be using logger instead of fmt.println
type Application struct {
	Logger *log.Logger
}

func NewApplication()(*Application, error){
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	app := &Application{
		Logger: logger,
	}

	return app, nil
}