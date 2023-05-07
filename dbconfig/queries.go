package dbconfig

import (
	"database/sql"
	"errors"
	"net/http"

	schema "github.com/sridharu346/loginPractice/schema"
)

func IsCustomerExist(db *sql.DB, email string) bool {
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM customers WHERE Email = ?", email).Scan(&count)
	if err != nil {
		return false
	}
	if count > 0 {
		return true
	}
	return false
}

func InsertCustomer(db *sql.DB, newCustomer schema.CustomerDetails) (int, error) {

	stmt, err := db.Prepare(`INSERT INTO customers(Name,Email,PhoneNumber,Password,Address) Values(?,?,?,?,?)`)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	_, err = stmt.Exec(newCustomer.Name, newCustomer.Email, newCustomer.Phonenumber, newCustomer.Password, newCustomer.Address)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusCreated, nil
}

func IsLoginvalid(db *sql.DB, email string, password string) (int, error) {
	var Password string
	err := db.QueryRow("SELECT Password FROM customers WHERE Email = ?", email).Scan(&Password)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	if Password == "" {
		return http.StatusBadRequest, errors.New("Email not found")
	}
	if Password != password {
		return http.StatusUnauthorized, errors.New("Invalid password")
	}
	return http.StatusOK, nil
}

func IsChangepasswordValid(db *sql.DB, email string, password string) (int, error) {
	var Password string
	err := db.QueryRow("SELECT Password FROM customers WHERE Email = ?", email).Scan(&Password)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	if Password == "" {
		return http.StatusBadRequest, errors.New("Email not found")
	}
	if Password != password {
		return http.StatusUnauthorized, errors.New("Invalid password")
	}
	return http.StatusOK, nil
}
