package routes

import (
	"github.com/gorilla/mux"
	"user-management-front-end/controllers"
)

func RegisterProductRoutes(router *mux.Router) {
	router.HandleFunc("/", controllers.HomeController).Methods("GET")
	router.HandleFunc("/login", controllers.LoginUserFromController).Methods("GET")
	router.HandleFunc("/login", controllers.LoginUserController).Methods("POST")
	router.HandleFunc("/registration", controllers.RegisterUserController).Methods("GET")
	router.HandleFunc("/registration", controllers.RegisterUserFromController).Methods("POST")
	router.HandleFunc("/edit/profile", controllers.EditProfileController).Methods("GET")
	router.HandleFunc("/edit/profile", controllers.UpdateProfileController).Methods("POST")
	router.HandleFunc("/profile", controllers.UserProfileController).Methods("GET")
	router.HandleFunc("/logout", controllers.LogoutController).Methods("GET")
	router.HandleFunc("/google-login", controllers.GoogleLoginUserController).Methods("POST")
}
