package schema

type signUpDetails struct {
	Name        string `json:"Name"`
	Email       string `json:"Email"`
	Password    string `json:"Password"`
	Phonenumber string `json:"Phonenumber"`
	Address     string `json:"Address"`
}

type response struct {
	StatusCode string `json:"StatusCode"`
	Message    string `json:"Message"`
	Data       string `json:"Data"`
}
