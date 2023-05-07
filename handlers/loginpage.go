package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	DBconfig "github.com/sridharu346/loginPractice/dbconfig"
	"github.com/sridharu346/loginPractice/schema"
	"github.com/sridharu346/loginPractice/validation"
)

func LoginpageHandler(w http.ResponseWriter, r *http.Request) {

	//Loginpage
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
	var logincustomer schema.LoginCustomer
	err = json.Unmarshal(body, &logincustomer)
	if err != nil {
		ResponseFormat(w, err.Error(), http.StatusInternalServerError, nil)
		return
	}

	// validation
	validate, err := validation.LoginpageValidation(db, logincustomer)
	if err != nil {
		ResponseFormat(w, "Login failed: "+err.Error(), validate, nil)
		return
	}

	ResponseFormat(w, "Login successful!", http.StatusOK, nil)
}
