package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"user-management-front-end/config"
	"user-management-front-end/routes"
)

func main() {
	config.LoadAppConfig()
	// Initialize the router
	router := mux.NewRouter() //.StrictSlash(true)

	// Register Routes
	routes.RegisterProductRoutes(router)

	http.Handle("/", router)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	port := os.Getenv("PORT")
	err := http.ListenAndServe(":" + port, nil)
	if err != nil {
		log.Println("An error occurred starting HTTP listener at port 6000")
		log.Println("Error: " + err.Error())
	}
}


