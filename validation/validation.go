package validation

import (
	"database/sql"
	"errors"
	"net/http"

	DBconfig "github.com/sridharu346/loginPractice/dbconfig"
	schema "github.com/sridharu346/loginPractice/schema"
)

func SignupValidation(db *sql.DB, newCustomer schema.CustomerDetails) (int, error) {

	if !IsEmailValid(newCustomer.Email) {
		return http.StatusBadRequest, errors.New("Email is not valid")
	}
	if DBconfig.IsCustomerExist(db, newCustomer.Email) {
		return http.StatusForbidden, errors.New("Customers already exists")
	}
	if !IsPasswordValid(newCustomer.Password) {
		return http.StatusBadRequest, errors.New("Password is not valid")
	}
	if newCustomer.Password != newCustomer.ConfirmPassword {
		return http.StatusBadRequest, errors.New("Password Mismatch")

	}
	if !IsPhoneNumberValid(newCustomer.Phonenumber) {
		return http.StatusBadRequest, errors.New("PhoneNumber is invalid")
	}
	if newCustomer.Name == "" {
		return http.StatusBadRequest, errors.New("Name is required")
	}
	return http.StatusCreated, nil

}

func LoginpageValidation(db *sql.DB, logincustomer schema.LoginCustomer) (int, error) {

	if !IsEmailValid(logincustomer.Email) {
		return http.StatusBadRequest, errors.New("Email is not valid")
	}

	if !IsPasswordValid(logincustomer.Password) {
		return http.StatusUnauthorized, errors.New("Password is not valid")
	}

	status, err := DBconfig.IsLoginvalid(db, logincustomer.Email, logincustomer.Password)
	if err != nil {
		return status, err
	}

	return http.StatusAccepted, nil
}

func ChangePasswordValidation(db *sql.DB, logincustomer schema.ChangePassword) (int, error) {
	if !IsEmailValid(logincustomer.Email) {
		return http.StatusBadRequest, errors.New("Email is not valid")
	}

	if !IsPasswordValid(logincustomer.Password) {
		return http.StatusUnauthorized, errors.New("Password is not valid")
	}

	if !IsPasswordValid(logincustomer.NewPassword) {
		return http.StatusUnauthorized, errors.New("NewPassword is not valid")
	}
	status, err := DBconfig.IsChangepasswordValid(db, logincustomer.Email, logincustomer.Password)
	if err != nil {
		return status, err
	}
	return http.StatusAccepted, nil
}
