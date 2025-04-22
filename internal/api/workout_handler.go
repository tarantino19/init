package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type WorkoutHandler struct {

}

//returns a pointer w/data type - WorkoutHandler (struct)
func NewWorkoutHandler() *WorkoutHandler{
	return &WorkoutHandler{}
}

//create methods for the handler
//(wh *WorkoutHandler) - this mean WorkoutHandler can access this function
func (wh *WorkoutHandler) HandleGetWorkoutByID(w http.ResponseWriter, r *http.Request){
	paramsWorkoutID := chi.URLParam(r, "id") //the slug (id) is defined in our routes

	if paramsWorkoutID == ""{
		http.NotFound(w, r)
		return
	}

	workoutID, err := strconv.ParseInt(paramsWorkoutID, 10, 64)

	if err != nil {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "this is the workout id %d\n", workoutID)

}

func (wh *WorkoutHandler) HandleCreateWorkout(w http.ResponseWriter, r *http.Request){
	
	fmt.Fprintf(w, "create a workout\n")

}