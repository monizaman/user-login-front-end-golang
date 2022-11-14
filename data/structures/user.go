package structures

type LoginForm struct {
	Email    string
	Password string
}

type RegistrationForm struct {
	Email     string
	Fullname  string
	Telephone string
	Password  string
}

type UpdateProfileForm struct {
	Email     string
	Telephone string
	Fullname  string
}

type UserToken struct {
	Token      string
	Email      string `json:"email,omitempty"`
	Password   string `json:"password,omitempty"`
	Success    bool
	StatusCode int
}

type UserProfile struct {
	Email      string `json:"email,omitempty"`
	Fullname   string `json:"fullname,omitempty"`
	Telephone  string `json:"telephone,omitempty"`
	GoogleId   string `json:"google_id,omitempty"`
	Success    bool
	StatusCode int
}
