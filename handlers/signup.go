package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	DBconfig "github.com/sridharu346/loginPractice/dbconfig"
)

func SignupHandler(w http.ResponseWriter, r *http.Request) {
	//signup
	if r.Method != "POST" {
		ResponseFormat(w, "Method not allowed", http.StatusMethodNotAllowed, nil)
		return
	}
	db, err := DBconfig.DBConnection()
	if err != nil {
		ResponseFormat(w, "DB connection failed", http.StatusInternalServerError, nil)
	}
	defer db.Close()
	//reading the body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		ResponseFormat(w, err.Error(), http.StatusInternalServerError, nil)
		return
	}
	// unmarshal the body
	var newCustomer Schema.signUpDetails
	err = json.Unmarshal(body, &newCustomer)
	if err != nil {
		ResponseFormat(w, err.Error(), http.StatusInternalServerError, nil)
		return
	}

}
