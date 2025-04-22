package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/tarantino19/init/internal/app"
	"github.com/tarantino19/init/internal/routes"
)

func main(){
	var port int
	flag.IntVar(&port, "port", 8080, "Go backend server port")
	flag.Parse()
	addr := fmt.Sprintf(":%d", port)

	app, err := app.NewApplication()

	if err != nil {
		panic(err)
	}

	defer app.DB.Close()
	
	//setup route
	r := routes.SetupRoutes(app)

	server := &http.Server{
		Addr: addr,
		Handler: r,
		IdleTimeout: time.Minute,
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}


	app.Logger.Println("We are running on port:", port)


	err = server.ListenAndServe()
	if err != nil {
		app.Logger.Fatal(err)
	}
}