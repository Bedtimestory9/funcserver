package db

type Record interface {
	MoodRecord | UserRecord
}

type Response interface {
	LoginResponse | UserSignupResponse
}

type MoodRecord struct {
	Name string `json:"name"`
	Mood int    `json:"mood"`
}

type UserRecord struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
	// ,string is for marshalling string to int for Go value
	Age int `json:"age,string"`
}

type LoginResponse struct {
	Result      string `json:"result"`
	Message     string `json:"message"`
	RedirectURL string `json:"redirectURL"`
}

type UserSignupResponse struct {
	EmailError     string `json:"emailError"`
	UsernameError  string `json:"usernameError"`
	PasswordError  string `json:"passwordError"`
	AgeError       string `json:"ageError"`
	GeneralMessage string `json:"generalMessage"`
	RedirectURL    string `json:"redirectURL"`
}
