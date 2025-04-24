package app

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/tarantino19/init/internal/api"
	"github.com/tarantino19/init/internal/store"
	"github.com/tarantino19/init/migrations"
)

type Application struct {
	Logger *log.Logger
	WorkoutHandler *api.WorkoutHandler
	DB *sql.DB
}

func NewApplication()(*Application, error){
	pgDB, err := store.Open()
	if err != nil {
		return nil, err
	}

	err = store.MigrateFS(pgDB, migrations.FS, ".")
	if err != nil {
		panic(err)
	}


	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)
	
	//our stores will go here

	workoutStore := store.NewPostgresWorkoutStore(pgDB)

	//our handler will go here
	workoutHandler := api.NewWorkoutHandler(workoutStore)

	app := &Application{
		Logger: logger,
		WorkoutHandler: workoutHandler,
		DB: pgDB,
	}

	return app, nil
}

//our Application struct now has a method called healthcheck because of this (a *Application)
//use for testing initial routes if working, FprintF returns something
func (a *Application) HealthCheck(w  http.ResponseWriter, r *http.Request){ 
	fmt.Fprintf(w, "Status is Available\n")
}
 