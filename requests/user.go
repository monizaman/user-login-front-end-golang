package requests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"io/ioutil"
	"net/http"
	"time"
	"user-management-front-end/data/structures"
)

// GetAuthToken TODO Has a duplicate code need to optimize
func GetAuthToken(loginForm structures.LoginForm, response chan structures.UserToken)  {
	apiUrl := viper.GetString("api.base_url") + "/api/login"
	loginFormJSON, err := json.Marshal(loginForm)
	req, _ := http.NewRequest("POST", apiUrl, bytes.NewBuffer(loginFormJSON))
	var responseObject structures.UserToken
	responseObject.Success = true
	responseObject.StatusCode = 200
	if err != nil {
		responseObject.Success = false
		responseObject.StatusCode = 400
		response <- responseObject
	}
	client := &http.Client{Timeout: time.Second * 60}
	resp, err := client.Do(req)

	if err != nil {
		responseObject.Success = false
		responseObject.StatusCode = 500
		response <- responseObject
		return
	}

	defer func() {
		err := resp.Body.Close()
		if err != nil {
			responseObject.Success = false
			response <- responseObject
			return
		}
	}()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		responseObject.Success = false
		responseObject.StatusCode = 500
		response <- responseObject
		return
	}
	err = json.Unmarshal(body, &responseObject)
	if err != nil {
		responseObject.Success = false
		response <- responseObject
		return
	}
	response <- responseObject
}

// GetGoogleLoginToken TODO Has a duplicate code need to optimize
func GetGoogleLoginToken(credential string, response chan structures.UserToken)  {
	apiUrl := viper.GetString("api.base_url") +  "/api/google-login"
	googleCredential := structures.GoogleToken{Credential: credential}
	credentialJSON, err := json.Marshal(googleCredential)
	fmt.Println(string(credentialJSON), "credentialJSON")
	req, _ := http.NewRequest("POST", apiUrl, bytes.NewBuffer(credentialJSON))
	var responseObject structures.UserToken
	responseObject.Success = true
	responseObject.StatusCode = 200
	if err != nil {
		responseObject.Success = false
		responseObject.StatusCode = 400
		response <- responseObject
	}
	client := &http.Client{Timeout: time.Second * 60}
	resp, err := client.Do(req)

	if err != nil {
		responseObject.Success = false
		responseObject.StatusCode = 500
		response <- responseObject
		return
	}

	defer func() {
		err := resp.Body.Close()
		if err != nil {
			responseObject.Success = false
			response <- responseObject
			return
		}
	}()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		responseObject.Success = false
		responseObject.StatusCode = 500
		response <- responseObject
		return
	}
	err = json.Unmarshal(body, &responseObject)
	if err != nil {
		responseObject.Success = false
		response <- responseObject
		return
	}
	response <- responseObject
}

// GetUserProfile TODO Has a duplicate code need to optimize
func GetUserProfile(jwtToken string, response chan structures.UserProfile)  {
	apiUrl := viper.GetString("api.base_url") + "/api/profile"
	req, err := http.NewRequest("GET", apiUrl, nil)
	var bearer = "Bearer " + jwtToken
	// add authorization header to the req
	req.Header.Add("Authorization", bearer)
	var responseObject structures.UserProfile
	responseObject.Success = true
	if err != nil {
		responseObject.Success = false
		responseObject.StatusCode = 400
		response <- responseObject
	}
	client := &http.Client{Timeout: time.Second * 60}
	resp, err := client.Do(req)

	if err != nil {
		responseObject.Success = false
		responseObject.StatusCode = 500
		response <- responseObject
		return
	}

	defer func() {
		err := resp.Body.Close()
		if err != nil {
			responseObject.Success = false
			response <- responseObject
			return
		}
	}()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		responseObject.Success = false
		responseObject.StatusCode = 500
		response <- responseObject
		return
	}
	err = json.Unmarshal(body, &responseObject)
	if err != nil {
		responseObject.Success = false
		response <- responseObject
		return
	}
	responseObject.StatusCode = resp.StatusCode
	response <- responseObject
}

