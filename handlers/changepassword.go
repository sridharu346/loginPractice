package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	DBconfig "github.com/sridharu346/loginPractice/dbconfig"
	"github.com/sridharu346/loginPractice/schema"
	"github.com/sridharu346/loginPractice/validation"
)

func ChangepasswordHandler(w http.ResponseWriter, r *http.Request) {
	//changepassword

	if r.Method != "PUT" {
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

	// un marshal the body
	var newPassword schema.ChangePassword
	err = json.Unmarshal(body, &newPassword)
	if err != nil {
		ResponseFormat(w, err.Error(), http.StatusInternalServerError, nil)
		return
	}

	//validation
	validate, err := validation.ChangePasswordValidation(db, newPassword)
	if err != nil {
		ResponseFormat(w, err.Error(), validate, nil)
	}
	ResponseFormat(w, "ChangePassword successful!", http.StatusOK, nil)
}
