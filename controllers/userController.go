package controllers

import (
	"fmt"
	"html/template"
	"net/http"
	"time"
	"user-management-front-end/data/structures"
	"user-management-front-end/requests"
)

func LoginUserFromController(w http.ResponseWriter, r *http.Request) {
	viewTemplate, _ := template.ParseFiles("view/login.html")
	c, err := r.Cookie("token")
	if err == nil && c.Value != ""{
		userTokenResponseData := make(chan structures.UserProfile)
		go requests.GetUserProfile(c.Value, userTokenResponseData)
		userProfile := <-userTokenResponseData
		if userProfile.Email != "" {
			http.Redirect(w, r, "/profile", http.StatusFound)
		}
	}
	err = viewTemplate.Execute(w, map[string]string{"email": "", "password": ""})
	if err != nil {
		return
	}
}

func HomeController(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/login", http.StatusFound)
}

func LoginUserController(w http.ResponseWriter, r *http.Request) {
	userInfo := structures.LoginForm{
		Email:   r.FormValue("email"),
		Password: r.FormValue("password"),
	}
	userTokenResponseData := make(chan structures.UserToken)
	go requests.GetAuthToken(userInfo, userTokenResponseData)
	userToken := <-userTokenResponseData
	if userToken.Token != "" {
		http.SetCookie(w, &http.Cookie{
			Name:    "token",
			Value:   userToken.Token,
			Expires: time.Now().Add(60 * time.Minute),
		})
		http.Redirect(w, r, "/profile", http.StatusFound)
	}else {
		viewTemplate, _ := template.ParseFiles("view/login.html")
		err := viewTemplate.Execute(w, userToken)
		if err != nil {
			return
		}
	}

}

func RegisterUserFromController(w http.ResponseWriter, r *http.Request) {
	userInfo := structures.RegistrationForm{
		Email:   r.FormValue("email"),
		Fullname:   r.FormValue("fullname"),
		Telephone:   r.FormValue("telephone"),
		Password: r.FormValue("password"),
	}
	userRegistrationResponseData := make(chan structures.UserProfile)
	go requests.UserRegistration(userInfo, userRegistrationResponseData)
	userRgRes := <-userRegistrationResponseData
	if userRgRes.StatusCode == 200 {
		http.Redirect(w, r, "/login", http.StatusFound)
	}else {
		viewTemplate, _ := template.ParseFiles("view/registration.html")
		err := viewTemplate.Execute(w, map[string]string{"error": "Registration Failed"})
		if err != nil {
			return
		}
	}
}

func RegisterUserController(w http.ResponseWriter, r *http.Request) {
	viewTemplate, _ := template.ParseFiles("view/registration.html")
	err := viewTemplate.Execute(w, "")
	if err != nil {
		return
	}
}

func EditProfileController(w http.ResponseWriter, r *http.Request) {
	viewTemplate, _ := template.ParseFiles("view/edit.profile.html")
	c, err := r.Cookie("token")
	if err != nil {
		// TODO need to implement refresh token functionality
		if err == http.ErrNoCookie {
			// If the cookie is not set, return an unauthorized status
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}
	userProfileResponseData := make(chan structures.UserProfile)
	go requests.GetUserProfile(c.Value, userProfileResponseData)
	userProfile := <-userProfileResponseData
	err = viewTemplate.Execute(w, userProfile)
	if err != nil {
		return
	}
}

func UpdateProfileController(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			// TODO need to implement refresh token functionality
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}
	userInfo := structures.UpdateProfileForm{
		Email:   r.FormValue("email"),
		Fullname: r.FormValue("fullname"),
		Telephone: r.FormValue("telephone"),
	}
	userProfileResponseData := make(chan structures.UserProfile)
	go requests.UpdateProfile(c.Value, userInfo, userProfileResponseData)
	userToken := <-userProfileResponseData
	if userToken.StatusCode == 200 {
		http.Redirect(w, r, "/profile", http.StatusFound)
	}else {
		viewTemplate, _ := template.ParseFiles("view/edit.profile.html")
		err := viewTemplate.Execute(w, map[string]string{"error": "Edit Profile Failed"})
		if err != nil {
			return
		}
	}
}

func UserProfileController(w http.ResponseWriter, r *http.Request) {
	viewTemplate, _ := template.ParseFiles("view/profile.html")
	c, err := r.Cookie("token")
	if err != nil {
		// TODO need to implement refresh token functionality
		if err == http.ErrNoCookie {
			// If the cookie is not set, return an unauthorized status
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}
	userTokenResponseData := make(chan structures.UserProfile)
	go requests.GetUserProfile(c.Value, userTokenResponseData)
	userProfile := <-userTokenResponseData
	if userProfile.StatusCode != 200 {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}
	fmt.Println(userProfile.StatusCode, "userProfile.StatusCode")
	fmt.Println(userProfile.Fullname, "userProfile.Fullname")
	err = viewTemplate.Execute(w, userProfile)
	if err != nil {
		return
	}
}


func GoogleLoginUserController(w http.ResponseWriter, r *http.Request) {
	userTokenResponseData := make(chan structures.UserToken)
	fmt.Println(r.FormValue("credential"), "r.FormValue(credential)")
	go requests.GetGoogleLoginToken(r.FormValue("credential"), userTokenResponseData)
	userToken := <-userTokenResponseData
	fmt.Println(userToken.Token, "userToken.Token")
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   userToken.Token,
		Expires: time.Now().Add(60 * time.Minute),
	})
	http.Redirect(w, r, "/profile", http.StatusFound)
}

func LogoutController(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   "",
		Expires: time.Now().Add(0 * time.Minute),
	})
	http.Redirect(w, r, "/login", http.StatusFound)
}

// GetNewToken TODO Need to implement Refresh token functionality
func updateUserToken(w http.ResponseWriter, r *http.Request)  {
	_, err := r.Cookie("token")
	refreshToken, refreshTokenErr := r.Cookie("token")
	if err == http.ErrNoCookie && refreshTokenErr != http.ErrNoCookie{
		userTokenResponseData := make(chan structures.UserToken)
		go requests.GetNewToken(refreshToken.Value, userTokenResponseData)
		newToken := <-userTokenResponseData
		if newToken.StatusCode != 200 {
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}
		http.SetCookie(w, &http.Cookie{
			Name:    "token",
			Value:   newToken.Token,
			Expires: time.Now().Add(60 * time.Minute),
		})
	} else {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}