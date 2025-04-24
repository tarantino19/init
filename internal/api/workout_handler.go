package api

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/tarantino19/init/internal/store"
)

type WorkoutHandler struct {
	workoutStore store.WorkoutStore
}

//returns a pointer w/data type - WorkoutHandler (struct)
func NewWorkoutHandler(workoutStore store.WorkoutStore) *WorkoutHandler{
	return &WorkoutHandler{
		workoutStore: workoutStore,
	}
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

	workout, err := wh.workoutStore.GetWorkoutByID(workoutID)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "failed to fetch the workout", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(workout)

	fmt.Fprintf(w, "this is the workout id %d\n", workoutID)
}

func (wh *WorkoutHandler) HandleCreateWorkout(w http.ResponseWriter, r *http.Request){
	
	//struct from workout store
	var workout store.Workout

	err := json.NewDecoder(r.Body).Decode(&workout) //quite similar to unmarshalling but new decoder for outside req

	if err != nil {
		fmt.Print(err)
		http.Error(w, "failed to create workout", http.StatusInternalServerError)
		return //don't forget this
	}

	//creates the workout db
	createdWorkout, err := wh.workoutStore.CreateWorkout(&workout)

	if err != nil {
		fmt.Println(err)
		http.Error(w, "failed to create workout", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(createdWorkout)


}

func (wh *WorkoutHandler) HandleUpdateWorkoutById(w http.ResponseWriter, r *http.Request){
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

	existingWorkout, err := wh.workoutStore.GetWorkoutByID(workoutID)

	if err != nil {
		http.Error(w, "failed to fetch workout", http.StatusInternalServerError)
		return
	}

	if existingWorkout == nil {
		http.NotFound(w, r)
		return
	}

	
	// at this point we can assume we are able to find an existing workout
	//pointers to see value actually exist
	var updateWorkoutRequest struct {
			Title           *string              `json:"title"`
			Description     *string              `json:"description"`
			DurationMinutes *int                 `json:"duration_minutes"`
			CaloriesBurned  *int                 `json:"calories_burned"`
			Entries         []store.WorkoutEntry `json:"entries"`
		}

	err = json.NewDecoder(r.Body).Decode(&updateWorkoutRequest)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	if updateWorkoutRequest.Title != nil {
		existingWorkout.Title = *updateWorkoutRequest.Title
	}
	if updateWorkoutRequest.Description != nil {
		existingWorkout.Description = *updateWorkoutRequest.Description
	}
	if updateWorkoutRequest.DurationMinutes != nil {
		existingWorkout.DurationMinutes = *updateWorkoutRequest.DurationMinutes
	}
	if updateWorkoutRequest.CaloriesBurned != nil {
		existingWorkout.CaloriesBurned = *updateWorkoutRequest.CaloriesBurned
	}
	if updateWorkoutRequest.Entries != nil {
		existingWorkout.Entries = updateWorkoutRequest.Entries
	}

	//ypdating the workout on the db
	err = wh.workoutStore.UpdateWorkout(existingWorkout)

	if err != nil {
		fmt.Println("update workout error", err)
		http.Error(w, "failed to update the workout", http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(existingWorkout)
	//returning the workout as response

}

func (wh *WorkoutHandler) HandleDeleteWorkoutByID (w http.ResponseWriter, r * http.Request){
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

	err = wh.workoutStore.DeleteWorkout(workoutID)

	if err == sql.ErrNoRows {
		http.Error(w, "workout not found", http.StatusNotFound)
	}

	if err != nil {
		http.Error(w, "error deleting workout", http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusNoContent)


}