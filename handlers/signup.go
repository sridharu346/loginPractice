package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	DBconfig "github.com/sridharu346/loginPractice/dbconfig"
	schema "github.com/sridharu346/loginPractice/schema"
	"github.com/sridharu346/loginPractice/validation"
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
	var newCustomer schema.CustomerDetails
	err = json.Unmarshal(body, &newCustomer)
	if err != nil {
		ResponseFormat(w, err.Error(), http.StatusInternalServerError, nil)
		return
	}
	// validation
	validate, err := validation.SignupValidation(db, newCustomer)
	if err != nil {
		ResponseFormat(w, err.Error(), validate, nil)

	}
	//pushing data
	ErrorCode, err := DBconfig.InsertCustomer(db, newCustomer)
	if err != nil {
		ResponseFormat(w, err.Error(), ErrorCode, nil)
		return
	}
	ResponseFormat(w, fmt.Sprintf("Customer %s created successfully", newCustomer.Email), http.StatusCreated, nil)
}
