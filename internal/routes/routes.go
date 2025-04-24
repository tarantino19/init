package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/tarantino19/init/internal/app"
)

//funcs are first class citizens in go that's why no need to call
func SetupRoutes (app *app.Application) *chi.Mux {
	r := chi.NewRouter()
 
	r.Get("/health", app.HealthCheck)
	r.Get("/workouts/{id}", app.WorkoutHandler.HandleGetWorkoutByID)

	r.Post("/workouts", app.WorkoutHandler.HandleCreateWorkout)
	r.Put("/workouts/{id}", app.WorkoutHandler.HandleUpdateWorkoutById)

	r.Delete("/workouts/{id}", app.WorkoutHandler.HandleDeleteWorkoutByID)
	return r
} 