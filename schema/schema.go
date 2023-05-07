package schema

type CustomerDetails struct {
	Name            string `json:"Name"`
	Email           string `json:"Email"`
	Password        string `json:"Password"`
	ConfirmPassword string `json:"ConfirmPassword"`
	Phonenumber     string `json:"Phonenumber"`
	Address         string `json:"Address"`
}

type Response struct {
	StatusCode int         `json:"StatusCode"`
	Message    string      `json:"Message"`
	Data       interface{} `json:"Data"`
}

type LoginCustomer struct {
	Email    string `json:"Email"`
	Password string `json:"Password"`
}

type ChangePassword struct {
	Email       string `json:"Email"`
	NewPassword string `json:"NewPassword"`
	Password    string `json:"Password"`
}