// UpdateProfile TODO Has a duplicate code need to optimize
func UpdateProfile(jwtToken string, user structures.UpdateProfileForm, response chan structures.UserProfile)  {
	apiUrl := viper.GetString("api.base_url") + "/api/user/update"
	userJSON, err := json.Marshal(user)
	req, _ := http.NewRequest("PATCH", apiUrl, bytes.NewBuffer(userJSON))
	var bearer = "Bearer " + jwtToken
	req.Header.Add("Authorization", bearer)
	var responseObject structures.UserProfile
	responseObject.Success = true
	if err != nil {
		responseObject.Success = false
		responseObject.StatusCode = 400
		response <- responseObject
	}
	client := &http.Client{Timeout: time.Second * 60}
	resp, err := client.Do(req)
	if err != nil {
		responseObject.Success = false
		responseObject.StatusCode = 500
		response <- responseObject
		return
	}

	defer func() {
		err := resp.Body.Close()
		if err != nil {
			responseObject.Success = false
			response <- responseObject
			return
		}
	}()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		responseObject.Success = false
		responseObject.StatusCode = 500
		response <- responseObject
		return
	}
	err = json.Unmarshal(body, &responseObject)
	if err != nil {
		responseObject.Success = false
		response <- responseObject
		return
	}
	responseObject.StatusCode = resp.StatusCode
	response <- responseObject
}


func UserRegistration(userRgFrom structures.RegistrationForm, response chan structures.UserProfile)  {
	apiUrl := viper.GetString("api.base_url") +  "/api/registration"
	userRgFromJSON, err := json.Marshal(userRgFrom)
	req, _ := http.NewRequest("POST", apiUrl, bytes.NewBuffer(userRgFromJSON))
	var responseObject structures.UserProfile
	responseObject.Success = true
	if err != nil {
		responseObject.Success = false
		responseObject.StatusCode = 400
		response <- responseObject
	}
	client := &http.Client{Timeout: time.Second * 60}
	resp, err := client.Do(req)
	if err != nil {
		responseObject.Success = false
		responseObject.StatusCode = 500
		response <- responseObject
		return
	}

	defer func() {
		err := resp.Body.Close()
		if err != nil {
			responseObject.Success = false
			response <- responseObject
			return
		}
	}()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		responseObject.Success = false
		responseObject.StatusCode = 500
		response <- responseObject
		return
	}
	err = json.Unmarshal(body, &responseObject)
	if err != nil {
		responseObject.Success = false
		response <- responseObject
		return
	}
	responseObject.StatusCode = resp.StatusCode
	response <- responseObject
}

// GetNewToken TODO Need to implement Refresh token functionality
func GetNewToken(jwtToken string, response chan structures.UserToken)  {
	apiUrl := viper.GetString("api.base_url") + "/api/refresh-token"
	req, err := http.NewRequest("GET", apiUrl, nil)
	var bearer = "Bearer " + jwtToken
	req.Header.Add("Authorization", bearer)
	var responseObject structures.UserToken
	responseObject.Success = true
	if err != nil {
		responseObject.Success = false
		responseObject.StatusCode = 400
		response <- responseObject
	}
	client := &http.Client{Timeout: time.Second * 60}
	resp, err := client.Do(req)

	if err != nil {
		responseObject.Success = false
		responseObject.StatusCode = 500
		response <- responseObject
		return
	}

	defer func() {
		err := resp.Body.Close()
		if err != nil {
			responseObject.Success = false
			response <- responseObject
			return
		}
	}()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		responseObject.Success = false
		responseObject.StatusCode = 500
		response <- responseObject
		return
	}
	err = json.Unmarshal(body, &responseObject)
	if err != nil {
		responseObject.Success = false
		response <- responseObject
		return
	}
	responseObject.StatusCode = resp.StatusCode
	response <- responseObject
}